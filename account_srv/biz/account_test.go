package biz

import (
	"context"
	"fmt"
	"learn-go-ms/account_srv/internal"
	"learn-go-ms/account_srv/proto/pb"
	"testing"
)

func init() {
	internal.InitDB()
}

func TestAccountServer_AddAccount(t *testing.T) {
	accountServer := AccountServer{}
	for i := 0; i < 5; i++ {
		s := fmt.Sprintf("1300000000%d", i)
		res, err := accountServer.AddAccount(context.Background(), &pb.AddAccountRequest{
			Mobile:   s,
			Password: s,
			NickName: s,
			Gender:   "male",
		})
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(res.Id)
	}
}
