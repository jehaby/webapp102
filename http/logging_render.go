package http

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/jehaby/webapp102/pkg/log"
)

func loggingRespond(l *log.Logger) func(w http.ResponseWriter, r *http.Request, v interface{}) {
	return func(w http.ResponseWriter, r *http.Request, v interface{}) {

		switch err := v.(type) {

		case *ErrResponse:

			switch err.HTTPStatusCode {
			case http.StatusNotFound, http.StatusUnauthorized, http.StatusBadRequest:
				break
			default:
				l.Errorw("bad error", "error", err.Err, "http_code", err.HTTPStatusCode)
			}

			v = map[string][]interface{}{"errors": {v}}

		case error:
			l.Errorw("unknown error", "error", err)
			v = map[string][]interface{}{"errors": {v}}
		}

		render.DefaultResponder(w, r, v)
	}
}
