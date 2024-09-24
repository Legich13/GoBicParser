//package models
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
//type SWBIC struct {
//	SWBIC        string `xml:"SWBIC,attr"`
//	DefaultSWBIC string `xml:"DefaultSWBIC,attr"`
//}
//
//type BICResponse struct {
//	BIC   string `json:"bic"`
//	NameP string `json:"namep"`
//}
//
//type BICDirectoryEntry struct {
//	BIC             string          `xml:"BIC,attr"`
//	ParticipantInfo ParticipantInfo `xml:"ParticipantInfo"`
//	Accounts        []Account       `xml:"Accounts"`
//	SWBICS          []SWBIC         `xml:"SWBICS"`
//}
//
//type ParticipantInfo struct {
//	NameP             string  `xml:"NameP,attr"`
//	EnglName          *string `xml:"EnglName,attr,omitempty"`
//	RegN              *string `xml:"RegN,attr,omitempty"`
//	CntrCd            *string `xml:"CntrCd,attr,omitempty"`
//	Rgn               *string `xml:"Rgn,attr,omitempty"`
//	Ind               *string `xml:"Ind,attr,omitempty"`
//	Tnp               *string `xml:"Tnp,attr,omitempty"`
//	Nnp               *string `xml:"Nnp,attr,omitempty"`
//	Adr               *string `xml:"Adr,attr,omitempty"`
//	PrntBIC           *string `xml:"PrntBIC,attr,omitempty"`
//	DateIn            *string `xml:"DateIn,attr,omitempty"`
//	PtType            *string `xml:"PtType,attr,omitempty"`
//	Srvcs             *string `xml:"Srvcs,attr,omitempty"`
//	XchType           *string `xml:"XchType,attr,omitempty"`
//	UID               *string `xml:"UID,attr,omitempty"`
//	ParticipantStatus *string `xml:"ParticipantStatus,attr,omitempty"`
//}
//
//type ED807 struct {
//	BICDirectoryEntry []BICDirectoryEntry `xml:"BICDirectoryEntry"`
//}

package models

type Account struct {
	Account               string  `xml:"Account,attr"`
	RegulationAccountType string  `xml:"RegulationAccountType,attr"`
	CK                    string  `xml:"CK,attr"`
	AccountCBRBIC         string  `xml:"AccountCBRBIC,attr"`
	DateIn                *string `xml:"DateIn,attr,omitempty"`
	AccountStatus         string  `xml:"AccountStatus,attr"`
}

type SWBIC struct {
	SWBIC        string `xml:"SWBIC,attr"`
	DefaultSWBIC string `xml:"DefaultSWBIC,attr"`
}

type BICDirectoryEntry struct {
	BIC             string          `xml:"BIC,attr"`
	ParticipantInfo ParticipantInfo `xml:"ParticipantInfo"`
	Accounts        []Account       `xml:"Accounts"`
	SWBICS          []SWBIC         `xml:"SWBICS"`
}

type ParticipantInfo struct {
	NameP             string  `xml:"NameP,attr"`
	EnglName          *string `xml:"EnglName,attr,omitempty"`
	RegN              *string `xml:"RegN,attr,omitempty"`
	CntrCd            *string `xml:"CntrCd,attr,omitempty"`
	Rgn               *string `xml:"Rgn,attr,omitempty"`
	Ind               *string `xml:"Ind,attr,omitempty"`
	Tnp               *string `xml:"Tnp,attr,omitempty"`
	Nnp               *string `xml:"Nnp,attr,omitempty"`
	Adr               *string `xml:"Adr,attr,omitempty"`
	PrntBIC           *string `xml:"PrntBIC,attr,omitempty"`
	DateIn            *string `xml:"DateIn,attr,omitempty"`
	PtType            *string `xml:"PtType,attr,omitempty"`
	Srvcs             *string `xml:"Srvcs,attr,omitempty"`
	XchType           *string `xml:"XchType,attr,omitempty"`
	UID               *string `xml:"UID,attr,omitempty"`
	ParticipantStatus *string `xml:"ParticipantStatus,attr,omitempty"`
}

type ED807 struct {
	BICDirectoryEntry []BICDirectoryEntry `xml:"BICDirectoryEntry"`
}

type BICResponse struct {
	BIC      string    `json:"bic"`
	NameP    string    `json:"namep"`
	Accounts []Account `json:"accounts"` // Добавляем поле для аккаунтов
}
