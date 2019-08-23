// Package kslack
package kslack


// Report object should implement this interface
type Reportable interface {
	ToSlackReport() Report
}

type Report struct {
	Key string
	Title string `json:"title"`
	Fields []ReportField `json:"fields"`
}

type ReportField struct {
	Title string `json:"title"`
	Data interface{} `json:"data"`
}
