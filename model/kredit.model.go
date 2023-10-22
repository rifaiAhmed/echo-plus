package models

import "time"

type Consumer struct {
	NIK       string
	FullName  string
	LegalName string
	DOB       time.Time
	Salary    float64
	// KTPImage  []byte
	// Selfie    []byte
	// Limits    map[int]float64
}

type Transaction struct {
	ContractNo   string
	AssetType    string
	OTR          float64
	AdminFee     float64
	Installments int
	InterestRate float64
	AssetName    string
}

type Limit struct {
	NIK       string
	Tenor     int
	LoanLimit float64
}

type ConsumerFormatter struct {
	NIK       string
	FullName  string
	LegalName string
	DOB       time.Time
	Salary    float64
	Limit     Limit
	// KTPImage  []byte
	// Selfie    []byte
	// Limits    map[int]float64
}
