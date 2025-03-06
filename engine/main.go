package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/gin-gonic/gin"

	"modb"
	"web"

	log "github.com/tengfei-xy/go-log"
	tools "github.com/tengfei-xy/go-tools"
)

type App struct {
	loglevel int
	fontpath string
}

var app App

func init_mongo() {
	log.Infof("mongo连接中...")
	str := os.Getenv("FONTLIB_DATABASE")
	if str == "" {
		str = "mongodb://localhost:27017/"
	}
	err := modb.Init(str)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("mongo连接成功!!")
}
func init_log() {
	log.SetLevelInt(app.loglevel)
	_, g := log.GetLevel()
	fmt.Printf("日志等级:%s\n", g)
}
func init_flag() {
	flag.IntVar(&app.loglevel, "v", log.LEVELINFOINT, fmt.Sprintf("日志等级,%d-%d", log.LEVELFATALINT, log.LEVELDEBUG3INT))
	// 字体文件路径
	flag.StringVar(&app.fontpath, "p", "fonts/", "字体文件路径")
	flag.Parse()
	init_path()
}
func init_path() {
	paths := []string{filepath.Join(".", app.fontpath), filepath.Join("..", app.fontpath)}
	for _, p := range paths {
		if tools.FileExist(p) {
			app.fontpath = p
			return
		}
	}
	log.Fatalf("字体文件路径不存在:%s", app.fontpath)
}
func env(c *gin.Context) {
	c.Set("fontpath", app.fontpath)
	c.Next()
}
func main() {
	init_flag()
	init_log()

	port := "21520"
	log.Infof("监听端口:%s", port)

	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.Use(env)
	web.FontRoute(g)
	web.FontsRoute(g)

	// init_mongo()
	go quit()
	log.Info("已启动")
	err := g.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}

}
func quit() {
	// 创建一个通道来接收信号通知
	sigs := make(chan os.Signal, 1)

	// 监听 SIGINT 和 SIGTERM 信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGINT)
	log.Infof("PID: %d", os.Getpid())
	// 阻塞等待信号
	sig := <-sigs
	fmt.Println(sig)

	err := modb.Disconnect()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(1)
}
