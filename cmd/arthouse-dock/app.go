package main

import (
	"log"
	"net/http"

	"github.com/artHouse-Docs/backend/internal"
	"github.com/artHouse-Docs/backend/pkg/config"
	"github.com/rs/cors"
)

func InitServer(cfg config.ServerConfig) error {
	r := internal.InitRoute()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	return http.ListenAndServe(
		cfg.Host+":"+cfg.Port,
		c.Handler(r),
	)
}

func main() {
	srvCfg := config.Configure().Server

	if err := InitServer(srvCfg); err != nil {
		log.Panic(err)
	}
}
