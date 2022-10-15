package models

// Weather ...
type Weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Data struct {
	Status Weather `json:"status"`
}
