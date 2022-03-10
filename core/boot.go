package core

import (
	"github.com/mashnoor/blind_cat/settings"
	"sync"
)

func initHealthCheckSystem() {
	var wg sync.WaitGroup

	for _, service := range settings.SystemAppConfig.Services {
		currentService := service
		//fmt.Println(service.Endpoint)
		go checkHealth(&currentService, &wg)
		wg.Add(1)

	}

	wg.Wait()
}

func BootApp() {
	settings.LoadAppConfig()
	settings.InitRedis()
	initHealthCheckSystem()

}
