package redisPg

import (
	"fmt"

	_ "github.com/foxiswho/blog-go/pkg/cachePg/rdsPg"
	"github.com/foxiswho/blog-go/pkg/configPg/pg"
	"github.com/go-spring/spring-core/gs"
	g "github.com/redis/go-redis/v9"
)

func init() {
	gs.Provide(newClient, gs.TagArg("${pg.redis}")).
		Condition(gs.OnProperty("pg.redis.enabled").HavingValue("true").MatchIfMissing()).
		Destroy(destroyClient).
		Name("__default__")
	gs.Group("${pg.redis.instances}", newClient, nil)
}

func newClient(c pg.Redis) (*g.Client, error) {
	d, ok := driverRegistry[c.Driver]
	if !ok {
		return nil, fmt.Errorf("redis driver not found: %s", c.Driver)
	}
	return d.CreateClient(c)
}

// destroyClient closes the Redis client.
func destroyClient(client *g.Client) error {
	return client.Close()
}
