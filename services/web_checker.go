package services

import (
	"encoding/csv"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/parmcoder/website-checker-backend/repositories"
)

//go:generate mockgen -destination=mocks/web_checker_service.go -package=mocks

type CheckerService interface {
	PerformCheck(list *[]string) (siteUps int, downs int, duration time.Duration, err error)
	ExtractLinesFromCsv(file *multipart.FileHeader) ([]string, error)
}

type CheckerServiceImpl struct {
	webCheckerRepository repositories.WebCheckerRepository
}

func NewCheckerService(w repositories.WebCheckerRepository) CheckerService {
	return CheckerServiceImpl{
		webCheckerRepository: w,
	}
}

func (c CheckerServiceImpl) ExtractLinesFromCsv(file *multipart.FileHeader) ([]string, error) {
	var lines []string

	src, err := file.Open()
	if err != nil {
		return lines, err
	}

	defer src.Close()

	csvReader := csv.NewReader(src)
	records, err := csvReader.ReadAll()

	if err != nil {
		return lines, err
	}

	for i := range records {
		urlString := records[i][0]
		_, err := url.ParseRequestURI(urlString)
		if err != nil {
			lines = append(lines, "https://"+records[i][0])
		} else {
			lines = append(lines, records[i][0])
		}
	}

	return lines, nil
}

func (c CheckerServiceImpl) PerformCheck(list *[]string) (siteUps int, downs int, duration time.Duration, err error) {
	start := time.Now()

	result, ups := c.webCheckerRepository.ParallelCheck(list)

	siteUps = ups
	downs = len(result) - ups
	duration = time.Since(start)
	duration = time.Duration(duration.Milliseconds())

	return
}
