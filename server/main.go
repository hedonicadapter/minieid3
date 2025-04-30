package main

import (
	"context"
	"fmt"

	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"net/url"

	"gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http/httptrace"

	"github.com/gin-gonic/gin"
	"github.com/hedonicadapter/gopher/api/routes"
	"github.com/hedonicadapter/gopher/config"
	"github.com/hedonicadapter/gopher/models"
	"github.com/hedonicadapter/gopher/services/queue"
	"github.com/hedonicadapter/gopher/services/user"
)

func main() {
	config.InitEnv()
	db := config.InitDb()
	rdb := config.InitRedis()
	config.IdempotentDummyData(db)

	r := gin.Default()

	userService := user.InitService(db)
	queueService := queue.InitService(rdb, "main")
	routes.UserRoutes(r.Group("api"), userService, queueService)

	r.GET("health", func(ctx *gin.Context) {
		ctx.JSONP(200, gin.H{
			"status": "OK",
		})
	})

	go queueService.Poll(context.Background(), func(task models.Task) any {
		fmt.Println(task.Action)

		return ""
	})

	// ----------------------------------------------------------------------------------

	backendURL := "add new service and put url here"
	parsedURL, _ := url.Parse(backendURL)

	proxy := httputil.NewSingleHostReverseProxy(parsedURL)
	proxy.Transport = httptrace.WrapRoundTripper(&http.Transport{
		MaxIdleConnsPerHost: 20,
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	})

	r.Any("/proxy/*path", func(c *gin.Context) {
		c.Request.URL.Path = c.Param("path")
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	r.Run()
}

// /*   httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http" */ proxy := httputil.NewSingleHostReverseProxy(&cfg.BackendURL) proxy.Transport = httptrace.WrapRoundTripper(&http.Transport{     MaxIdleConnsPerHost: 20,     TLSClientConfig: &tls.Config{         MinVersion:   tls.VersionTLS12,         Certificates: clientAndCertificates.certificates.Certificates,         RootCAs:      clientAndCertificates.certificates.RootCAs,     }, })
//
// proxy.Serve.HTTP(c.Write, c.Request)
