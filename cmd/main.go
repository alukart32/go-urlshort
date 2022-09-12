package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"alukart32.com/urlshort/internal/config"
	"alukart32.com/urlshort/internal/handler"
)

var (
	input    = flag.String("input", "yaml", "The input to use between 'yaml/json' files or read from db")
	filepath = flag.String("filepath", "../assets/config.yaml", "The filepath of the yaml/json/... config file")
)

func main() {
	flag.Parse()

	reader, err := config.NewFileConfigReader(*input)
	if err != nil {
		log.Fatal(err)
	}
	conf, err := config.GetConfig(reader, *input, *filepath)
	if err != nil {
		log.Fatal(err)
	}

	conf.Print()

	fmt.Println("\ninit a new http handler...")
	mux := handler.AddHandlerFuncs(http.NewServeMux())
	handler := handler.GetRedirectHandlerFunc(conf.PathsToMap(), mux)

	fmt.Println("start http server at 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
