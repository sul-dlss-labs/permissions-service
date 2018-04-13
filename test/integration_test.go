package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/sul-dlss-labs/permissions-service/config"
	baloo "gopkg.in/h2non/baloo.v3"
)

func setupTest() *baloo.Client {
	remoteHost, ok := os.LookupEnv("TEST_REMOTE_ENDPOINT")
	if !ok {
		port := config.NewConfig().Port
		remoteHost = fmt.Sprintf("localhost:%v", port)
	}
	return baloo.New(fmt.Sprintf("http://%s", remoteHost))
}

func TestHealthCheck(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	setupTest().Get("/v1/healthcheck").
		Expect(t).
		Status(200).
		Type("json").
		JSON(map[string]string{"status": "OK"}).
		Done()
}
