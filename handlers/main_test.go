package handlers

import (
	"net/http"

	app "github.com/sul-dlss-labs/permissions-service"
)

func setupFakeRuntime() *TestEnv {
	return &TestEnv{}
}

type TestEnv struct {
}

func (d *TestEnv) Handler() http.Handler {
	// For testing we can inject a fake config/services here.
	rt, _ := app.NewRuntime(nil)
	return BuildAPI(rt).Serve(nil)
}
