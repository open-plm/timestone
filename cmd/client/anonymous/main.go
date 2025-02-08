package main

import (
	"github.com/open-plm/timestone/pkg/client"
	"github.com/open-plm/timestone/pkg/config"
)

func main() {
	conf := config.NewBaseConfig()
	client.ClientForAnonymousUser(conf)
}