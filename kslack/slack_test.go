package kslack

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)
var (
	ReportTotalNotification = Report{
		Key: "report#totalNotification",
		Title:  "TotalNotification",
		Fields: []ReportField{
			{
				Title: "MarketingNotification",
			},
			{
				Title: "OrderNotification",
			},
		},
	}
)
//region Test Slack
type TestSlackSuite struct {
	suite.Suite
	slack Slacker
}

func (s *TestSlackSuite) SetupTest() {
	fmt.Print("Setup test")
}

func (s *TestSlackSuite) TearUpTest() {
	fmt.Print("Tear up test")
}

func (s *TestSlackSuite) TearDownTest() {
	fmt.Print("Tear down test")
}

func (s *TestSlackSuite) Test_Report() {
	if err := s.slack.Report(ReportTotalNotification); err != nil {
		return
	}
}

func (s *TestSlackSuite) Test_Err() {

}

//endregion End test Slack

func TestSlackTestSuite(t *testing.T) {
	suite.Run(t, new(TestSlackSuite))
}

