// Package kslack
package kslack

import (
	"fmt"
	"github.com/nlopes/slack"
)

var (
	apiKey  = ""
	channel = ""
)

type Slacker interface {
	//Cheat for testing purpose
	Error(msg string, args ...interface{})
	CustomMsg(ops ...slack.MsgOption) error
	Msg(attach slack.Attachment) error
	Report(report Report) error
}

type slacker struct {
	userId uint
	api    *slack.Client
}

func (c *slacker) Report(report Report) error {
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

func (c *slacker) Msg(attach slack.Attachment) error {

	_, _, err := c.api.PostMessage(
		channel,
		slack.MsgOptionAttachments(attach),
		)
	return err
}
func (c *slacker) Error(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args)
	}

	if _, _, err := c.api.PostMessage(
		channel,
		slack.MsgOptionText(msg, false),
		slack.MsgOptionBroadcast(),
	); err != nil {
		fmt.Println("err ", err.Error())
	}

}
func (c *slacker) Info(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args)
	}

	if _, _, err := c.api.PostMessage(
		channel,
		slack.MsgOptionText(msg, false),
		slack.MsgOptionBroadcast(),
	); err != nil {
		fmt.Println("err ", err.Error())
	}

}
func (c *slacker) Debug(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args)
	}

	if _, _, err := c.api.PostMessage(
		channel,
		slack.MsgOptionText(msg, false),
		slack.MsgOptionBroadcast(),
	); err != nil {
		fmt.Println("err ", err.Error())
	}

}
func (c *slacker) Warn(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args)
	}

	if _, _, err := c.api.PostMessage(
		channel,
		slack.MsgOptionText(msg, false),
		slack.MsgOptionBroadcast(),
	); err != nil {
		fmt.Println("err ", err.Error())
	}

}
func (c *slacker) Critical(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args)
	}

	if _, _, err := c.api.PostMessage(
		channel,
		slack.MsgOptionText(msg, false),
		slack.MsgOptionBroadcast(),
	); err != nil {
		fmt.Println("err ", err.Error())
	}

}

func (c *slacker) SetDev() {
	apiKey = "xoxp-685696288598-672151870147-683127053956-191603b65e16df541fa64a495e7e1fe0"
	channel = "#reporter"
}

func (c *slacker) SetProd() {
	apiKey = "xoxp-158475431284-672224707747-701240168402-8afa211fb45e823da4d6e16d7ca40b66"
	channel = "#messaging_report"
}

func SlackInstance() *slacker {
	/*
		BMC_KEY: xoxp-158475431284-672224707747-701240168402-8afa211fb45e823da4d6e16d7ca40b66
		PRIVATE_KEY: xoxp-685696288598-672151870147-683127053956-191603b65e16df541fa64a495e7e1fe0
	*/

	slacker := &slacker{}
	slacker.SetDev()
	slacker.api = slack.New(apiKey, slack.OptionDebug(true))
	return slacker
}

type CSlack interface {
	Report()

	Error()
	Warn()
	Critical()
}

type cSlack struct {

}


func NewSlacker(conf slackNotifyConfig) CSlack {
	return nil
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



