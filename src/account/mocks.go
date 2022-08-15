package account

import "github.com/theNullP0inter/form3-api-library/src/models"

var MockAccount = models.Account{
	Type: "accounts",
	Attributes: &models.AccountAttributes{
		Country:       "GB",
		BankIDCode:    "GBDSC",
		Bic:           "NWBKGB22",
		BankID:        "123456",
		AccountNumber: "12345678",
		Name:          []string{"Jon", "Snow"},
	},
}

var MockAccountJson = `{"data":{"type":"accounts","id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","version":0,"attributes":{"country":"GB","base_currency":"GBP","account_number":"41426819","bank_id":"400300","bank_id_code":"GBDSC","bic":"NWBKGB22","iban":"GB11NWBK40030041426819","name":["Samantha Holder"],"alternative_names":["Sam Holder"],"account_classification":"Personal","joint_account":false,"account_matching_opt_out":false,"secondary_identification":"A1B2C3D4","switched":false,"name_matching_status":"supported","private_identification":{"birth_date":"2017-07-23","birth_country":"GB","identification":"13YH458762","address":["10 Avenue des Champs"],"city":"London","country":"GB"},"organisation_identification":{"identification":"123654","actors":[{"name":["Jeff Page"],"birth_date":"1970-01-01","residency":"GB"}],"address":["10 Avenue des Champs"],"city":"London","country":"GB"},"status":"confirmed","status_reason":"unspecified","user_defined_data":[{"key":"Some account related key","value":"Some account related value"}],"validation_type":"card","reference_mask":"############","acceptance_qualifier":"same_day"},"relationships":{"master_account":{"data":[{"type":"accounts","id":"a52d13a4-f435-4c00-cfad-f5e7ac5972df"}]},"account_events":{"data":[{"type":"account_events","id":"c1023677-70ee-417a-9a6a-e211241f1e9c"},{"type":"account_events","id":"437284fa-62a6-4f1d-893d-2959c9780288"}]}}}}`
var MockAccountListJson = `{"data":[{"type":"accounts","id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","version":0,"attributes":{"country":"GB","base_currency":"GBP","account_number":"41426819","bank_id":"400300","bank_id_code":"GBDSC","bic":"NWBKGB22","iban":"GB11NWBK40030041426819","account_classification":"Personal","joint_account":false,"switched":false,"account_matching_opt_out":false,"status":"confirmed"}},{"type":"accounts","id":"ea6239c1-99e9-42b3-bca1-92f5c068da6b","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","version":0,"attributes":{"country":"GB","base_currency":"GBP","account_number":"76357283","bank_id":"400305","bank_id_code":"GBDSC","bic":"LHVBEE22","iban":"EE209642394169415743","account_classification":"Personal","joint_account":false,"switched":false,"account_matching_opt_out":false,"status":"confirmed"}}]}`

var MockAccountJsonID = "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
