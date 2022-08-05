package controllers

import (
	"reflect"
	"testing"

	"github.com/parmcoder/website-checker-backend/repositories"
	"github.com/parmcoder/website-checker-backend/services"
)

func TestNewServer(t *testing.T) {
	type testSetup struct {
		name string
		args ServerImplParams
		want Server
	}

	service := services.NewCheckerService(repositories.NewWebCheckerRepository())

	server := &ServerImpl{
		&service,
	}

	sParam := ServerImplParams{
		Checker: service,
	}

	tests := []testSetup{
		{
			name: "Success",
			args: sParam,
			want: server,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := NewServer(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
