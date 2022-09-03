package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-go-ms/account_web/handler"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "输入ip")
	port := flag.Int("port", 8081, "输入port")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *ip, *port)
	r := gin.Default()
	accountGroup := r.Group("/v1/account")
	{
		accountGroup.GET("/list", handler.AccountListHandler)
		accountGroup.POST("/login", handler.LoginByPasswordHandler)
	}
	r.Run(addr)
}
