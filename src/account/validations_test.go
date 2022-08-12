package account

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
	"github.com/theNullP0inter/form3-api-library/src/models"
)

func TestAccountValidation(t *testing.T) {
	acc := &models.Account{}

	v := validator.New()
	v.RegisterStructValidation(AccountValidation, models.Account{})

	err := v.Struct(acc)
	assert.Equal(t, err.Error(), "Key: 'Account.attributes' Error:Field validation for 'attributes' failed on the 'attributes' tag")

	acc.Attributes = &models.AccountAttributes{}
	err = v.Struct(acc)
	assert.Equal(t,
		"Key: 'Account.country' Error:Field validation for 'country' failed on the 'supported_country' tag",
		err.Error())

	acc.Attributes.Country = "GB"
	err = v.Struct(acc)
	assert.Equal(t,
		"Key: 'Account.bank_id_code' Error:Field validation for 'bank_id_code' failed on the 'gb_bankidcode' tag\nKey: 'Account.bic' Error:Field validation for 'bic' failed on the 'gb_bic' tag\nKey: 'Account.bank_id' Error:Field validation for 'bank_id' failed on the 'gb_bankid' tag",
		err.Error())

	acc.Attributes.BankIDCode = "GBDSC"
	err = v.Struct(acc)
	assert.Equal(t,
		"Key: 'Account.bic' Error:Field validation for 'bic' failed on the 'gb_bic' tag\nKey: 'Account.bank_id' Error:Field validation for 'bank_id' failed on the 'gb_bankid' tag",
		err.Error())

	acc.Attributes.Bic = "NWBKGB22"
	err = v.Struct(acc)
	assert.Equal(t,
		"Key: 'Account.bank_id' Error:Field validation for 'bank_id' failed on the 'gb_bankid' tag",
		err.Error())

	acc.Attributes.BankID = "123"
	err = v.Struct(acc)
	assert.Equal(t,
		"Key: 'Account.bank_id' Error:Field validation for 'bank_id' failed on the 'gb_bankid_len' tag",
		err.Error())

	acc.Attributes.BankID = "123456"
	err = v.Struct(acc)
	assert.NoError(t, err)

	acc.Attributes.AccountNumber = "123"
	err = v.Struct(acc)
	assert.Equal(t,
		"Key: 'Account.account_number' Error:Field validation for 'account_number' failed on the 'gb_accountnumberlen' tag",
		err.Error())

	acc.Attributes.AccountNumber = "12345678"
	err = v.Struct(acc)
	assert.NoError(t, err)

}
