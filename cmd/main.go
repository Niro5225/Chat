package main

import (
	"chat/pkg/apiserver"

	"github.com/sirupsen/logrus"
)

func main() {
	conf := apiserver.New_config()
	s := apiserver.New(conf)

	if err := s.Start(); err != nil {
		logrus.Fatal(err.Error())
	}
}
