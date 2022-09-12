package handler

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"learn-go-ms/account_srv/proto/pb"
	"learn-go-ms/account_web/req"
	"learn-go-ms/account_web/res"
	"learn-go-ms/custom_error"
	"learn-go-ms/jwt_op"
	"learn-go-ms/log"
	"net/http"
	"strconv"
	"time"
)

func HandleError(err error) string {
	if err != nil {
		switch err.Error() {
		case custom_error.AccountExisted:
			return custom_error.AccountExisted
		case custom_error.AccountNotFound:
			return custom_error.AccountNotFound
		case custom_error.SaltError:
			return custom_error.SaltError
		default:
			return custom_error.InternalError
		}
	}
	return ""
}

func AccountListHandler(c *gin.Context) {
	pageNoStr := c.DefaultQuery("pageNo", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "3")
	pageNo, _ := strconv.ParseInt(pageNoStr, 10, 32)
	pageSize, _ := strconv.ParseInt(pageSizeStr, 10, 32)
	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure())
	if err != nil {
		s := fmt.Sprintf("AccountListHandler-GRPC拨号失败：%s", err.Error())
		log.Logger.Info(s)
		e := HandleError(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": e,
		})
		return
	}
	client := pb.NewAccountServiceClient(conn)
	r, err := client.GetAccountList(context.Background(), &pb.PagingRequest{
		PageNo:   uint32(pageNo),
		PageSize: uint32(pageSize),
	})
	if err != nil {
		s := fmt.Sprintf("AccountListHandler-GRPC调用失败：%s", err.Error())
		log.Logger.Info(s)
		e := HandleError(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": e,
		})
		return
	}
	var resList []res.Account4Res
	for _, item := range r.AccountList {
		resList = append(resList, pb2res(item))
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":   "ok",
		"total": r.Total,
		"data":  resList,
	})
}

func pb2res(accountRes *pb.AccountRes) res.Account4Res {
	return res.Account4Res{
		Mobile:   accountRes.Mobile,
		NickName: accountRes.Nickname,
		Gender:   accountRes.Gender,
	}
}

func LoginByPasswordHandler(c *gin.Context) {
	var loginByPassword req.LoginByPassword
	err := c.ShouldBind(&loginByPassword)
	if err != nil {
		log.Logger.Error("LoginByPassword出错" + err.Error())
		c.JSON(http.StatusOK, gin.H{
			"msg": "解析参数错误",
		})
		return
	}
	// TODO 校验手机号码正则表达式
	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure())
	if err != nil {
		s := fmt.Sprintf("LoginByPasswordHandler 拨号出错：%s", err.Error())
		log.Logger.Error(s)
		e := HandleError(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": e,
		})
		return
	}
	client := pb.NewAccountServiceClient(conn)
	r, err := client.GetAccountByMobile(context.Background(), &pb.MobileRequest{
		Mobile: loginByPassword.Mobile,
	})
	if err != nil {
		s := fmt.Sprintf("GRPC GetAccountByMobile 出错：%s", err.Error())
		log.Logger.Error(s)
		e := HandleError(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": e,
		})
		return
	}
	cheRes, err := client.CheckPassword(context.Background(), &pb.CheckPasswordRequest{
		Password:       loginByPassword.Password,
		HashedPassword: r.Password,
		AccountId:      uint32(r.Id),
	})
	if err != nil {
		s := fmt.Sprintf("GRPC GetAccountByMobile 出错：%s", err.Error())
		log.Logger.Error(s)
		e := HandleError(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": e,
		})
		return
	}
	checkResult := "登录失败"
	if cheRes.Result {
		checkResult = "登录成功"
		j := jwt_op.NewJWT()
		now := time.Now()
		claims := jwt_op.CustomClaims{
			StandardClaims: jwt.StandardClaims{
				NotBefore: now.Unix(),
				ExpiresAt: now.Add(time.Hour * 24 * 30).Unix(),
			},
			ID:          r.Id,
			NickName:    r.Nickname,
			AuthorityId: int32(r.Role),
		}
		token, err := j.GenerateJWT(claims)
		if err != nil {
			s := fmt.Sprintf("GRPC GenerateJWT 出错：%s", err.Error())
			log.Logger.Error(s)
			e := HandleError(err)
			c.JSON(http.StatusOK, gin.H{
				"msg": e,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":    "",
			"result": checkResult,
			"token":  token,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "",
		"result": checkResult,
		"token":  "",
	})
}

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
