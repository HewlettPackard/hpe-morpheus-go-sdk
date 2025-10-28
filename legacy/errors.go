package morpheus

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func errWithBody(err error, resp *http.Response) error {
	var msg string

	if err != nil {
		msg = err.Error()
	}

	if resp != nil {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		code := http.StatusText(resp.StatusCode)
		msg = fmt.Sprintf("%s (%s): %s", msg, code, string(bodyBytes))
	}

	return errors.New(msg)
}
