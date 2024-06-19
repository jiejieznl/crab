package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var conf *Config

func New() *gin.Engine {
	app := gin.New()
	app.Use(gin.Recovery())
	app.Use(gin.Logger())
	app.Use(cors)
	return app
}

func Run(app *gin.Engine) {
	gin.SetMode(conf.Mode)
	go func() {
		if err := app.Run(fmt.Sprintf(":%d", conf.Port)); err != nil {
			log.Fatal("服务端启动失败", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//处理退出的逻辑
	fmt.Println("关闭server")

}

func Register(cfg *Config) {
	conf = cfg
}

func cors(c *gin.Context) {
	method := c.Request.Method               // 获取当前请求的方法
	origin := c.Request.Header.Get("Origin") // 获取请求头中的Origin字段，用于判断是否跨域

	// 如果是跨域请求，则设置相应的Access-Control-*头部
	if origin != "" {
		c.Header("Access-Control-Allow-Origin", "*")                                                                                                                          // 允许所有来源访问
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")                                                                                   // 允许的请求方法
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Accept-URL")                                                 // 允许的请求头部
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type") // 允许响应中暴露的头部
		c.Header("Access-Control-Allow-Credentials", "true")                                                                                                                  // 允许使用凭证（如cookies）
	}

	// 如果请求方法为OPTIONS，则直接返回200，不继续处理后续中间件
	if strings.ToUpper(method) == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
	}

	c.Next() // 继续处理后续的中间件和处理函数
}
