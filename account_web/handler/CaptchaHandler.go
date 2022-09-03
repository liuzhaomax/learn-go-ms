package handler

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"learn-go-ms/internal"
	"net/http"
	"os"
	"time"
)

func CaptchaHandler(c *gin.Context) {
	mobile, ok := c.GetQuery("mobile")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	fileName := "captcha.png"
	f, err := os.Create(fileName)
	if err != nil {
		zap.S().Error("GenCaptcha 失败")
		return
	}
	defer f.Close()
	var w io.WriterTo
	d := captcha.RandomDigits(captcha.DefaultLen)
	w = captcha.NewImage("", d, captcha.StdWidth, captcha.StdHeight)
	_, err = w.WriteTo(f)
	if err != nil {
		zap.S().Error("GenCaptcha 失败")
		return
	}
	fmt.Println(d)
	_captcha := ""
	for _, item := range d {
		_captcha += fmt.Sprintf("%d", item)
	}
	fmt.Println(_captcha)
	internal.RedisClient.Set(context.Background(), mobile, _captcha, 120*time.Second)
	b64, err := GetBase64(fileName)
	if err != nil {
		zap.S().Error("GenCaptcha 失败")
		return
	}
	fmt.Println(b64)
	c.JSON(http.StatusOK, gin.H{
		"captcha": b64,
	})
}

func GetBase64(fileName string) (string, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	// TODO 判断文件大小，给一个相对合理的数值
	b := make([]byte, 102400)
	base64.StdEncoding.Encode(b, file)
	s := string(b)
	return s, nil
}
