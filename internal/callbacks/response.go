package callbacks

import (
	"encoding/json"

	"net/http"

	"github.com/go-chi/render"
)

// ApiResponse :
// Common API response object
type ApiResponse struct {
	Payload interface{}

	Session *data.Session
}

// ErrorResponse :
// Error response structure
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`         // user-level status message
	ErrCode    string `json:"code,omitempty"` // application-specific error code
}

// ErrInvalidRequest :
// Invalid request response
// @param : err error
// @return : render.Renderer
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "INVALID_REQUEST",
		ErrCode:        err.Error(),
	}
}

// ErrRender :
// Render the error response
// @param : err error
// @return : render.Renderer
func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "ERROR_RESPONSE",
		ErrCode:        err.Error(),
	}
}

// Render :
// Render error
// Embedded ErrResponse pointer
// @param : w http.ResponseWriter object
// @param : r http.Request object
// @return : error
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// SendResponse :
// send http response in json format
// Embedded ErrResponse pointer
// @param : w http.ResponseWriter object
// @param : r http.Request object
// @param : i interface is the response paylod
// @param : s *data.Session Session object
// @return : void
func SendResponse(w http.ResponseWriter, r *http.Request, i interface{}, s *data.Session) {

	a := ApiResponse{i, s}

	j, err := json.Marshal(a)

	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
