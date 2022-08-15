package ffc

import (
	"github.com/feature-flags-co/ffc-go-sdk/common"
	"github.com/timandy/routine"
)

var threadLocal routine.ThreadLocal
var inheritableThreadLocal routine.ThreadLocal

func init() {
	threadLocal = routine.NewThreadLocal()
	inheritableThreadLocal = routine.NewInheritableThreadLocal()
}

func GetCurrentUser() common.FFCUser {

	var ffUser common.FFCUser
	ffUser = inheritableThreadLocal.Get().(common.FFCUser)
	if ffUser.IsEmpty() {
		ffUser = threadLocal.Get().(common.FFCUser)
	}
	return ffUser
}

func Remove() {
	threadLocal.Remove()
	inheritableThreadLocal.Remove()
}

func SetCurrentUser(user common.FFCUser, inherit bool) {

	if inherit {
		inheritableThreadLocal.Set(user)
	} else {
		threadLocal.Set(user)
	}
}
