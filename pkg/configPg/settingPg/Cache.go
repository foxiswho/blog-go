package settingPg

var caches = NewConfigs()

// ConfigsData
// @Description: 缓存
type ConfigsData struct {
	data            map[string]any  //缓存数据
	loading         map[string]bool //加载状态
	loadingComplete bool            //加载完成
}

func NewConfigs() *ConfigsData {
	c := &ConfigsData{}
	c.init()
	return c
}

func (c *ConfigsData) init() {
	c.data = make(map[string]any)
	c.loadingComplete = false
}

func (c *ConfigsData) Get(key string) (any, bool) {
	a, ok := c.data[key]
	return a, ok
}

func (c *ConfigsData) Set(key string, val any) bool {
	c.data[key] = val
	return true
}

func (c *ConfigsData) Clear(key string) bool {
	delete(c.data, key)
	return true
}

func Get(key string) (any, bool) {
	return caches.Get(key)
}

func Set(key string, val any) bool {
	return caches.Set(key, val)
}

func Clear(key string) bool {
	return caches.Clear(key)
}
