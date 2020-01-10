package mises_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Handler_Alive(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	// return code & no errors
	res, err := http.Get(srv.URL + "/v1/alive")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// respnse body
	type response struct {
		Alive bool
	}

	data := &response{}
	err = json.NewDecoder(res.Body).Decode(data)
	assert.Equal(t, data.Alive, true)
}
