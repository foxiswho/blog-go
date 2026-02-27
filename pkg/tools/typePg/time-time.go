package typePg

import (
	"time"
)

type TimeOnly time.Time

func (t *TimeOnly) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+time.TimeOnly+`"`, string(data), time.Local)
	*t = TimeOnly(now)
	return
}

// MarshalJSON on JSONTime format Time field with Y-m-d H:i:s
func (t TimeOnly) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("null"), nil
	}
	b := make([]byte, 0, len(time.TimeOnly)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, time.TimeOnly)
	b = append(b, '"')
	return b, nil
}
func (t TimeOnly) String() string {
	return time.Time(t).Format(time.TimeOnly)
}
