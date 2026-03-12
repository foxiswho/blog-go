package eventBasicEvent

import (
	"context"

	"github.com/foxiswho/blog-go/app/core/basic/model/modCacheBasicEvent"
	"github.com/foxiswho/blog-go/app/event/basic/model/modEventBasicEvent"
	"github.com/foxiswho/blog-go/infrastructure/entityBasic"
	"github.com/foxiswho/blog-go/pkg/enum/state/enumStatePg"
	"github.com/foxiswho/blog-go/pkg/log2"
	"github.com/foxiswho/blog-go/pkg/sdk/basic/key/basicEventKey"
	"github.com/foxiswho/blog-go/pkg/tools/dbHelper/repositoryPg"
	"github.com/goccy/go-json"
	"github.com/jinzhu/copier"
	"github.com/pangu-2/go-tools/tools/strPg"
	"gorm.io/gorm"
)

type EventMakeCache struct {
	Log *log2.Logger `autowire:"?"`
	Sp  *Sp          `autowire:"?"`
	ct  modEventBasicEvent.EventDto
}

func NewEventMakeCache(sp *Sp, ct modEventBasicEvent.EventDto) *EventMakeCache {
	return &EventMakeCache{
		Sp:  sp,
		Log: sp.Log,
		ct:  ct,
	}
}

func (c *EventMakeCache) Processor(ctx context.Context) error {
	if c.ct.IsAll {
		return c.All(ctx)
	}
	if c.ct.IsThisTenantAll {
		return c.ThisTenantAll(ctx)
	}
	if strPg.IsNotBlank(c.ct.TenantNo) && !c.ct.IsThisTenantAll && len(c.ct.Nos) > 0 {
		return c.Nos(ctx)
	}
	return nil
}

func (c *EventMakeCache) ThisTenantAll(ctx context.Context) error {
	var query entityBasic.BasicConfigEventEntity
	if strPg.IsNotBlank(c.ct.TenantNo) {
		query.TenantNo = c.ct.TenantNo
		query.State = enumStatePg.ENABLE.Index()
		infos := c.Sp.repEvent.FindAll(query)
		if infos != nil {
			data := make(map[string]interface{})
			//keysAdd := make([]string, 0)
			for _, info := range infos {
				var obj modCacheBasicEvent.EventCache
				copier.Copy(&obj, info)
				//
				//key := basicEventKey.EventTenantNo(info.TenantNo, info.No)
				//data[key] = info.No
				//keysAdd = append(keysAdd, key)

				key2 := basicEventKey.EventTenantNo(info.TenantNo, info.No)
				str, err := json.Marshal(obj)
				if err == nil {
					data[key2] = str
				}
			}
			if len(data) > 0 {
				c.Sp.rdt.SetPipeline(ctx, data)
			}
			//if len(keysAdd) > 0 {
			//	keysAll := basicEventKey.EventTenantNoKeys(c.ct.TenantNo)
			//	err := c.Sp.rdt.GetRdb().SAdd(ctx, keysAll, keysAdd).Err()
			//	if err != nil {
			//		c.Sp.Log.Error("缓存失败:", err)
			//	}
			//}
		}
	}
	return nil
}

func (c *EventMakeCache) All(ctx context.Context) error {
	var query entityBasic.BasicConfigEventEntity
	if c.ct.IsAll {
		query.State = enumStatePg.ENABLE.Index()
		infos := c.Sp.repEvent.FindAll(query)
		if infos != nil {
			data := make(map[string]interface{})
			//mapKeysAdd := make(map[string][]string, 0)
			for _, info := range infos {
				var obj modCacheBasicEvent.EventCache
				copier.Copy(&obj, info)
				//
				key := basicEventKey.EventTenantNo(info.TenantNo, info.No)
				str, err := json.Marshal(obj)
				if err == nil {
					data[key] = str
				}

				//if _, ok := mapKeysAdd[info.TenantNo]; !ok {
				//	mapKeysAdd[info.TenantNo] = make([]string, 0)
				//}
				//mapKeysAdd[info.TenantNo] = append(mapKeysAdd[info.TenantNo], key)

				//key2 := basicEventKey.EventTenantNoByCode(info.TenantNo, info.No)
				//data[key2] = info.No
			}
			if len(data) > 0 {
				c.Sp.rdt.SetPipeline(ctx, data)
			}
			//if len(mapKeysAdd) > 0 {
			//	for tenantNo, keys := range mapKeysAdd {
			//		if nil == keys || len(keys) <= 0 {
			//			continue
			//		}
			//		keysAll := basicEventKey.EventTenantNoKeys(tenantNo)
			//		err := c.Sp.rdt.GetRdb().SAdd(ctx, keysAll, keys).Err()
			//		if err != nil {
			//			c.Sp.Log.Error("缓存失败:", err)
			//		}
			//	}
			//}
		}
	}
	return nil
}

func (c *EventMakeCache) Nos(ctx context.Context) error {
	var query entityBasic.BasicConfigEventEntity
	if strPg.IsNotBlank(c.ct.TenantNo) && !c.ct.IsThisTenantAll && len(c.ct.Nos) > 0 {
		nos := make([]string, 0)
		for _, item := range c.ct.Nos {
			if strPg.IsNotBlank(item) {
				nos = append(nos, item)
			}
		}
		if len(nos) <= 0 {
			return nil
		}
		query.TenantNo = c.ct.TenantNo
		query.State = enumStatePg.ENABLE.Index()
		infos := c.Sp.repEvent.FindAll(query, repositoryPg.ConditionOption(func(db *gorm.DB) *gorm.DB {
			db = db.Order("create_at desc")
			db.Where("no in ?", nos)
			return db
		}))
		if infos != nil {
			//keysAdd := make([]string, 0)
			data := make(map[string]interface{})
			for _, info := range infos {
				var obj modCacheBasicEvent.EventCache
				copier.Copy(&obj, info)
				//
				key := basicEventKey.EventTenantNo(info.TenantNo, info.No)
				str, err := json.Marshal(obj)
				if err == nil {
					data[key] = str
				}
				//keysAdd = append(keysAdd, key)

				//key2 := basicEventKey.EventTenantNoByCode(info.TenantNo, info.No)
				//data[key2] = info.No
			}
			if len(data) > 0 {
				c.Sp.rdt.SetPipeline(ctx, data)
			}
			//if len(keysAdd) > 0 {
			//	keysAll := basicEventKey.EventTenantNoKeys(c.ct.TenantNo)
			//	err := c.Sp.rdt.GetRdb().SAdd(ctx, keysAll, keysAdd).Err()
			//	if err != nil {
			//		c.Sp.Log.Error("缓存失败:", err)
			//	}
			//}
		}
	}
	return nil
}
