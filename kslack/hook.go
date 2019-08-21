// Package kslack
package kslack

import (
	"fmt"
	"github.com/nlopes/slack"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

const (
	ColorDanger  string = "danger"
	ColorWarning string = "warning"
	ColorGood    string = "good"
)

type Hook struct {
	client     Slacker

	ServiceName string
	AuthorName string
	Channel    string
	IconURL    string
	IconEmoji  string

	ErrorConfigs []ErrorConfig
}

type ErrorConfig struct {
	Pattern string
	Min int
	total int
}

func NewHook(serviceName string, configs ...ErrorConfig) *Hook {
	hook := &Hook{
		ServiceName:serviceName,
		Channel:    channel,
		client:SlackInstance(),
		ErrorConfigs:configs,
	}
	return hook
}

func (hook *Hook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
	}
}

func (hook *Hook) Fire(sourceEntry *logrus.Entry) error {
	fmt.Println("Get in here")
	usedConf := false
	var selectedConf ErrorConfig
	for index, conf := range hook.ErrorConfigs {
		fmt.Println("Source message ", sourceEntry.Message)
		fmt.Println("Pattern ", conf.Pattern)
		if strings.Contains(sourceEntry.Message, conf.Pattern) {
			hook.ErrorConfigs[index].total++
			usedConf = true
			selectedConf = hook.ErrorConfigs[index]
		}
	}

	if !usedConf || selectedConf.Min > selectedConf.total {
		return nil
	}

	priority := buildPriority(sourceEntry.Level)
	attachment :=  slack.Attachment{
		Color:      buildColor(sourceEntry.Level),
		Text:       fmt.Sprintf("%s - %s", hook.ServiceName, sourceEntry.Message),
		Fields : []slack.AttachmentField{priority, buildRepeat(selectedConf)},
		Footer:     "SlackNotification",
		FooterIcon: "https://platform.slack-edge.com/img/default_application_icon.png",
		//Ts: json.Number(time.Now().Unix()),
	}

	notifyOps := buildNotifyLevel(sourceEntry.Level)
	attachmentOps := slack.MsgOptionAttachments(attachment)
	return hook.client.CustomMsg(notifyOps, attachmentOps)
}

func buildNotifyLevel(level logrus.Level) slack.MsgOption {
	switch level {
	case logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel:
		return slack.MsgOptionText("<@channel>", false)
	case logrus.WarnLevel:
		return slack.MsgOptionText("<@channel>", false)
	case logrus.InfoLevel:
		return slack.MsgOptionText("<@channel>", false)
	default:
		return slack.MsgOptionText("<@channel>", false)
	}
}

func buildColor(level logrus.Level) string {
	switch level {
	case logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel:
		return ColorDanger
	case logrus.WarnLevel:
		return ColorWarning
	case logrus.InfoLevel:
		return ColorGood
	default:
		return ""
	}
}

func buildTitle(level logrus.Level) string {
	return "Log notification"
}

func buildPriority(level logrus.Level) slack.AttachmentField {
	switch level {
	case logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel:
		return slack.AttachmentField{
			Title: "Priority",
			Value: "High",
			Short: true,
		}
	case logrus.WarnLevel:
		return slack.AttachmentField{
			Title: "Priority",
			Value: "Normal",
			Short: true,
		}
	case logrus.InfoLevel:
		return slack.AttachmentField{
			Title: "Priority",
			Value: "Low",
			Short: true,
		}
	default:
		return slack.AttachmentField{}
	}
}

func buildRepeat(conf ErrorConfig) slack.AttachmentField {
	val := strconv.Itoa(conf.total)
	return slack.AttachmentField{
		Title: "Repeat",
		Value: val,
		Short: true,
	}
}

