package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/is-a-gamer/file-service/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start() {
	wd := NewWebDav()
	wd.Init()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	root := r.Group("/api")
	{
		// jwt2020-06-04: eyJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJzeXN0ZW0iLCJzdWIiOiJ7XCJ1c2VyaWRcIjpcIjEwMDBcIixcInVzZXJuYW1lXCI6XCJ1YnVudHVcIixcInByZWZlcnJlZF91c2VybmFtZVwiOlwidWJ1bnR1LCwsXCIsXCJlbWFpbFwiOlwiXCIsXCJlbWFpbF92ZXJpZmllZFwiOmZhbHNlLFwiYWRtaW5cIjpmYWxzZSxcImdyb3Vwc1wiOm51bGx9IiwiZXhwIjoxNTkxODYzODMyLCJuYmYiOjE1OTEyNTkwMzIsImlhdCI6MTU5MTI1OTAzMiwianRpIjoiNWY1ZjM3OTEtZjFmNC00NmJiLTk5MjUtNTUzODFjNWIwNjgyIn0.QV0amEYlM24uBk8GbaLH5bOtRVuAh10UjXM8JbOGkLc
		webDAVHandler(root, wd.HandlerFunc)
	}
	server := &http.Server{Addr: ":"+viper.GetString("port"), Handler: r}
	go server.ListenAndServe()
	gracefulExit(server)
}

func gracefulExit(server *http.Server) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Logger.Error("shutdown http server with error: ", zap.Error(err))
	} else {
		log.Logger.Info("shutdown http server successfully.")
	}
}