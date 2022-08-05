package repositories

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/smirzaei/parallel"
)

//go:generate mockgen -destination=mocks/web_checker_repositories.go -package=mocks

type WebCheckerRepository interface {
	ParallelCheck(list *[]string) ([]int, int)
}

type WebCheckerRepositoryImpl struct {
}

func NewWebCheckerRepository() WebCheckerRepository {
	return WebCheckerRepositoryImpl{}
}

func (w WebCheckerRepositoryImpl) ParallelCheck(list *[]string) ([]int, int) {
	result := parallel.Map(*list, checkLink)

	c1 := make(chan int)
	c2 := make(chan int)

	go sum(result[:len(result)/2], c1)
	go sum(result[len(result)/2:], c2)

	x, y := <-c1, <-c2

	return result, x + y
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
