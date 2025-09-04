package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default() // デフォルトのミドルウェア（ロギングやリカバリ）が設定される
    router.GET("/", func(c *gin.Context) {
        c.String(200, "Hello, World!")
    })
    router.Run(":8080") // ポート8080でサーバーを起動
}