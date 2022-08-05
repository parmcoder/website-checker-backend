package services

import (
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/parmcoder/website-checker-backend/mocks"
	"github.com/parmcoder/website-checker-backend/repositories"
)

func TestNewCheckerService(t *testing.T) {
	type args struct {
		w repositories.WebCheckerRepository
	}
	type testSetup struct {
		name string
		args args
		want CheckerService
	}
	controller := SetupTest(t)
	repo := mocks.NewMockWebCheckerRepository(controller)
	arg := args{repo}
	service := CheckerServiceImpl{
		webCheckerRepository: repo,
	}

	tests := []testSetup{
		{
			name: "Success",
			args: arg,
			want: service,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCheckerService(tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCheckerService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckerServiceImpl_PerformCheck(t *testing.T) {
	type fields struct {
		webCheckerRepository *mocks.MockWebCheckerRepository
	}
	type args struct {
		list *[]string
	}

	var webs []string
	arg := args{
		&webs,
	}

	tests := []struct {
		name         string
		prepare      func(*fields)
		args         args
		wantSiteUps  int
		wantDowns    int
		wantDuration time.Duration
		wantErr      assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{name: "Success",
			prepare: func(f *fields) {
				f.webCheckerRepository.EXPECT().ParallelCheck(gomock.Any()).Return([]int{}, 0)
			},
			args:         arg,
			wantSiteUps:  0,
			wantDowns:    0,
			wantDuration: 0,
			wantErr:      assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := SetupTest(t)
			f := fields{
				webCheckerRepository: mocks.NewMockWebCheckerRepository(controller),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}

			c := CheckerServiceImpl{
				webCheckerRepository: f.webCheckerRepository,
			}

			gotSiteUps, gotDowns, gotDuration, err := c.PerformCheck(tt.args.list)
			if !tt.wantErr(t, err) {
				t.Errorf("CheckerServiceImpl.PerformCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSiteUps != tt.wantSiteUps {
				t.Errorf("CheckerServiceImpl.PerformCheck() gotSiteUps = %v, want %v", gotSiteUps, tt.wantSiteUps)
			}
			if gotDowns != tt.wantDowns {
				t.Errorf("CheckerServiceImpl.PerformCheck() gotDowns = %v, want %v", gotDowns, tt.wantDowns)
			}
			if gotDuration != tt.wantDuration {
				t.Errorf("CheckerServiceImpl.PerformCheck() gotDuration = %v, want %v", gotDuration, tt.wantDuration)
			}
		})
	}
}

func TestCheckerServiceImpl_ExtractLinesFromCsv(t *testing.T) {
	type fields struct {
		webCheckerRepository repositories.WebCheckerRepository
	}
	type args struct {
		records [][]string
	}

	var records [][]string
	var results []string
	results = append(results, "")
	resultsOutput := []string{"https://"}
	records = append(records, results)
	arg := args{
		records,
	}
	type testSetup struct {
		name    string
		prepare func(*fields)
		args    args
		want    []string
		wantErr assert.ErrorAssertionFunc
	}
	tests := []testSetup{
		// TODO: Add test cases.
		{
			name: "Success",
			prepare: func(f *fields) {

			},
			args:    arg,
			want:    resultsOutput,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := SetupTest(t)
			f := fields{
				webCheckerRepository: mocks.NewMockWebCheckerRepository(controller),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			c := NewCheckerService(f.webCheckerRepository)
			got, err := c.ExtractLinesFromCsv(tt.args.records)
			if !tt.wantErr(t, err) {
				t.Errorf("CheckerServiceImpl.ExtractLinesFromCsv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckerServiceImpl.ExtractLinesFromCsv() = %v, want %v", got, tt.want)
			}
		})
	}
}
