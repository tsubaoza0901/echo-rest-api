package main

import (
	"os"

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

	// Heroku用設定
	if os.Getenv("PORT") != "" {
		e.Start(":" + os.Getenv("PORT"))
	} else {
		e.Start(":8080")
	}

	// // 通常用：サーバー起動
	// e.Start(":8080")

	// // Heroku用：サーバー起動
	// e.Start(":" + os.Getenv("PORT"))
}

// // 以下、初期ECS確認用---------------------------------

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// )

// // RequestLogger : Logs HTTP requests from specified routes
// func RequestLogger(targetMux http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()

// 		targetMux.ServeHTTP(w, r)

// 		// log request by who(IP address)
// 		requesterIP := r.RemoteAddr

// 		log.Printf(
// 			"%s\t\t%s\t\t%s\t\t%v",
// 			r.Method,
// 			r.RequestURI,
// 			requesterIP,
// 			time.Since(start),
// 		)
// 	})
// }
// func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
// 	w.WriteHeader(status)
// 	if status == http.StatusNotFound {
// 		fmt.Fprint(w, "404 Page Not Found")
// 	}
// }

// func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	html := "Hello World!4"
// 	if r.URL.Path != "/" {
// 		errorHandler(w, r, http.StatusNotFound)
// 		return
// 	}
// 	w.Write([]byte(html))
// }
// func byeWorld(w http.ResponseWriter, r *http.Request) {
// 	html := "Bye World!"

// 	if r.URL.Path != "/bye" {
// 		errorHandler(w, r, http.StatusNotFound)
// 		return
// 	}
// 	w.Write([]byte(html))
// }

// func main() {
// 	metrics := promhttp.Handler()
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", helloWorld)
// 	mux.HandleFunc("/bye", byeWorld)
// 	mux.Handle("/metrics", metrics)
// 	// fmt.Printf("Starting server at Port %d", strconv.Atoi(os.Getenv("PORT")))
// 	http.ListenAndServe(":"+os.Getenv("PORT"), RequestLogger(mux))

// 	// Herokuを使わない場合は以下でOK-------------
// 	// metrics := promhttp.Handler()
// 	// mux := http.NewServeMux()
// 	// mux.HandleFunc("/", helloWorld)
// 	// mux.HandleFunc("/bye", byeWorld)
// 	// mux.Handle("/metrics", metrics)
// 	// http.ListenAndServe(":8080", RequestLogger(mux))
// }

// // echoテスト用------------------------------

// package main

// import (
// 	"net/http"
// 	"os"

// 	"github.com/labstack/echo"
// )

// // Book ...
// type Book struct {
// 	Title  string `json:"title"`
// 	Author string `json:"author"`
// }

// func main() {
// 	// Echo インスタンスをNew
// 	e := echo.New()

// 	// ルーティング設定
// 	e.GET("/books/:id", bookHandler)

// 	// サーバ起動
// 	e.Start(":" + os.Getenv("PORT"))
// }

// // ハンドラ
// func bookHandler(c echo.Context) error {
// 	return c.JSON(http.StatusOK, getBook())
// }

// func getBook() *Book {
// 	b := &Book{
// 		Title:  "Golangの本",
// 		Author: "Golangの本の著者",
// 	}
// 	return b
// }
