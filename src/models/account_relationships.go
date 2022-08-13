package models

type AccountRelationships struct {
	AccountEvents *Relationship `json:"account_events,omitempty"`
	MasterAccount *Relationship `json:"master_account,omitempty"`
}
