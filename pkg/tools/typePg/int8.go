package typePg

import (
	"fmt"
	"strconv"
	"strings"
)

type Int8 int8

func (i Int8) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", i)), nil
}

func (i *Int8) UnmarshalJSON(value []byte) error {
	str := string(value)
	if strings.LastIndexAny(str, "\"") > -1 {
		str = string(value[1 : len(value)-1])
	}
	fmt.Printf("%+v\n", str)
	if 0 == len(str) {
		*i = 0
		return nil
	}
	if "null" == str {
		*i = 0
		return nil
	}
	m, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return err
	}
	*i = Int8(m)
	return nil
}

func (i Int8) ToString() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int8) ToInt8() int8 {
	return int8(i)
}
