// package parser
//
// import (
//
//	"database/sql"
//	"encoding/xml"
//	"fmt"
//	"goLen/models"
//	"io"
//	_ "log"
//	"os"
//
//	"golang.org/x/text/encoding/charmap"
//
// )
//
// // Функция для вставки данных в PostgreSQL
//
//	func InsertIntoDB(db *sql.DB, entry models.BICDirectoryEntry) error {
//		query := `
//	       INSERT INTO bic_directory_entries (bic, namep, engl_name, regn, cntr_cd, rgn, ind, tnp, nnp, adr, prnt_bic, date_in, pt_type, srvcs, xch_type, uid, participant_status)
//	       VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
//	   `
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
// // Функция для парсинга XML файла
//
//	func ParseXMLFile(filePath string) (models.ED807, error) {
//		xmlFile, err := os.Open(filePath)
//		if err != nil {
//			return models.ED807{}, fmt.Errorf("не удалось открыть файл: %v", err)
//		}
//		defer xmlFile.Close()
//
//		d := xml.NewDecoder(xmlFile)
//		d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
//			switch charset {
//			case "windows-1251", "Windows-1251":
//				return charmap.Windows1251.NewDecoder().Reader(input), nil
//			default:
//				return nil, fmt.Errorf("unknown charset: %s", charset)
//			}
//		}
//
//		var ed models.ED807
//		err = d.Decode(&ed)
//		if err != nil {
//			return models.ED807{}, fmt.Errorf("Ошибка при парсинге XML: %v", err)
//		}
//
//		return ed, nil
//	}
package parser

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"goLen/models"
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

// Функция для вставки данных в PostgreSQL
func InsertIntoDB(db *sql.DB, entry models.BICDirectoryEntry) error {
	query := `
        INSERT INTO bic_directory_entries (bic, namep, engl_name, regn, cntr_cd, rgn, ind, tnp, nnp, adr, prnt_bic, date_in, pt_type, srvcs, xch_type, uid, participant_status)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
    `
	_, err := db.Exec(query, entry.BIC, entry.ParticipantInfo.NameP, entry.ParticipantInfo.EnglName, entry.ParticipantInfo.RegN, entry.ParticipantInfo.CntrCd, entry.ParticipantInfo.Rgn,
		entry.ParticipantInfo.Ind, entry.ParticipantInfo.Tnp, entry.ParticipantInfo.Nnp, entry.ParticipantInfo.Adr, entry.ParticipantInfo.PrntBIC,
		entry.ParticipantInfo.DateIn, entry.ParticipantInfo.PtType, entry.ParticipantInfo.Srvcs, entry.ParticipantInfo.XchType, entry.ParticipantInfo.UID, entry.ParticipantInfo.ParticipantStatus)

	if err != nil {
		return fmt.Errorf("не удалось вставить данные: %v", err)
	}

	return nil
}

// Функция для парсинга XML файла
func ParseXMLFile(filePath string) (models.ED807, error) {
	xmlFile, err := os.Open(filePath)
	if err != nil {
		return models.ED807{}, fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer xmlFile.Close()

	d := xml.NewDecoder(xmlFile)
	d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251", "Windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}

	var ed models.ED807
	err = d.Decode(&ed)
	if err != nil {
		return models.ED807{}, fmt.Errorf("Ошибка при парсинге XML: %v", err)
	}

	return ed, nil
}
