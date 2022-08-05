package controllers

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func SetupTest(t *testing.T) *gomock.Controller {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	return mockCtrl
}
