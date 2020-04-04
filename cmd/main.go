package main

import (
	"github.com/Development/golang-echo-docker-compose/cmd/test"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/", test.Hello())

	e.POST("/routing_check", test.RoutingCheck())
	// サーバー起動
	e.Start(":8080") //ポート番号指定してね
}
