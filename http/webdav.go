package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/net/webdav"
	"net/http"
)
type WebDAV struct {
	fs          webdav.FileSystem
	HandlerFunc gin.HandlerFunc
}

func NewWebDav() (wd *WebDAV) {
	return &WebDAV{}
}

func (wd *WebDAV) Init() {
	wd.fs = webdav.Dir(viper.GetString("webdav.folder_path"))

	handler := &webdav.Handler{
		Prefix:     "/api/webdav",
		FileSystem: wd.fs,
		LockSystem: webdav.NewMemLS(),
	}

	wd.HandlerFunc = func(c *gin.Context) {
		r := c.Request
		w := c.Writer
		u, p, ok := r.BasicAuth()
		err := handleThirPartyAuth(u, p)
		if !ok || err != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="davfs"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		handler.ServeHTTP(w, r)
		//c.String(http.StatusOK, "")
	}
}

func handleThirPartyAuth(u, p string) (err error) {
	request := new(http.Request)
	request.Header = make(http.Header)
	if u == "" || p == "" {
		return fmt.Errorf("no user validate")
	}
	if u != viper.GetString("webdav.username") || p != viper.GetString("webdav.password") {
		return fmt.Errorf("username or password validate")
	}
	request.SetBasicAuth(u, p)
	request.Method = "OPTIONS"
	return nil
}


func webDAVHandler(rg *gin.RouterGroup, fn gin.HandlerFunc) {
	r := rg.Group("/webdav")
	{
		r.Any("/*any", fn)
		r.Handle("PROPFIND", "/*any", fn)
		r.Handle("PROPPATCH", "/*any", fn)
		r.Handle("MKCOL", "/*any", fn)
		r.Handle("COPY", "/*any", fn)
		r.Handle("MOVE", "/*any", fn)
		r.Handle("LOCK", "/*any", fn)
		r.Handle("UNLOCK", "/*any", fn)
	}
}