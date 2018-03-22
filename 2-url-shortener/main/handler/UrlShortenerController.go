package handler

import (
	"net/http"
	"../storage"
)

func EncodeHandler(storage storage.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		values, ok := r.URL.Query()["url"]

		if ok {
			w.Write([]byte(storage.Save(values[0])))
		} else {
			w.Write([]byte("Not a valid Url"))
		}
	}

	return http.HandlerFunc(handleFunc)
}

func DecodeHandler(storage storage.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Path[len("/dec/"):]

		url, err := storage.Load(code)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("URL Not Found. Error: " + err.Error() + "\n"))
			return
		}

		w.Write([]byte(url))
	}

	return http.HandlerFunc(handleFunc)
}

func RedirectHandler(storage storage.Storage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Path[len("/red/"):]

		url, err := storage.Load(code)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("URL Not Found. Error: " + err.Error() + "\n"))
			return
		}

		http.Redirect(w, r, string(url), 301)
	}

	return http.HandlerFunc(handleFunc)
}
