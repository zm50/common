package httpcli

import "net/http"

type Response[T any] struct {
	*http.Response

	Body *T
}
