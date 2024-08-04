package httpcli

import (
	"bytes"
	"io"
	"net/http"

	"github.com/pkg/errors"

	"github.com/zm50/common/serialize"
)

func DoRaw[T serialize.Serializer, D any](req *http.Request, body any) (*Response[D], error) {
	serializer := *new(T)

	if body != nil {
		reqData, err := serializer.Marshal(body)
		if err != nil {
			return nil, errors.WithMessage(err, "marshal request body failed")
		}

		req.Body = io.NopCloser(bytes.NewReader(reqData))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.WithMessage(err, "do request failed")
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "read response body failed")
	}

	respBody := new(D)

	err = serializer.Unmarshal(respData, respBody)
	if err != nil {
		return nil, errors.WithMessage(err, "unmarshal response body failed")
	}

	res := &Response[D]{
		Response: resp,
		Body:     respBody,
	}

	return res, nil
}

func Do[T serialize.Serializer, D any](method, url string, body any) (*Response[D], error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "create request failed")		
	}

	res, err := DoRaw[T, D](req, body)
	if err != nil {
		return nil, errors.WithMessage(err, "do request failed")
	}

	return res, nil
}

func Get[T serialize.Serializer, D any](url string, body any) (*Response[D], error) {
	return Do[T, D](http.MethodGet, url, body)
}

func Post[T serialize.Serializer, D any](url string, body any) (*Response[D], error) {
	return Do[T, D](http.MethodPost, url, body)
}

func Head[T serialize.Serializer, D any](url string, body any) (*Response[D], error) {
	return Do[T, D](http.MethodHead, url, body)
}

func Put[T serialize.Serializer, D any](url string, body any) (*Response[D], error) {
	return Do[T, D](http.MethodPut, url, body)
}

func Patch[T serialize.Serializer, D any](url string, body any) (*Response[D], error) {
	return Do[T, D](http.MethodPatch, url, body)
}

func Delete[T serialize.Serializer, D any](url string, body any) (*Response[D], error) {
	return Do[T, D](http.MethodDelete, url, body)
}

func Connect[T serialize.Serializer, D any](url string, body any) (*Response[D], error) {
	return Do[T, D](http.MethodConnect, url, body)
}

func Options[T serialize.Serializer, D any](url string, body any) (*Response[D], error) {
	return Do[T, D](http.MethodOptions, url, body)
}


func Trace[T serialize.Serializer, D any](url string, body any) (*Response[D], error) {
	return Do[T, D](http.MethodTrace, url, body)
}
