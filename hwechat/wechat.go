package hwechat

import (
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

type GomOAConf struct {
	RedisCacheOpts      *cache.RedisOpts  // cache redis的配置
	OfficialAccountConf *offConfig.Config // officialAccount 配置
}

func NewGomOA(conf *GomOAConf) *GomOA {
	// 创建wechat
	wc := wechat.NewWechat()

	// 设置redis存储accesstoken
	if conf.RedisCacheOpts != nil {
		redisCache := cache.NewRedis(conf.RedisCacheOpts)
		wc.SetCache(redisCache)
	}

	// 创建officialAccount
	officialAccount := wc.GetOfficialAccount(conf.OfficialAccountConf)

	return &GomOA{
		conf:            conf,
		wc:              wc,
		officialAccount: officialAccount,
	}
}

type GomOA struct {
	wc              *wechat.Wechat
	conf            *GomOAConf
	officialAccount *officialaccount.OfficialAccount
}

func (t *GomOA) Get() *officialaccount.OfficialAccount {
	return t.officialAccount
}
