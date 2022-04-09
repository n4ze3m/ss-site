package controllers

import (
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func GenerateScreenshot(ctx *gin.Context) {
	paramUrl := ctx.Query("url")
	_, err := url.ParseRequestURI(paramUrl)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "invalid url",
		})
		return
	}
	path, _ := launcher.LookPath()
	l := launcher.New().Bin(path).Headless(true).Leakless(true).RemoteDebuggingPort(10000)
	defer l.Cleanup()
	u, _ := l.Launch()
	b := rod.New().ControlURL(u).MustConnect().Trace(true).Timeout(60 * time.Second)
	defer b.MustClose()
	p := b.MustPage(paramUrl)
	p.MustWaitLoad()

	c := p.MustScreenshot()

	ctx.Data(200, "image/png", c)
}
