package account

import (
	"github.com/go-playground/validator"
	"github.com/theNullP0inter/form3-api-library/src/models"
)

func AccountValidation(s validator.StructLevel) {

	account := s.Current().Interface().(models.Account)

	if account.Attributes == nil {
		s.ReportError(account.Attributes, "attributes", "Attributes", "attributes", "")
		return
	}

	countryValidation := map[string]func(acc models.Account, s validator.StructLevel){
		"GB": validateAccForGB,
	}

	if countryValidator, ok := countryValidation[account.Attributes.Country]; ok {
		countryValidator(account, s)
	} else {
		s.ReportError(account.Attributes.Country, "country", "Country", "supported_country", "")
	}

}

func validateAccForGB(acc models.Account, s validator.StructLevel) {

	if acc.Attributes.BankIDCode != "GBDSC" {
		s.ReportError(acc.Attributes.BankIDCode, "bank_id_code", "BankIDCode", "gb_bankidcode", "")
	}

	if acc.Attributes.Bic == "" {
		s.ReportError(acc.Attributes.Bic, "bic", "Bic", "gb_bic", "")
	}

	if acc.Attributes.BankID == "" {
		s.ReportError(acc.Attributes.BankID, "bank_id", "BankID", "gb_bankid", "")
	}

	if acc.Attributes.BankID != "" && len(acc.Attributes.BankID) != 6 {
		s.ReportError(acc.Attributes.BankID, "bank_id", "BankID", "gb_bankid_len", "")
	}

	if acc.Attributes.AccountNumber != "" && len(acc.Attributes.AccountNumber) != 8 {
		s.ReportError(acc.Attributes.AccountNumber, "account_number", "AccountNumber", "gb_accountnumberlen", "")
	}

}
