package main

import (
	"log"
	"net/http"
)

func main() {
	//HTTP - Protocolo de comunicacao - Base da comunicacao WEB

	//Cliente (Faz Requisicao) - Servidor (Processa requisicao e envia a resposta)
	//Rotas

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pagina Raiz"))
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ola Mundo!"))
	})

    http.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregando usuario"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
