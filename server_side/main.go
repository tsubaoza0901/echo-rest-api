package main

import (
	"github.com/labstack/echo"
	"github.com/tsubaoza0901/echo-rest-api/infrastructure/rdb"
	"github.com/tsubaoza0901/echo-rest-api/interactor"
	"github.com/tsubaoza0901/echo-rest-api/presenter/http/middleware"
	"github.com/tsubaoza0901/echo-rest-api/presenter/http/router"
)

func main() {
	// DBセットアップ
	conn := rdb.InitDB()
	// defer func() {
	// 	if db != nil {
	// 		if err := db.Close(); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 	}
	// }()

	i := interactor.NewInteractor(conn)
	h := i.NewAppHandler()

	// ctx := context.BackGround()

	// Echoのインスタンス生成
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェアをセット
	middleware.InitMiddleware(e)

	// ルーティングをセット
	router.InitRouting(e, h)

	// サーバー起動
	e.Start(":8080")
}
