package services

import (
	"mime/multipart"
	"reflect"
	"testing"
	"time"

	"github.com/parmcoder/website-checker-backend/repositories"
)

func TestNewCheckerService(t *testing.T) {
	type args struct {
		w repositories.WebCheckerRepository
	}
	tests := []struct {
		name string
		args args
		want CheckerService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCheckerService(tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCheckerService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckerServiceImpl_ExtractLinesFromCsv(t *testing.T) {
	type fields struct {
		webCheckerRepository repositories.WebCheckerRepository
	}
	type args struct {
		file *multipart.FileHeader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CheckerServiceImpl{
				webCheckerRepository: tt.fields.webCheckerRepository,
			}
			got, err := c.ExtractLinesFromCsv(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckerServiceImpl.ExtractLinesFromCsv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckerServiceImpl.ExtractLinesFromCsv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckerServiceImpl_PerformCheck(t *testing.T) {
	type fields struct {
		webCheckerRepository repositories.WebCheckerRepository
	}
	type args struct {
		list *[]string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantSiteUps  int
		wantDowns    int
		wantDuration time.Duration
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CheckerServiceImpl{
				webCheckerRepository: tt.fields.webCheckerRepository,
			}
			gotSiteUps, gotDowns, gotDuration, err := c.PerformCheck(tt.args.list)
			if (err != nil) != tt.wantErr {
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
