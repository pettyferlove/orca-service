package security

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"orca-service/global"
	security "orca-service/global/security/entity"
)

var text = `
[request_definition]
r = subject, object, action, service

[policy_definition]
p = subject, object, action, service, effect

[role_definition]
g = _, _

[policy_effect]
e = priority(p.effect) || deny

[matchers]
m = g(r.subject, p.subject) && keyMatch(r.object, p.object) && (r.action == p.action || p.action == "*") && (r.service == p.service || p.service == "*")
`

func InitSecurityEngine() error {
	securityRule := &security.SecurityRule{}
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(global.DatabaseClient, securityRule, securityRule.TableName())
	if err != nil {
		return err
	}
	m, err := model.NewModelFromString(text)
	if err != nil {
		panic(err)
	}
	enforcer, err := casbin.NewSyncedEnforcer(m, adapter)
	securityEffector := NewSecurityEffector()
	enforcer.SetEffector(securityEffector)
	if err != nil {
		return err
	}
	if err != nil {
		panic(err)
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		panic(err)
	}
	enforcer.EnableLog(false)
	global.Enforcer = enforcer
	return nil
}
