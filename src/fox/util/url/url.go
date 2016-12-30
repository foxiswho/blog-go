package url

import (
	"net/url"
	"strconv"
)

type Url struct {
	url.Values
}
//数值
func (t Url)GetInt(key string, def ...int) (int) {
	if t.Values == nil {
		if len(def) > 0 {
			return def[0]
		} else {
			return 0
		}
	}

	strv :=t.Get(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0]
	}

	v, _ := strconv.Atoi(strv)
	return v
}