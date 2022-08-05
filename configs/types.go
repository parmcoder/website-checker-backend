package configs

import "time"

type WebsiteList struct {
	Rows []string `json:"rows"`
}

type WebsiteListResponse struct {
	Ups      int           `json:"ups"`
	Downs    int           `json:"downs"`
	Duration time.Duration `json:"duration"`
}
