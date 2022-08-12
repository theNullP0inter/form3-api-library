package validations

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

	if account.Attributes.Country == "GB" {
		validateAccForGB(account, s)
	}
}

func validateAccForGB(acc models.Account, s validator.StructLevel) {

	if acc.Attributes.BankIDCode != "GBDSC" {
		s.ReportError(acc.Attributes.BankIDCode, "bank_id_code", "BankIDCode", "gbbankidcode", "")
	}

	if acc.Attributes.Bic == "" {
		s.ReportError(acc.Attributes.Bic, "bic", "Bic", "gbbic", "")
	}

	if acc.Attributes.BankID == "" {
		s.ReportError(acc.Attributes.BankID, "bank_id", "BankID", "gbbankid", "")
	}

	if acc.Attributes.BankID != "" && len(acc.Attributes.BankID) != 6 {
		s.ReportError(acc.Attributes.BankID, "bank_id", "BankID", "gbbankidlen", "")
	}

	if acc.Attributes.AccountNumber != "" && len(acc.Attributes.AccountNumber) != 8 {
		s.ReportError(acc.Attributes.AccountNumber, "account_number", "AccountNumber", "gbaccountnumberlen", "")
	}

}
