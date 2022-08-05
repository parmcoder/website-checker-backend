package services

import "github.com/sirupsen/logrus"

type CheckerService interface {
	PerformCheck(list *[]string) error
}

type CheckerServiceImpl struct {
}

func NewCheckerService() CheckerService {
	return CheckerServiceImpl{}
}

func (c CheckerServiceImpl) PerformCheck(list *[]string) error {
	logrus.Info("Checking!")
	return nil
}
