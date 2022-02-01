package main

import (
	"flag"
	"fmt"
	"incorta-cloud/upload-service/config"
	"incorta-cloud/upload-service/datasources"
	"incorta-cloud/upload-service/server"
	"os"
)

// @title Service Title
// @version         1.1
// @description Service Description

// @contact.name   Incorta Cloud
// @contact.url    http://cloud.incorta.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /api/v1

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        X-Api-Token

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	datasources.InitDatastore()
	datasources.InitOauthHttpClient()
	datasources.InitStorage()
	server.Init()
}
