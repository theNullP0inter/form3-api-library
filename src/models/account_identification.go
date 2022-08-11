package models

type Actor struct {
	BirthDate string   `json:"birth_date,omitempty"`
	Name      []string `json:"name,omitempty"`
	Residency string   `json:"residency,omitempty"`
}

type OrganisationIdentification struct {
	Actors         []Actor  `json:"actors"`
	Address        []string `json:"address"`
	City           string   `json:"city,omitempty"`
	Country        string   `json:"country,omitempty"`
	Identification string   `json:"identification,omitempty"`
}

type PrivateIdentification struct {
	Address        []string `json:"address"`
	BirthDate      string   `json:"birth_date,omitempty"`
	BirthCountry   string   `json:"birth_country,omitempty"`
	City           string   `json:"city,omitempty"`
	Country        string   `json:"country,omitempty"`
	Identification string   `json:"identification,omitempty"`
}
