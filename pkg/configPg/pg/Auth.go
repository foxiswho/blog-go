package pg

type Auth struct {
	LoginEncrypt map[string]bool `json:"loginEncrypt" toml:"loginEncrypt" value:"${loginEncrypt}"`
}
