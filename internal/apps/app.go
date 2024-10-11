package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"ses_back/internal/config"
	"ses_back/internal/service"
)

func App() {
	config.ReadConfig()
	var svrCFG = config.Config.Server
	server := gin.Default()

	// Services:
	go server.GET(
		"/events",
		service.GetEvents,
	)

	go server.POST(
		"/events",
		service.GetEvByFilter,
	)

	//---------------------------------------------------------------//
	var addr = fmt.Sprintf("%s:%s", svrCFG.Host, svrCFG.Port)
	fmt.Printf("Server run on https://%s\n", addr)

	err := server.Run(addr)
	if err != nil {
		log.Fatal("Server run error: ", err)
	}
}
