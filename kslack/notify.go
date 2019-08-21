// Package kslack
package kslack


type NotifyConfig struct {
	Pattern string `json:"pattern"`
	NotifyLevel SlackNotifyLevel
	NotifyTarget SlackNotifyTarget
	Min int `json:"min"`
	Max int `json:"max"`
}

type slackNotifyConfig struct {
	Pattern string
	RepeatCount int

	NotifyLevel int
	NotifyUserId string
	NotifyChannel string
}

func NewSlackNotifyConfig(pattern string, repeatCount int, level int, notifyTarget...string) *slackNotifyConfig  {
	return &slackNotifyConfig{

	}
}
