package main

import (
	"github.com/labstack/echo"
	"github.com/tsubaoza0901/echo-rest-api/presenter/http/middleware"
	"github.com/tsubaoza0901/echo-rest-api/presenter/http/router"
)

func main() {
	// Echoのインスタンス生成
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェアをセット
	middleware.InitMiddleware(e)

	// ルーティングをセット
	router.InitRouting(e)

	// サーバー起動
	e.Start(":8080")
}
