package customType

import (
	"encoding/json"
	"time"
)

type JsonDateTime struct {
	time.Time
}

func (mt *JsonDateTime) UnmarshalJSON(bs []byte) error {
	var timestamp string
	err := json.Unmarshal(bs, &timestamp)
	if err != nil {
		return err
	}
	parsed, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return err
	}
	*mt = JsonDateTime{parsed}
	return nil
}

func (mt JsonDateTime) MarshalJSON() ([]byte, error) {
	timestamp := mt.Time.Format(time.RFC3339)
	return json.Marshal(timestamp)
}
