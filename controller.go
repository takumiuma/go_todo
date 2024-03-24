// package main

// import (
// 	"time"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// )

// func main() {

//   r := gin.Default()

//   // ここからCorsの設定
//   r.Use(cors.New(cors.Config{
//     // アクセスを許可したいアクセス元
//     AllowOrigins: []string{
//         "https://localhost:3000",
//     },
//     // アクセスを許可したいHTTPメソッド
//     AllowMethods: []string{
//       "POST",
//       "GET",
//       "OPTIONS",
// 		  "PUT",
// 		  "DELETE",
//     },
//     // 許可したいHTTPリクエストヘッダ
//     AllowHeaders: []string{
//         "Access-Control-Allow-Credentials",
//         "Access-Control-Allow-Headers",
//         "Content-Type",
//         "Content-Length",
//         "Accept-Encoding",
//         "Authorization",
//     },
//     // cookieなどの情報を必要とするかどうか
//     AllowCredentials: false,
//     // preflightリクエストの結果をキャッシュする時間
//     MaxAge: 24 * time.Hour,
//   }))

//   r.GET("/", func(c *gin.Context) {
//     c.String(200, "Hello,World!")
//   })
//   // r.POST("/api/test", controller.TestMethod)

//   r.Run(":9000")
// }