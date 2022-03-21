package fbaOutbound

import (
	"encoding/json"
	"fmt"
	"time"
)

func (d *Timestamp) UnmarshalJSON(b []byte) error {
	var t time.Time

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	*d = Timestamp(t)

	return nil
}

func (d Timestamp) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf("\"%s\"", time.Time(d).Format(time.RFC3339Nano))
	return []byte(str), nil
}
