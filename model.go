package nverify

import "time"

type Article struct {
	Title       string
	Description string
	Keywords    []string
	Content     string
	URL         string
	Date        *time.Time
}
