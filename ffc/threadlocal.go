package ffc

// Package ffc ThreadLocal is not in use, it will be use later
import (
	"github.com/feature-flags-co/ffc-go-sdk/model"
	"github.com/timandy/routine"
)

var threadLocal routine.ThreadLocal
var inheritableThreadLocal routine.ThreadLocal

func init() {
	threadLocal = routine.NewThreadLocal()
	inheritableThreadLocal = routine.NewInheritableThreadLocal()
}

func GetCurrentUser() model.FFCUser {

	var ffUser model.FFCUser
	ffUser = inheritableThreadLocal.Get().(model.FFCUser)
	if ffUser.IsEmpty() {
		ffUser = threadLocal.Get().(model.FFCUser)
	}
	return ffUser
}

func Remove() {
	threadLocal.Remove()
	inheritableThreadLocal.Remove()
}

func SetCurrentUser(user model.FFCUser, inherit bool) {

	if inherit {
		inheritableThreadLocal.Set(user)
	} else {
		threadLocal.Set(user)
	}
}
