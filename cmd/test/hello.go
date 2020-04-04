package test

import (
	"encoding/json"
	"github.com/Development/golang-echo-docker-compose/internal/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"

)

type test struct {
	Test string `json:"test"`
}

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		var t  user.User
		t = user.User{}
		t.GetUser()
		log.Printf("%+v", t)
		log.Printf("%+v", &t)

		bytes, e := json.Marshal(t)
		if e != nil {
			log.Error(e)
		}
		println(string(bytes))

		log.Printf("%+v", t)
		//var tt *test
		//e = json.Unmarshal(marshal,tt)


		return c.JSONBlob(http.StatusOK, bytes)
	}
}
