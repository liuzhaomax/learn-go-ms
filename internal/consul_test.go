package internal

import (
	"fmt"
	"testing"
)

func TestReg(t *testing.T) {
	err := Reg(ViperConf.AccountWebConfig.Host, ViperConf.AccountWebConfig.SrvName, ViperConf.AccountWebConfig.SrvName,
		ViperConf.AccountWebConfig.Port, ViperConf.AccountWebConfig.Tags)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("注册成功")
	}
}
