// Package kslack
package kslack


type Report struct {
	Key string
	Title string `json:"title"`
	Fields []ReportField `json:"fields"`
}

type ReportField struct {
	Title string `json:"title"`
	Data interface{} `json:"data"`
}
