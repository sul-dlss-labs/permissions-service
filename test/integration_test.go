package test

import (
	"fmt"
	"testing"

	"github.com/sul-dlss-labs/permissions-service/config"
	baloo "gopkg.in/h2non/baloo.v3"
)

func setupTest() *baloo.Client {
	port := config.NewConfig().Port
	return baloo.New(fmt.Sprintf("http://localhost:%v", port))
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
