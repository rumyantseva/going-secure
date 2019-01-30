package pprofed

import (
	"net/http"
	_ "net/http/pprof" // add pprof handlers
)

func serve(addr string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// provide some business logic
	})

	http.ListenAndServe(addr, nil)
}
