package url

import (
	"net/url"
	"strconv"
)

type Url url.Values
//数值
func (t Url)GetInt(key string, def ...int) (int) {
	if t == nil {
		if len(def) > 0 {
			return def[0]
		} else {
			return 0
		}
	}
	strv := t[key]
	if len(strv) == 0 && len(def) > 0 {
		return def[0]
	}
	v, _ := strconv.Atoi(strv)
	return v
}