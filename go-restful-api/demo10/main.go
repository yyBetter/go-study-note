package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go-study-note/go-restful-api/demo8/config"
	"go-study-note/go-restful-api/demo8/model"
	"go-study-note/go-restful-api/demo8/router"
	"go-study-note/go-restful-api/demo8/router/middleware"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "demo2 config file path.")
)

func main() {
	pflag.Parse()
	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// create the gin engine.
	g := gin.New()
	// gin middlewares
	// routes
	router.Load(
		// Cores.
		g,
		// 全局中间件
		middleware.Logging(),
		middleware.RequestId(),
	)
	// 设置gin运行模式 debug 会打出更多日志信息
	gin.SetMode(viper.GetString("runmode"))
	// init db
	//model.DB.Init()
	defer model.DB.Close()

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Infof("The router has been deployed successfully.")
	}()

	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("tls.addr"))
		log.Infof(http.ListenAndServeTLS(":"+viper.GetString("tls.addr"), cert, key, g).Error())
	}

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Infof(http.ListenAndServe(":"+viper.GetString("addr"), g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Infof("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}

// 此外可以从环境变量读取配置，export SERVER_ADDR=:8880
// 配置文件可以热更新
