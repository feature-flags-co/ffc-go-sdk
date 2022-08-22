package utils

import (
	"log"
	"testing"
)

func TestPercentageOfKey(t *testing.T) {

	cases := []struct {
		args string
		ret  float64
	}{
		{
			args: "test-value",
			ret:  0.14653629204258323,
		},
		{
			args: "qKPKh1S3FolC",
			ret:  0.9105919692665339,
		},
		{
			args: "3eacb184-2d79-49df-9ea7-edd4f10e4c6f",
			ret:  0.08994403155520558,
		},
	}

	for _, v := range cases {

		t.Run(v.args, func(t *testing.T) {
			value := PercentageOfKey(v.args)
			if value == v.ret {
				log.Printf("param = %v, exp result = %v, real result = %v", v.args, v.ret, value)
			}else{
				log.Printf("error param = %v, exp result = %v, real result = %v", v.args, v.ret, value)
			}
		})
	}

}
