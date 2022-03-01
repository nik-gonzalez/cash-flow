package models

import (
	"encoding/json"
	"time"
)

type CashFlow struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Remarks string  `json:"remarks"`
	Amount  float64 `json:"amount"`
	Date    JTime   `json:"date"`
}

type JTime time.Time

func (mt *JTime) UnmarshalJSON(bs []byte) error {
	var timestamp string
	err := json.Unmarshal(bs, &timestamp)
	if err != nil {
		return err
	}
	parse, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return err
	}
	*mt = JTime(parse)
	return nil
}

func (mt JTime) MarshalJSON() ([]byte, error) {
	timestamp := time.Time(mt).Format(time.RFC3339)
	return json.Marshal(timestamp)
}
