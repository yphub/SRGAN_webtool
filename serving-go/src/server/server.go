package main

import "net/http"

var config *Config

func main() {
	var err error
	config, err = LoadConfig("config.json")
	if err != nil {
		panic(err)
	}

	err = InitTfGrpc(config.TFServing.Address, config.TFServing.ModelName, config.TFServing.SigName)
	if err != nil {
		panic(err)
	}
	router := getMainRouter()
	if config.StaticFile != "" {
		router.ServeFiles("/*filepath", http.Dir(config.StaticFile))
	}

	err = http.ListenAndServe(config.HTTPAddress, router)
	if err != nil {
		panic(err)
	}
}
