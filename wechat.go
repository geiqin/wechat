package wechat

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/geiqin/wechat/cache"
	"github.com/geiqin/wechat/miniprogram"
	miniConfig "github.com/geiqin/wechat/miniprogram/config"
	"github.com/geiqin/wechat/officialaccount"
	offConfig "github.com/geiqin/wechat/officialaccount/config"
	"github.com/geiqin/wechat/openplatform"
	openConfig "github.com/geiqin/wechat/openplatform/config"
	"github.com/geiqin/wechat/pay"
	payConfig "github.com/geiqin/wechat/pay/config"
	"github.com/geiqin/wechat/work"
	workConfig "github.com/geiqin/wechat/work/config"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

// Wechat struct
type Wechat struct {
	cache cache.Cache
}

// NewWechat init
func NewWechat() *Wechat {
	return &Wechat{}
}

// SetCache 设置cache
func (wc *Wechat) SetCache(cache cache.Cache) {
	wc.cache = cache
}

// GetOfficialAccount 获取微信公众号实例
func (wc *Wechat) GetOfficialAccount(cfg *offConfig.Config) *officialaccount.OfficialAccount {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return officialaccount.NewOfficialAccount(cfg)
}

// GetMiniProgram 获取小程序的实例
func (wc *Wechat) GetMiniProgram(cfg *miniConfig.Config) *miniprogram.MiniProgram {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return miniprogram.NewMiniProgram(cfg)
}

// GetPay 获取微信支付的实例
func (wc *Wechat) GetPay(cfg *payConfig.Config) *pay.Pay {
	return pay.NewPay(cfg)
}

// GetOpenPlatform 获取微信开放平台的实例
func (wc *Wechat) GetOpenPlatform(cfg *openConfig.Config) *openplatform.OpenPlatform {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return openplatform.NewOpenPlatform(cfg)
}

// GetWork 获取企业微信的实例
func (wc *Wechat) GetWork(cfg *workConfig.Config) *work.Work {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return work.NewWork(cfg)
}
