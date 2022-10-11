package main

import (
	"chat/pkg/apiserver"

	"github.com/sirupsen/logrus"
)

func main() {
	//Создание конфига лдя сервера
	conf := apiserver.New_config()

	//Создание обьекта сервера
	s := apiserver.New(conf)

	if err := s.Start(); err != nil {
		logrus.Fatal(err.Error())
	}
}
