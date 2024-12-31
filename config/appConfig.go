package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	ServerPort            string
	Dsn                   string
	AppSecret             string
	TwilioAccountSid      string
	TwilioAuthToken       string
	TwilioFromPhoneNumber string
}

func SetupEnv() (cfg AppConfig, err error) {

	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load(".env")
	}

	httpPort := os.Getenv("HTTP_PORT")

	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variable HTTP_PORT is not set")
	}

	Dsn := os.Getenv("DSN")

	if len(Dsn) < 1 {
		return AppConfig{}, errors.New("env variable DSN is not set")
	}

	appSecret := os.Getenv("APP_SECRET")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("env variable APP_SECRET is not set")
	}

	TwilioAccountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	if len(TwilioAccountSid) < 1 {
		return AppConfig{}, errors.New("env variable TWILIO_ACCOUNT_SID is not set")
	}

	TwilioAuthToken := os.Getenv("TWILIO_AUTH_TOKEN")
	if len(TwilioAuthToken) < 1 {
		return AppConfig{}, errors.New("env variable TWILIO_AUTH_TOKEN is not set")
	}

	TwilioFromPhoneNumber := os.Getenv("TWILIO_FROM_PHONE_NUMBER")
	if len(TwilioFromPhoneNumber) < 1 {
		return AppConfig{}, errors.New("env variable TWILIO_FROM_PHONE_NUMBER is not set")
	}

	return AppConfig{
		ServerPort: httpPort, Dsn: Dsn, AppSecret: appSecret,
		TwilioAccountSid: TwilioAccountSid, TwilioAuthToken: TwilioAuthToken,
		TwilioFromPhoneNumber: TwilioFromPhoneNumber}, nil

}
