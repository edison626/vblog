package protocol

import (
	"context"
	"fmt"
	"net/http"

	"github.com/edison626/vblog/conf"
	"github.com/edison626/vblog/ioc"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() *HttpServer {
	r := gin.Default()
	ioc.ApiHandler().RouteRegistry(r.Group("/api/vblog"))

	return &HttpServer{
		server: &http.Server{
			// 服务监听的地址
			Addr: conf.C().App.HttpAddr(),
			// 监听关联的路由处理
			Handler: r,
		},
	}
}

// 封装一个自己的http server
type HttpServer struct {
	server *http.Server
}

func (s *HttpServer) Run() error {
	fmt.Printf("listen addr: %s\n", conf.C().App.HttpAddr())
	return s.server.ListenAndServe()
}

// 优雅关闭
func (s *HttpServer) Close(ctx context.Context) {
	s.server.Shutdown(ctx)
}
