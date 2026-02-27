package typePg

import (
	"time"
)

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+time.DateTime+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

// MarshalJSON on JSONTime format Time field with Y-m-d H:i:s
func (t Time) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("null"), nil
	}
	b := make([]byte, 0, len(time.DateTime)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, time.DateTime)
	b = append(b, '"')
	return b, nil
}

// Value insert timestamp into mysql need this function.
//func (t Time) Value() (driver.Value, error) {
//	var zeroTime time.Time
//	if time.Time(t).UnixNano() == zeroTime.UnixNano() {
//		return nil, nil
//	}
//	return time.Time(t), nil
//}

func (t Time) String() string {
	return time.Time(t).Format(time.DateTime)
}

func (t Time) ToTime() time.Time {
	return time.Time(t)
}

func (t *Time) FormTimePointer(val *time.Time) *Time {
	if val != nil {
		*t = Time(*val)
	}
	return t
}

func (t *Time) FormTime(val time.Time) *Time {
	*t = Time(val)
	return t
}

func TimeFormTimePointer(val *time.Time) *Time {
	v := new(Time)
	return v.FormTimePointer(val)
}

func TimeFormTime(val time.Time) *Time {
	v := Time(val)
	return &v
}

//func (t *Time) Scan(v interface{}) error {
//	value, ok := v.(time.Time)
//	if ok {
//		*t = Time(value)
//		return nil
//	}
//	return fmt.Errorf("can not convert %v to timestamp", v)
//}
