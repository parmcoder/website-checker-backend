package controllers

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/parmcoder/website-checker-backend/services"
)

func TestServerImpl_CheckHealthCsv(t *testing.T) {
	type fields struct {
		checker *services.CheckerService
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServerImpl{
				checker: tt.fields.checker,
			}
			if err := s.CheckHealthCsv(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("ServerImpl.CheckHealthCsv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServerImpl_CheckHealth(t *testing.T) {
	type fields struct {
		checker *services.CheckerService
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ServerImpl{
				checker: tt.fields.checker,
			}
			if err := s.CheckHealth(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("ServerImpl.CheckHealth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
