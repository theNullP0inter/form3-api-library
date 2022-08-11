package models

type RelatedItem struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Relationship struct {
	Data []RelatedItem `json:"data"`
}
