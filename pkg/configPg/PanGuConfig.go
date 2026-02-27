package configPg

type PanGuConfig struct {
	Pg       Pg       `json:"pg" value:"${pg}"`
	Database Database `json:"database" value:"${database}"`
	Server   Server   `json:"serverPg" value:"${serverPg}"`
}
