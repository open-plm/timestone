package main

import (
	"github.com/open-plm/timestone/pkg/config"
	"github.com/open-plm/timestone/pkg/server"
)

func main() {
	conf := config.NewBaseConfig()
	server.Main(conf)
}