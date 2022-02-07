package finances

import (
	"encoding/json"
	"fmt"
	"time"
)

func (d *Date) UnmarshalJSON(b []byte) error {
	var t time.Time

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	*d = Date(t)

	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf("\"%s\"", time.Time(d).Format(time.RFC3339Nano))
	return []byte(str), nil
}
