package controllers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/parmcoder/website-checker-backend/mocks"
)

func TestServerImpl_CheckHealthCsv(t *testing.T) {
	type fields struct {
		checker *mocks.MockCheckerService
	}
	type args struct {
		ctx echo.Context
	}

	path := "../mocks/test.csv"

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", path)
	assert.NoError(t, err)

	sample, err := os.Open(path)
	assert.NoError(t, err)

	_, err = io.Copy(part, sample)
	assert.NoError(t, err)
	assert.NoError(t, writer.Close())

	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/csv", body)

	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())

	response := httptest.NewRecorder()

	ctx := e.NewContext(req, response)

	tests := []struct {
		name    string
		prepare func(*fields)
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			prepare: func(f *fields) {
				f.checker.EXPECT().ExtractLinesFromCsv(gomock.Any()).Return([]string{}, nil)
				f.checker.EXPECT().PerformCheck(gomock.Any()).Return(0, 0, time.Duration(0), nil)
			},
			args:    args{ctx},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := SetupTest(t)

			f := fields{
				checker: mocks.NewMockCheckerService(controller),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			params := ServerImplParams{
				f.checker,
			}

			s := NewServer(params)

			if err := s.CheckHealthCsv(tt.args.ctx); !tt.wantErr(t, err) {
				t.Errorf("ServerImpl.CheckHealthCsv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
