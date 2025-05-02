package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	// "time"

	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"

	"github.com/gin-gonic/gin"
	"github.com/hedonicadapter/gopher/utils"
)

func main() {
	cert, err := tls.LoadX509KeyPair("client-cert.pem", "client-key.pem")
	if err != nil {
		fmt.Println("Error loading client certificate:", err)
		return
	}

	rootCA := x509.NewCertPool()
	caCert, err := utils.ReadFile("ca-cert.pem")
	if err != nil {
		fmt.Println("Error loading root CA:", err)
		return
	}
	if !rootCA.AppendCertsFromPEM([]byte(caCert)) {
		fmt.Println("Failed to append root CA")
		return
	}

	backendURL := "http://mock-external-server:8080/"
	parsedURL, err := url.Parse(backendURL)
	if err != nil {
		fmt.Println("Invalid backend URL:", err)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(parsedURL)
	proxy.Transport = httptrace.WrapRoundTripper(&http.Transport{
		MaxIdleConnsPerHost: 20,
		TLSClientConfig: &tls.Config{
			MinVersion:   tls.VersionTLS12,
			Certificates: []tls.Certificate{cert},
			RootCAs:      rootCA,
		},
		// TLSHandshakeTimeout:   10 * time.Second,
		// ResponseHeaderTimeout: 33 * time.Second,
		// IdleConnTimeout:       30 * time.Second,
		// ExpectContinueTimeout: 1 * time.Second,
	})

	// config.InitEnv()
	// db := config.InitDb()
	// rdb := config.InitRedis()
	// config.IdempotentDummyData(db)
	//
	r := gin.Default()
	//
	// userService := user.InitService(db)
	// queueService := queue.InitService(rdb, "main")
	// routes.UserRoutes(r.Group("api"), userService, queueService)
	//
	// r.GET("health", func(ctx *gin.Context) {
	// 	ctx.JSONP(200, gin.H{
	// 		"status": "OK",
	// 	})
	// })
	//
	// go queueService.Poll(context.Background(), func(task models.Task) any {
	// 	fmt.Println(task.Action)
	//
	// 	return ""
	// })
	r.Any("/proxy/*path", func(c *gin.Context) {
		c.Request.URL.Path = c.Param("path")
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	r.Run()
}

// proxy := httputil.NewSingleHostReverseProxy(&cfg.BackendURL)
// proxy.Transport = httptrace.WrapRoundTripper(&http.Transport{
// 	MaxIdleConnsPerHost: 20,
// 	TLSClientConfig: &tls.Config{
// 		MinVersion: tls.VersionTLS12,
// 		Certificates: clientAndCertificates.certificates.Certificates,
// 		RootCAs: clientAndCertificates.certificates.RootCAs
// 	}
// })
