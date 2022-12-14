package repository

import "github.com/jinzhu/configor"

type Config struct {
	AppName string `default:"ami-back"`
	Port    string `default:"8000"`
	DB      struct {
		Use   string `default:"mysql"`
		Mysql []struct {
			Enabled  bool   `default:"true"`
			Host     string `default:"localhost"`
			Port     string `default:"3306"`
			UserName string `default:"root"`
			Password string `default:"MyPassword@123"`
			Database string `default:"mb_altera"`
		}
	}
}

func (c *Config) NewConfig() (*Config, error) {
	err := configor.Load(c, "config.yml")
	if err != nil {
		return nil, err
	}
	return c, nil
}
