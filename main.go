package main

import (
	"database/sql"
	"goLen/parser"
	"goLen/server"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib" // драйвер для PostgreSQL
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите команду: 'parse' или 'server'")
	}

	db, err := sql.Open("pgx", "postgres:")
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer db.Close()

	switch os.Args[1] {
	case "parse":
		if len(os.Args) < 3 {
			log.Fatal("Укажите путь к XML файлу")
		}
		filePath := os.Args[2]
		ed, err := parser.ParseXMLFile(filePath)
		if err != nil {
			log.Fatalf("Ошибка при парсинге XML: %v", err)
		}

		for _, entry := range ed.BICDirectoryEntry {
			if err := parser.InsertIntoDB(db, entry); err != nil {
				log.Printf("Ошибка при вставке записи BIC %s: %v", entry.BIC, err)
			}
		}
		log.Println("Все данные успешно сохранены в базу данных.")

	case "server":
		http.HandleFunc("/bik/", server.BICHandler(db)) // Изменен путь на /bik/
		log.Println("Сервер запущен на порту 8080")
		log.Fatal(http.ListenAndServe(":8080", nil))

	default:
		log.Fatal("Неверная команда. Используйте 'parse' или 'server'")
	}
}

//package main
//
//import (
//
//	"database/sql"
//	"encoding/xml"
//	"fmt"
//	"io"
//	"log"
//	"os"
//
//	"golang.org/x/text/encoding/charmap"
//
//	_ "github.com/jackc/pgx/v4/stdlib" // драйвер для PostgreSQL
//
//)
//
//	type ED807 struct {
//		XMLName           xml.Name            `xml:"ED807"`
//		EDNo              string              `xml:"EDNo,attr"`
//		EDDate            string              `xml:"EDDate,attr"`
//		EDAuthor          string              `xml:"EDAuthor,attr"`
//		CreationReason    string              `xml:"CreationReason,attr"`
//		CreationDateTime  string              `xml:"CreationDateTime,attr"`
//		InfoTypeCode      string              `xml:"InfoTypeCode,attr"`
//		BusinessDay       string              `xml:"BusinessDay,attr"`
//		DirectoryVersion  string              `xml:"DirectoryVersion,attr"`
//		BICDirectoryEntry []BICDirectoryEntry `xml:"BICDirectoryEntry"`
//	}
//
//	type BICDirectoryEntry struct {
//		BIC             string          `xml:"BIC,attr"`
//		ParticipantInfo ParticipantInfo `xml:"ParticipantInfo"`
//		Accounts        []Account       `xml:"Accounts"`
//		SWBICS          []SWBIC         `xml:"SWBICS"`
//	}
//
//	type ParticipantInfo struct {
//		NameP             string `xml:"NameP,attr"`
//		EnglName          string `xml:"EnglName,attr"`
//		RegN              string `xml:"RegN,attr"`
//		CntrCd            string `xml:"CntrCd,attr"`
//		Rgn               string `xml:"Rgn,attr"`
//		Ind               string `xml:"Ind,attr"`
//		Tnp               string `xml:"Tnp,attr"`
//		Nnp               string `xml:"Nnp,attr"`
//		Adr               string `xml:"Adr,attr"`
//		PrntBIC           string `xml:"PrntBIC,attr,omitempty"`
//		DateIn            string `xml:"DateIn,attr"`
//		PtType            string `xml:"PtType,attr"`
//		Srvcs             string `xml:"Srvcs,attr"`
//		XchType           string `xml:"XchType,attr"`
//		UID               string `xml:"UID,attr"`
//		ParticipantStatus string `xml:"ParticipantStatus,attr"`
//	}
//
//	type Account struct {
//		Account               string `xml:"Account,attr"`
//		RegulationAccountType string `xml:"RegulationAccountType,attr"`
//		CK                    string `xml:"CK,attr"`
//		AccountCBRBIC         string `xml:"AccountCBRBIC,attr"`
//		DateIn                string `xml:"DateIn,attr"`
//		AccountStatus         string `xml:"AccountStatus,attr"`
//	}
//
//	type SWBIC struct {
//		SWBIC        string `xml:"SWBIC,attr"`
//		DefaultSWBIC string `xml:"DefaultSWBIC,attr"`
//	}
//
//// Функция для записи данных в PostgreSQL
//
//	func insertIntoDB(db *sql.DB, entry BICDirectoryEntry) error {
//		query := `
//	     INSERT INTO bic_directory_entries (bic, namep, engl_name, regn, cntr_cd, rgn, ind, tnp, nnp, adr, prnt_bic, date_in, pt_type, srvcs, xch_type, uid, participant_status)
//	     VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
//	 `
//		_, err := db.Exec(query, entry.BIC, entry.ParticipantInfo.NameP, entry.ParticipantInfo.EnglName, entry.ParticipantInfo.RegN, entry.ParticipantInfo.CntrCd, entry.ParticipantInfo.Rgn,
//			entry.ParticipantInfo.Ind, entry.ParticipantInfo.Tnp, entry.ParticipantInfo.Nnp, entry.ParticipantInfo.Adr, entry.ParticipantInfo.PrntBIC,
//			entry.ParticipantInfo.DateIn, entry.ParticipantInfo.PtType, entry.ParticipantInfo.Srvcs, entry.ParticipantInfo.XchType, entry.ParticipantInfo.UID, entry.ParticipantInfo.ParticipantStatus)
//
//		if err != nil {
//			return fmt.Errorf("не удалось вставить данные: %v", err)
//		}
//
//		return nil
//	}
//
//	func main() {
//		// 1. Читаем XML файл
//		xmlFile, err := os.Open("20240923_ED807_full.xml") // Убедитесь, что файл находится на одном уровне с main.go
//		if err != nil {
//			log.Fatalf("Не удалось открыть файл: %v", err)
//		}
//		defer xmlFile.Close()
//
//		// 2. Создаем новый декодер XML с кастомным CharsetReader
//		d := xml.NewDecoder(xmlFile)
//		d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
//			switch charset {
//			case "windows-1251", "Windows-1251": // добавляем обработку для обоих вариантов
//				return charmap.Windows1251.NewDecoder().Reader(input), nil
//			default:
//				return nil, fmt.Errorf("unknown charset: %s", charset)
//			}
//		}
//
//		// 3. Парсим XML данные
//		var ed ED807
//		err = d.Decode(&ed)
//		if err != nil {
//			log.Fatalf("Ошибка при парсинге XML: %v", err)
//		}
//
//		// 4. Подключаемся к базе данных PostgreSQL
//		db, err := sql.Open("pgx", "postgres:")
//		if err != nil {
//			log.Fatalf("Не удалось подключиться к базе данных: %v", err)
//		}
//		defer db.Close()
//
//		// 5. Записываем данные в базу данных
//		for _, entry := range ed.BICDirectoryEntry {
//			err := insertIntoDB(db, entry)
//			if err != nil {
//				log.Printf("Ошибка при вставке записи BIC %s: %v", entry.BIC, err)
//			}
//		}
//
//		log.Println("Все данные успешно сохранены в базу данных.")
//	}
//package main
//
//import (
//	"database/sql"
//	"encoding/xml"
//	"fmt"
//	"io"
//	"log"
//	"os"
//	"time"
//
//	"golang.org/x/text/encoding/charmap"
//
//	_ "github.com/jackc/pgx/v4/stdlib" // драйвер для PostgreSQL
//)
//
//type ED807 struct {
//	XMLName           xml.Name            `xml:"ED807"`
//	EDNo              string              `xml:"EDNo,attr"`
//	EDDate            string              `xml:"EDDate,attr"`
//	EDAuthor          string              `xml:"EDAuthor,attr"`
//	CreationReason    string              `xml:"CreationReason,attr"`
//	CreationDateTime  string              `xml:"CreationDateTime,attr"`
//	InfoTypeCode      string              `xml:"InfoTypeCode,attr"`
//	BusinessDay       string              `xml:"BusinessDay,attr"`
//	DirectoryVersion  string              `xml:"DirectoryVersion,attr"`
//	BICDirectoryEntry []BICDirectoryEntry `xml:"BICDirectoryEntry"`
//}
//
//type BICDirectoryEntry struct {
//	BIC             string          `xml:"BIC,attr"`
//	ParticipantInfo ParticipantInfo `xml:"ParticipantInfo"`
//	Accounts        Account         `xml:"Accounts"`
//}
//
//type ParticipantInfo struct {
//	NameP             string `xml:"NameP,attr"`
//	EnglName          string `xml:"EnglName,attr,omitempty"`
//	RegN              string `xml:"RegN,attr"`
//	CntrCd            string `xml:"CntrCd,attr"`
//	Rgn               string `xml:"Rgn,attr"`
//	Ind               string `xml:"Ind,attr"`
//	Tnp               string `xml:"Tnp,attr"`
//	Nnp               string `xml:"Nnp,attr"`
//	Adr               string `xml:"Adr,attr"`
//	PrntBIC           string `xml:"PrntBIC,attr,omitempty"`
//	DateIn            string `xml:"DateIn,attr"`
//	PtType            string `xml:"PtType,attr"`
//	Srvcs             string `xml:"Srvcs,attr"`
//	XchType           string `xml:"XchType,attr"`
//	UID               string `xml:"UID,attr"`
//	ParticipantStatus string `xml:"ParticipantStatus,attr"`
//}
//
//type Account struct {
//	Account               string `xml:"Account,attr"`
//	RegulationAccountType string `xml:"RegulationAccountType,attr"`
//	CK                    string `xml:"CK,attr"`
//	AccountCBRBIC         string `xml:"AccountCBRBIC,attr"`
//	DateIn                string `xml:"DateIn,attr"`
//	AccountStatus         string `xml:"AccountStatus,attr"`
//}
//
//// Функция для создания таблицы
//func createTable(db *sql.DB) error {
//	query := `
//    CREATE TABLE IF NOT EXISTS bic_accounts (
//        id SERIAL PRIMARY KEY,
//        bic VARCHAR(20) NOT NULL,
//        namep TEXT,
//        engl_name TEXT,
//        regn VARCHAR(20),
//        cntr_cd VARCHAR(5),
//        rgn VARCHAR(5),
//        ind VARCHAR(20),
//        tnp VARCHAR(10),
//        nnp VARCHAR(100),
//        adr TEXT,
//        prnt_bic VARCHAR(20),
//        date_in DATE,
//        pt_type VARCHAR(10),
//        srvcs VARCHAR(10),
//        xch_type VARCHAR(10),
//        uid VARCHAR(20),
//        participant_status VARCHAR(10),
//        account VARCHAR(34) NOT NULL,
//        regulation_account_type VARCHAR(10),
//        ck VARCHAR(10),
//        account_cbr_bic VARCHAR(20),
//        account_date_in DATE,
//        account_status VARCHAR(10)
//    );`
//
//	_, err := db.Exec(query)
//	if err != nil {
//		return fmt.Errorf("не удалось создать таблицу: %v", err)
//	}
//	return nil
//}
//
//// Функция для вставки данных в таблицу
//func insertIntoDB(db *sql.DB, entry BICDirectoryEntry) error {
//	// Преобразуем строки дат в тип date
//	dateIn, err := parseDate(entry.ParticipantInfo.DateIn)
//	if err != nil {
//		return fmt.Errorf("ошибка при преобразовании даты участника: %v", err)
//	}
//
//	accountDateIn, err := parseDate(entry.Accounts.DateIn)
//	if err != nil {
//		return fmt.Errorf("ошибка при преобразовании даты счета: %v", err)
//	}
//
//	query := `
//       INSERT INTO bic_accounts (
//           bic, namep, engl_name, regn, cntr_cd, rgn, ind, tnp, nnp, adr, prnt_bic, date_in, pt_type, srvcs, xch_type, uid, participant_status,
//           account, regulation_account_type, ck, account_cbr_bic, account_date_in, account_status
//       )
//       VALUES (
//           $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11,
//           $12, $13, $14, $15, $16, $17,
//           $18, $19, $20, $21, $22, $23
//       )
//   `
//	_, err = db.Exec(query, entry.BIC, entry.ParticipantInfo.NameP, entry.ParticipantInfo.EnglName, entry.ParticipantInfo.RegN, entry.ParticipantInfo.CntrCd, entry.ParticipantInfo.Rgn,
//		entry.ParticipantInfo.Ind, entry.ParticipantInfo.Tnp, entry.ParticipantInfo.Nnp, entry.ParticipantInfo.Adr, entry.ParticipantInfo.PrntBIC,
//		dateIn, entry.ParticipantInfo.PtType, entry.ParticipantInfo.Srvcs, entry.ParticipantInfo.XchType, entry.ParticipantInfo.UID, entry.ParticipantInfo.ParticipantStatus,
//		entry.Accounts.Account, entry.Accounts.RegulationAccountType, entry.Accounts.CK, entry.Accounts.AccountCBRBIC, accountDateIn, entry.Accounts.AccountStatus)
//
//	if err != nil {
//		return fmt.Errorf("не удалось вставить данные: %v", err)
//	}
//
//	return nil
//}
//
//// Функция для преобразования строкового представления даты в тип date
//func parseDate(dateStr string) (sql.NullString, error) {
//	if dateStr == "" {
//		return sql.NullString{Valid: false}, nil
//	}
//
//	parsedDate, err := time.Parse("2006-01-02", dateStr)
//	if err != nil {
//		return sql.NullString{}, fmt.Errorf("недопустимый формат даты: %s", dateStr)
//	}
//
//	return sql.NullString{String: parsedDate.Format("2006-01-02"), Valid: true}, nil
//}
//
//func main() {
//	// 1. Читаем XML файл
//	xmlFile, err := os.Open("20240923_ED807_full.xml") // Убедитесь, что файл находится на одном уровне с main.go
//	if err != nil {
//		log.Fatalf("Не удалось открыть файл: %v", err)
//	}
//	defer xmlFile.Close()
//
//	// 2. Создаем новый декодер XML с кастомным CharsetReader
//	d := xml.NewDecoder(xmlFile)
//	d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
//		switch charset {
//		case "windows-1251", "Windows-1251": // добавляем обработку для обоих вариантов
//			return charmap.Windows1251.NewDecoder().Reader(input), nil
//		default:
//			return nil, fmt.Errorf("unknown charset: %s", charset)
//		}
//	}
//
//	// 3. Парсим XML данные
//	var ed ED807
//	err = d.Decode(&ed)
//	if err != nil {
//		log.Fatalf("Ошибка при парсинге XML: %v", err)
//	}
//
//	// 4. Подключаемся к базе данных PostgreSQL
//	db, err := sql.Open("pgx", "postgres:")
//	if err != nil {
//		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
//	}
//	defer db.Close()
//
//	// 5. Создаем таблицу
//	if err := createTable(db); err != nil {
//		log.Fatalf("Ошибка при создании таблицы: %v", err)
//	}
//
//	// 6. Записываем данные в базу данных
//	for _, entry := range ed.BICDirectoryEntry {
//		err := insertIntoDB(db, entry)
//		if err != nil {
//			log.Printf("Ошибка при вставке записи BIC %s: %v", entry.BIC, err)
//		}
//	}
//
//	log.Println("Все данные успешно сохранены в базу данных.")
//}
