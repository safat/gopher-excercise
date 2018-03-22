package main

import (
	"runtime"

	"github.com/mitchellh/go-homedir"
	"./storage"
	"./handler"
	"log"
	"path/filepath"
	"net/http"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	dir, _ := homedir.Dir()

	storage := &storage.FileSystem{}

	err := storage.Init(filepath.Join(dir, "url-shortener-dir"))

	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is response"));
	}))

	http.Handle("/encode", handler.EncodeHandler(storage))
	http.Handle("/dec/", handler.DecodeHandler(storage))
	http.Handle("/red/", handler.RedirectHandler(storage))

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
