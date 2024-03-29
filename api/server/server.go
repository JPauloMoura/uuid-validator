package server

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/JPaulo-Moura/uuid-validator/api/handlers"
	_ "github.com/JPaulo-Moura/uuid-validator/api/handlers/docs"
)

type IServer interface {
	Start()
}

type Server struct {
	HandlerUse handlers.ValidatorUUID
	Port       string
}

func NewServer(handler handlers.Validator, port string) IServer {
	return Server{
		HandlerUse: handler,
		Port:       port,
	}
}

func (s Server) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", httpSwagger.WrapHandler)
	mux.HandleFunc("/validator", s.HandlerUse.Validate)
	routers := cors.AllowAll().Handler(mux)

	log.Println("Server running on port: ", s.Port)
	panic(http.ListenAndServe(":"+s.Port, routers))
}
