package typePg

import (
	"fmt"
	"strconv"
	"strings"
)

type Int32String int32

func (i Int32String) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%v\"", i)), nil
}

func (i *Int32String) UnmarshalJSON(value []byte) error {
	str := string(value)
	if strings.LastIndexAny(str, "\"") > -1 {
		str = string(value[1 : len(value)-1])
	}
	m, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return err
	}
	*i = Int32String(m)
	return nil
}

func (i Int32String) ToString() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int32String) ToInt32() int32 {
	return int32(i)
}

type Int32 int32

func (i Int32) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", i)), nil
}

func (i *Int32) UnmarshalJSON(value []byte) error {
	str := string(value)
	if strings.LastIndexAny(str, "\"") > -1 {
		str = string(value[1 : len(value)-1])
	}
	m, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return err
	}
	*i = Int32(m)
	return nil
}

func (i Int32) ToString() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int32) ToInt32() int32 {
	return int32(i)
}

func (i Int32) ToInt64() int64 {
	return int64(i)
}
