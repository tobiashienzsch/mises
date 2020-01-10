package mises_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Handler_Index(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	res, err := http.Get(srv.URL + "/v1/")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// respnse body
	type response struct {
		Name string
	}

	data := &response{}
	err = json.NewDecoder(res.Body).Decode(data)
	assert.Equal(t, "test_service", data.Name)

}
