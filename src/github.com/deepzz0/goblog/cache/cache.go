package cache

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/deepzz0/go-com/log"
	"github.com/deepzz0/goblog/models"
)

type cache struct {
	BackgroundLeftBar  map[string]string
	BackgroundLeftBars []*models.Leftbar
	BuildVersion       string
}

var Cache = NewCache()

func NewCache() *cache {
	return &cache{BackgroundLeftBar: make(map[string]string)}
}

func init() {
	doReadBackLeftBarConfig()
	doReadBuildVersionConfig()
}

var path, _ = os.Getwd()

func doReadBackLeftBarConfig() {
	b, err := ioutil.ReadFile(path + "/conf/backleft.conf")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &Cache.BackgroundLeftBars)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range Cache.BackgroundLeftBars {
		if v.ID != "" {
			Cache.BackgroundLeftBar[v.ID] = v.ID
		}
	}
}

func doReadBuildVersionConfig() {
	b, err := ioutil.ReadFile(path + "/version")
	if err != nil {
		log.Error(err)
	}
	Cache.BuildVersion = string(b)
}
