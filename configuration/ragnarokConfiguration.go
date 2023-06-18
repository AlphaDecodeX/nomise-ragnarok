package configuration

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var RagnarokConfig *Configuration

type Configuration struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

func ConfigureRagnarok() error {
	data, err := ioutil.ReadFile("/Users/lovepreetsingh/Desktop/Side/nomise-ragnarok/config.yaml")
	if err != nil {
		return err
	}
	RagnarokConfig := &Configuration{}
	err = yaml.Unmarshal(data, RagnarokConfig)
	if err != nil {
		return err
	}
	return nil
}
