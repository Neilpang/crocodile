package model

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter"
	"github.com/labulaka521/crocodile/common/log"
	"github.com/labulaka521/crocodile/core/config"
	"go.uber.org/zap"
)

var (
	enforcer *casbin.Enforcer
)

func InitRabc() {
	modeltext := `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`
	dbcfg := config.CoreConf.Db
	m, err := model.NewModelFromString(modeltext)
	if err != nil {
		log.Panic("NewModelFromString Err", zap.Error(err))
	}
	a, err := xormadapter.NewAdapter(dbcfg.Drivename, dbcfg.Dsn)
	if err != nil {
		log.Panic("NewAdapter Err", zap.Error(err))
	}

	enforcer, err = casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatal("InitRabc failed", zap.Error(err))
	}

}

func GetEnforcer() *casbin.Enforcer {
	return enforcer
}
