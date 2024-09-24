package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"goLen/models"
	"net/http"
	"strings"
)

// Функция для поиска записи по BIC
func getBICEntry(db *sql.DB, bic string) (models.BICResponse, error) {
	var response models.BICResponse

	// Получаем основные данные о BIC
	query := `
        SELECT bic, namep
        FROM bic_accounts
        WHERE bic = $1
    `
	row := db.QueryRow(query, bic)
	err := row.Scan(&response.BIC, &response.NameP)
	if err != nil {
		if err == sql.ErrNoRows {
			return response, nil // Если запись не найдена, вернем пустую запись
		}
		return response, fmt.Errorf("ошибка при выполнении запроса: %v", err)
	}

	// Получаем аккаунты, связанные с данным BIC
	queryAccounts := `
        SELECT account, regulation_account_type, ck, account_cbr_bic, date_in, account_status
        FROM bic_accounts
        WHERE bic = $1
    `
	rows, err := db.Query(queryAccounts, bic)
	if err != nil {
		return response, fmt.Errorf("ошибка при выполнении запроса для аккаунтов: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var account models.Account
		if err := rows.Scan(&account.Account, &account.RegulationAccountType, &account.CK, &account.AccountCBRBIC, &account.DateIn, &account.AccountStatus); err != nil {
			return response, fmt.Errorf("ошибка при сканировании аккаунта: %v", err)
		}
		response.Accounts = append(response.Accounts, account)
	}

	return response, nil
}

// Обработчик для API по BIC
func BICHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получаем BIC из URL
		bic := strings.TrimPrefix(r.URL.Path, "/bik/")
		if bic == "" {
			http.Error(w, "BIC is required", http.StatusBadRequest)
			return
		}

		entry, err := getBICEntry(db, bic)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if entry.BIC == "" {
			http.Error(w, "Запись не найдена", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(entry); err != nil {
			http.Error(w, "Ошибка при кодировании JSON", http.StatusInternalServerError)
		}
	}
}
