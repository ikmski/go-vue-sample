package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	server serverConfig
}

type serverConfig struct {
	TLS      bool   `default:"false"`
	Port     int    `default:"8080"`
	CertFile string `envconfig:"cert_file" default:""`
	KeyFile  string `envconfig:"key_file" default:""`
}

func (c *serverConfig) getAddr() string {
	return fmt.Sprintf(":%d", c.Port)
}

func newConfig() *config {

	var server serverConfig
	err := envconfig.Process("server", &server)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	c := &config{
		server: server,
	}

	return c
}
