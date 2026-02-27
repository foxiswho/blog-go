package typePg

import (
	"time"
)

type DateOnly time.Time

func (t *DateOnly) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+time.DateOnly+`"`, string(data), time.Local)
	*t = DateOnly(now)
	return
}

// MarshalJSON on JSONTime format Time field with Y-m-d H:i:s
func (t DateOnly) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("null"), nil
	}
	b := make([]byte, 0, len(time.DateOnly)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, time.DateOnly)
	b = append(b, '"')
	return b, nil
}
func (t DateOnly) String() string {
	return time.Time(t).Format(time.DateOnly)
}
