package pg

import "strconv"

// Redis app.redis
type Redis struct {
	Password       string   `json:"password" value:"${password:=}"`
	Username       string   `json:"username" value:"${username:=}"`
	Addrs          []string `json:"addrs" value:"${addrs:=}"  `
	Host           string   `json:"host" value:"${host:=127.0.0.1}" `
	Port           int      `json:"port" value:"${port:=6379}"`
	Database       int      `value:"${database:=0}"`
	Ping           bool     `value:"${ping:=true}"`
	IdleTimeout    int      `value:"${idle-timeout:=0}"`
	ConnectTimeout int      `value:"${connect-timeout:=0}"`
	ReadTimeout    int      `value:"${read-timeout:=0}"`
	WriteTimeout   int      `value:"${write-timeout:=0}"`
}

func (r Redis) GetAddress() string {
	return r.Host + ":" + r.PortToString()
}

func (r Redis) PortToString() string {
	return strconv.Itoa(r.Port)
}
