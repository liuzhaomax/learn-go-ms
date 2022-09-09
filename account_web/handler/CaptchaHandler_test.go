package handler

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestCaptchaHandler(t *testing.T) {
	CaptchaHandler(&gin.Context{})
}
