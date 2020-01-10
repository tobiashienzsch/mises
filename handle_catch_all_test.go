package mises_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Handler_CatchAll(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	res, err := http.Get(srv.URL + "/v1/not_implemented")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)

	// respnse body
	type response struct {
		Error string
	}

	data := &response{}
	err = json.NewDecoder(res.Body).Decode(data)
	assert.Equal(t, "endpoint not found", data.Error)

}
