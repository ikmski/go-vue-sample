package api

import (
	"time"
)

var timeFormat string = "2006-01-02 15:04:05"
var jst *time.Location = time.FixedZone("Asia/Tokyo", 9*60*60)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
