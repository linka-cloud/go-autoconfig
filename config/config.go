package config

import (
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

type Config struct {
	ServiceAddr string `yaml:"service_addr" mapstructure:"service_addr"`

	Domain string `yaml:"domain" mapstructure:"domain" validate:"required"`

	IMAP Server `yaml:"imap" mapstructure:"imap"`
	SMTP Server `yaml:"smtp" mapstructure:"smtp"`
}

func (c Config) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

type Server struct {
	Server     string `yaml:"server" mapstructure:"server" validate:"required"`
	Port     int    `yaml:"port" mapstructure:"port" validate:"required"`
	STARTTLS bool   `yaml:"starttls" mapstructure:"starttls"`
}

const defaultConfig = `
service_addr: ":1323"
domain: ""
imap:
  server: ""
  port: 993
  starttls: false
smtp:
  server: ""
  port: 587
  starttls: true
`

func New(path string) (*Config, error) {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(strings.NewReader(defaultConfig)); err != nil {
		return nil, err
	}
	path = strings.TrimSpace(path)
	if path != "" {
		viper.SetConfigFile(path)
		if err := viper.MergeInConfig(); err != nil {
			return nil, err
		}
	}
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	if err := c.Validate(); err != nil {
		return nil, err
	}
	if c.ServiceAddr == "" {
		c.ServiceAddr = ":1323"
	}
	return &c, nil
}
