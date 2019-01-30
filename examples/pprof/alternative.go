package pprofed

import (
	"net/http"
	"net/http/pprof"

	"github.com/gorilla/mux"
)

func diagnostics(addr string) error {
	r := mux.NewRouter()

	r.HandleFunc("/diag/pprof", pprof.Index)
	r.HandleFunc("/diag/cmdline", pprof.Cmdline)
	r.HandleFunc("/diag/profile", pprof.Profile)
	r.HandleFunc("/diag/symbol", pprof.Symbol)
	r.HandleFunc("/diag/trace", pprof.Trace)
	r.Handle("/diag/goroutine", pprof.Handler("goroutine"))
	r.Handle("/diag/heap", pprof.Handler("heap"))
	r.Handle("/diag/threadcreate", pprof.Handler("threadcreate"))
	r.Handle("/diag/block", pprof.Handler("block"))

	return http.ListenAndServe(addr, r)
}

// OMIT
func external(addr string) error {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// provide some business logic
	})

	return http.ListenAndServe(addr, r)
}

// OMIT
