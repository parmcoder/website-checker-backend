package services

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/smirzaei/parallel"
)

type CheckerService interface {
	PerformCheck(list *[]string) (siteUps int, downs int, duration time.Duration, err error)
}

type CheckerServiceImpl struct {
}

func NewCheckerService() CheckerService {
	return CheckerServiceImpl{}
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
		logrus.Info(link, " might be down!")
		return 0
	}
	logrus.Info(link, " is up!")
	return 1
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}
