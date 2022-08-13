package models

type AccountAttributes struct {
	AccountClassification   string   `json:"account_classification,omitempty"`
	AccountMatchingOptOut   bool     `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty" validate:"max=11"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 string   `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            bool     `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty" validate:"required"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty" validate:"max=140"`
	Status                  string   `json:"status,omitempty"`
	Switched                bool     `json:"switched,omitempty"`
}
