package cmd

import (
	"log"
	"net/http"

	"github.com/Israel-Ferreira/servidor-em-go/pkg/handlers"
	"github.com/Israel-Ferreira/servidor-em-go/pkg/middlewares"
)


func ChainedMiddlewares(f http.HandlerFunc, mddlwrs ...middlewares.Middleware) http.HandlerFunc {
	for _, m :=  range mddlwrs {
		f = m(f)
	}

	return f
}





func StartServer() {
	port := ":9090"

	log.Println("Servidor subindo na porta ", port)
	
	http.Handle("/api/pokego/exclusive-in-raids", ChainedMiddlewares(handlers.GetRaidExclusivePokemons, middlewares.JsonMiddleware(), middlewares.LogMiddleware()))

	if err :=  http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Falha ao subir o servidor http na porta %s \n", port)
	}
}