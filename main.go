package main

import (
	"flag"
	"fmt"
	"net/http"
	"viry_sun/lib/config"
	"viry_sun/lib/controller"
	"viry_sun/lib/dao/clipboard"
	"viry_sun/lib/log"
	"viry_sun/lib/request"
	"viry_sun/lib/util"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

type myHandler struct {
	handler http.Handler
}

func (my myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//附加上下文
	isMobile := util.IsMobileUserAgent(r.UserAgent())
	dir := "./pc"
	if isMobile {
		dir = "./mobile"
	}
	r2 := request.SetTemplateDir(r, dir)

	my.handler.ServeHTTP(w, r2)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	var (
		task          string
		expireSeconds int
	)

	flag.StringVar(&task, "task", "", "命令行任务类型，如：removeExpire")
	flag.IntVar(&expireSeconds, "expireSeconds", 600, "过期时间（秒），如：600")
	flag.Parse()

	switch task {
	default:
		flag.Usage()
	case "removeExpire": //清理过期字段
		num, err := clipboard.RemoveExpire(expireSeconds)
		if err != "" {
			log.L.Error(fmt.Sprintf("Task removeExpire Error: %v", err))
		}
		log.L.Info(fmt.Sprintf("Task removeExpire Num: %d", num))
	case "":
		//路由定义
		router := httprouter.New()
		router.GET("/get/:name", controller.GetAct)
		router.POST("/save/:name", controller.SaveAct)
		router.GET("/hello/:name", Hello)

		//错误拦截
		router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
			log.L.Error(fmt.Sprintf("Request Error: %v", err), zap.Stack("Default Stack:"))
			http.Error(w, "Request Error!", http.StatusInternalServerError)
		}

		h := &myHandler{}
		h.handler = router

		log.L.Fatal(fmt.Sprintf("Start Web Server Fatal: %v", http.ListenAndServe(":"+config.C.Site.Port, h)))
	}

	//t1 := time.NewTicker(3 * time.Second)

	// for {

	// 	<-t1.C
	// 	log.L.Debug("Debug!!!")
	// 	log.L.Info("test!!!")
	// 	log.L.Error("error test!!!")
	// }
}
