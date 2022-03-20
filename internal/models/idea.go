package models

type Idea struct {
	Id            int    `json:"id"`
	EstimatedTime int    `json:"estimatedTime"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	CreatedDate   string `json:"createdData"`
}
