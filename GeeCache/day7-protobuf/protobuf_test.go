package main

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPB(t *testing.T) {
	for i := 0; i < 100; i++ {
		resp, _ := http.Get("http://localhost:9999/api?key=Tom")
		content, _ := ioutil.ReadAll(resp.Body)
		require.Equal(t, "630", string(content))
		resp.Body.Close()
	}
}
