package services

import (
	"encoding/csv"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/smirzaei/parallel"
)

type CheckerService interface {
	PerformCheck(list *[]string) (siteUps int, downs int, duration time.Duration, err error)
	ExtractLinesFromCsv(file *multipart.FileHeader) ([]string, error)
}

type CheckerServiceImpl struct {
}

func NewCheckerService() CheckerService {
	return CheckerServiceImpl{}
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

	result := parallel.Map(*list, checkLink)

	c1 := make(chan int)
	c2 := make(chan int)

	go sum(result[:len(result)/2], c1)
	go sum(result[len(result)/2:], c2)

	x, y := <-c1, <-c2 // receive from c1 aND C2

	siteUps = x + y
	downs = len(result) - x - y
	duration = time.Since(start)
	duration = time.Duration(duration.Milliseconds())
	return
}

func checkLink(link string) int {
	_, err := http.Get(link)
	if err != nil {
		logrus.Debug(link, " might be down!")
		return 0
	}
	logrus.Debug(link, " is up!")
	return 1
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}
