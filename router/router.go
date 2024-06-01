package router

import (
	"practice/di"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.Default()

		  // ここからCorsの設定
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:5173",
		},
		// アクセスを許可したいHTTPメソッド
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Accept",
			"Accept-Encoding",
			"Accept-Language",
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Authorization",
			"Connection",
			"Host",
			"Origin",
			"Referer",
			"Sec-Ch-Ua",
			"Sec-Ch-Ua-Mobile",
			"Sec-Ch-Ua-Platform",
			"Sec-Fetch-Dest",
			"Sec-Fetch-Mode",
			"Sec-Fetch-Site",
			"User-Agent",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: false,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	v1 := r.Group("/v1")

	{
		// systemHandler := di.InitSystemHandler()
		// v1.GET("/systems/ping", systemHandler.Ping)

		todoHandler := di.InitTodoHandler()
		v1.GET("/users", todoHandler.GetAllUser)
		v1.GET("/todos", todoHandler.GetAll)
		v1.GET("/todos/:id", todoHandler.GetById)
		v1.POST("/users", todoHandler.RegistUser)
		v1.POST("/todos", todoHandler.Create)
		v1.PUT("/todos/:id", todoHandler.Update)
		v1.DELETE("/todos/:id", todoHandler.Delete)
	}

	return r
}