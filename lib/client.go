package lib

import (
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}
