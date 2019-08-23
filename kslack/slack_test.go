package kslack

import (
	"github.com/stretchr/testify/suite"
	"testing"
)



type ReportSample struct {
	Total int
	New int
}

func (r ReportSample) ToSlackReport() Report {
	return Report{}
}

var (
	ReportTotalNotification = ReportSample{}
)
//region Test Slack
type TestSlackSuite struct {
	suite.Suite
	slack Slacker
}


func (s *TestSlackSuite) Test_Report() {
	if err := s.slack.Report(ReportSample{}); err != nil {
		return
	}
}

func (s *TestSlackSuite) Test_Err() {

}

//endregion End test Slack

func TestSlackTestSuite(t *testing.T) {
	suite.Run(t, new(TestSlackSuite))
}

