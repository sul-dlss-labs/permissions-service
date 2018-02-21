package handlers

import (
	"net/http"

	app "github.com/sul-dlss-labs/swagger-go-template"
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
