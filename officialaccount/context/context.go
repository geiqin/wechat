package context

import (
	"github.com/geiqin/wechat/credential"
	"github.com/geiqin/wechat/officialaccount/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
