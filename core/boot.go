package core

import (
	"github.com/mashnoor/blind_cat/settings"
	"github.com/mashnoor/blind_cat/structures"
	"gopkg.in/yaml.v2"
	"os"
	"sync"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var (
	SystemAppConfig structures.AppConfig
)

func readConfigFile() string {
	dat, err := os.ReadFile("services.yaml")
	check(err)

	return string(dat)
}

func loadAppConfig() {
	config := readConfigFile()
	//fmt.Println(config)
	err := yaml.Unmarshal([]byte(config), &SystemAppConfig)
	check(err)
}

func initHealthCheckSystem() {
	var wg sync.WaitGroup

	for _, service := range SystemAppConfig.Services {
		currentService := service
		//fmt.Println(service.Endpoint)
		go checkHealth(&currentService, &wg)
		wg.Add(1)

	}

	wg.Wait()
}

func BootApp() {
	loadAppConfig()
	settings.InitRedis(SystemAppConfig.RedisHost, SystemAppConfig.RedisPort)
	initHealthCheckSystem()

}
