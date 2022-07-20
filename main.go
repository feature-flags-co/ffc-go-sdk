package main

import (
	"fmt"
	"github.com/feature-flags-co/ffc-go-sdk/common"
)

func main() {

	ffuser := new(common.FFCUser)
	ffuser.UserName = "test"
	fmt.Println(ffuser.UserName)

}
