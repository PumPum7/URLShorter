package utils

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		Domain   string `yaml:"domain"`
		DBName   int    `yaml:"database"`
	} `yaml:"database"`
	Logging struct {
		Logfile  string `yaml:"logfile"`
		LogLevel string `yaml:"loglevel"`
	} `yaml:"log"`
}

var Cfg Config

func init() {
	f, err := os.Open("config.yml")
	if err != nil {
		Log.Fatal(err, " Config.yml file not found")
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Cfg)
	if err != nil {
		Log.Fatal(err)
	}
}
