package zdb

import (
	"database/sql/driver"
	"fmt"
	"github.com/zohu/zutils"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" {
		return nil
	}
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse(zutils.FormatTime, timeStr)
	*t = Time{t1}
	return err
}

func (t *Time) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *Time) GetTime() time.Time {
	if t == nil {
		return time.Time{}
	} else {
		return t.Time
	}
}
