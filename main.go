package main

import (
	"github.com/sirupsen/logrus"
	"gitlab.sendo.vn/longnd/spump/kslack"
)

func main() {
	log := logrus.New()
	errorConf := []kslack.ErrorConfig{kslack.ErrorConfig{
		Pattern: "ERROR 1045",
		Min:     1,
	},}
	log.AddHook(kslack.NewSlackrusHook("Notification",errorConf...))
	log.Warnf("log check log for fun %s ", "test")
	log.Errorf("ERROR 1045 (28000): Access denied for user 'foo'@'::1' (using password: NO)")
}


