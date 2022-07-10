package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	EmailConfig
	ListenAddress string `json:"listen_address"`
}

type EmailConfig struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     int    `json:"email_port"`
}

var Conf *Config

func ReadConfig(source string) (err error) {
	os.Stat(".env")
	if _, err := os.Stat(".env"); !(errors.Is(err, os.ErrNotExist)) {
		err1 := godotenv.Load()
		if err1 != nil {
			return err1
		}
	}

	var raw []byte
	raw, err = ioutil.ReadFile(source)
	if err != nil {
		eMsg := "Couldn't read configuration"
		logrus.WithError(err).Error(eMsg)
		return err
	}

	emailConfig := EmailConfig{}
	err = json.Unmarshal(raw, &emailConfig)
	if err != nil {
		eMsg := "Couldn't parse config file from json"
		logrus.WithError(err).Error(eMsg)
		return err
	}

	Conf = &Config{
		ListenAddress: os.Getenv("LISTEN_ADDR"),
		EmailConfig:   emailConfig,
	}

	return
}
