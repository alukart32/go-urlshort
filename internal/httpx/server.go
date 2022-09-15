package httpx

import (
	"fmt"
	"net/http"
)

// GetRedirectHandlerFunc redirects the request relative to the url or returns a handler depending on the request path.
func GetRedirectHandlerFunc(paths map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url, isMapped := paths[r.URL.Path]
		if isMapped {
			http.Redirect(w, r, url, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello. This is a home page.")
}

func redirectPath3Handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Redirection for /path3.")
}

// AddHandlerFuncs add a set of http.HandlerFunc to http.ServeMux.
// It will return the passed ServeMux pointer back.
func AddHandlerFuncs(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/redirections/path3", redirectPath3Handler)

	return mux
}
