package main

import "github.com/sirupsen/logrus"

func main() {
	logger := logrus.WithField("some_key", "some_value")
	logger.Println("printing some args")
}
