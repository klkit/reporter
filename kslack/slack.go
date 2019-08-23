// Package kslack
package kslack

import (
	"fmt"
	"github.com/nlopes/slack"
	"os"
)

var (
	apiKey  = ""
	channel = ""
)

type Slacker interface {
	CustomMsg(ops ...slack.MsgOption) error
	Report(reportable Reportable) error
}

type slacker struct {
	userId uint
	api    *slack.Client
}

func (c *slacker) Report(reportable Reportable) error {
	report := reportObj.ToSlackReport()
	headerText := slack.NewTextBlockObject("plain_text", report.Title, false, false)
	headerSection := slack.NewSectionBlock(headerText, []*slack.TextBlockObject{}, nil)

	// Fields
	fieldSlice := make([]*slack.TextBlockObject, 0)
	for _, field := range report.Fields {
		typeField := slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("*%s:*\n %+v", field.Title, field.Data), false, false)
		fieldSlice = append(fieldSlice, typeField)
	}

	fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)

	blocks := slack.MsgOptionBlocks(
		headerSection,
		slack.NewDividerBlock(),
		fieldsSection,
	)

	if _, _, err := c.api.PostMessage(
		channel,
		blocks,
	); err != nil {
		fmt.Println("err ", err.Error())
	}
	return nil
}

func (c *slacker) CustomMsg(ops ...slack.MsgOption) error {
	_, _, err := c.api.PostMessage(
		channel,
		ops...,
	)
	return err
}

//func (c *slacker) Msg(attach slack.Attachment) error {
//
//	_, _, err := c.api.PostMessage(
//		channel,
//		slack.MsgOptionAttachments(attach),
//		)
//	return err
//}
//func (c *slacker) Error(msg string, args ...interface{}) {
//	if len(args) > 0 {
//		msg = fmt.Sprintf(msg, args)
//	}
//
//	if _, _, err := c.api.PostMessage(
//		channel,
//		slack.MsgOptionText(msg, false),
//		slack.MsgOptionBroadcast(),
//	); err != nil {
//		fmt.Println("err ", err.Error())
//	}
//
//}
//func (c *slacker) Info(msg string, args ...interface{}) {
//	if len(args) > 0 {
//		msg = fmt.Sprintf(msg, args)
//	}
//
//	if _, _, err := c.api.PostMessage(
//		channel,
//		slack.MsgOptionText(msg, false),
//		slack.MsgOptionBroadcast(),
//	); err != nil {
//		fmt.Println("err ", err.Error())
//	}
//
//}
//func (c *slacker) Debug(msg string, args ...interface{}) {
//	if len(args) > 0 {
//		msg = fmt.Sprintf(msg, args)
//	}
//
//	if _, _, err := c.api.PostMessage(
//		channel,
//		slack.MsgOptionText(msg, false),
//		slack.MsgOptionBroadcast(),
//	); err != nil {
//		fmt.Println("err ", err.Error())
//	}
//
//}
//func (c *slacker) Warn(msg string, args ...interface{}) {
//	if len(args) > 0 {
//		msg = fmt.Sprintf(msg, args)
//	}
//
//	if _, _, err := c.api.PostMessage(
//		channel,
//		slack.MsgOptionText(msg, false),
//		slack.MsgOptionBroadcast(),
//	); err != nil {
//		fmt.Println("err ", err.Error())
//	}
//
//}
//func (c *slacker) Critical(msg string, args ...interface{}) {
//	if len(args) > 0 {
//		msg = fmt.Sprintf(msg, args)
//	}
//
//	if _, _, err := c.api.PostMessage(
//		channel,
//		slack.MsgOptionText(msg, false),
//		slack.MsgOptionBroadcast(),
//	); err != nil {
//		fmt.Println("err ", err.Error())
//	}
//
//}

func SlackInstance() *slacker {
	slacker := &slacker{}
	apiKey = os.Getenv("SLACK_API_KEY")
	channel = os.Getenv("SLACK_CHANNEL")
	fmt.Println("Key ", apiKey)
	fmt.Println("Key ", apiKey)
	slacker.api = slack.New(apiKey, slack.OptionDebug(true))
	return slacker
}

//region Object
type SlackRunMode int
const (
	DebugMode      SlackRunMode = 1
	DevelopMode    SlackRunMode = 2
	ProductionMode SlackRunMode = 3
)

type SlackNotifyTarget string
type SlackNotifyLevel int
const(
	NotifyUser    SlackNotifyLevel = 1
	NotifyChannel SlackNotifyLevel = 2
)

//endregion Object



