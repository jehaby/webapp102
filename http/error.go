package http

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"

	"github.com/jehaby/webapp102/service"
)

var (

	// for cases when user must be logged in but is not;
	errNotLoggedIn500 = err500(errors.New("user not logged in"))
)

//--
// Error response payloads & renderers
//--

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	AppCode int64  `json:"code,omitempty"` // application-specific error code TODO: add
	Msg     string `json:"msg"`            // what the user will see
}

func (a *app) createRendererErr(err error) render.Renderer {
	// TODO: ok (status 200) response if err == nil
	// TODO: metrics maybe (or middleware?)

	switch err.(type) {
	case *service.ErrBadRequest:
		return errInvalidRequest(err)
	}

	switch err {
	case service.ErrNotFound:
		return errNotFound(err)
	case service.ErrNotAuthorized, service.ErrNotAllowed:
		return errUnauthorized(err)
	default:
		return err500(err)
	}
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func errInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		Msg:            err.Error(),
	}
}

func errUnauthorized(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 401,
		Msg:            "Unauthorized",
	}
}

func errRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		Msg:            "Error rendering response.",
	}
}

func errNotFound(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 404,
		Msg:            "Not found",
	}
}

func err500(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		Msg:            "Service error",
	}
}
