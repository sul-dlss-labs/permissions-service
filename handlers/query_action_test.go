package handlers

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/buger/jsonparser"
	"github.com/stretchr/testify/assert"
)

func TestQueryAction(t *testing.T) {
	r := gofight.New()
	r.GET("/v1/permissions/create/collection").
		SetHeader(gofight.H{
			"On-Behalf-Of": "balbriton",
		}).
		Run(setupFakeRuntime().Handler(),
			func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
				assert.Equal(t, http.StatusOK, r.Code)
				stat, _ := jsonparser.GetBoolean(r.Body.Bytes(), "authorized")
				assert.True(t, stat)
			})
}
