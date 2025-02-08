package server

import "github.com/open-plm/timestone/pkg/config"

func Main(conf config.BaseConfig) {
	router := InitRouter(conf)
	router.Run(":8000")
}