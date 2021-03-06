package main

import (
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

type Config struct {
	SecretKey string `toml:"cframe_secret"`
	Log       Log    `toml:"log"`
}

type Log struct {
	Days  int64  `toml:"days"`
	Level string `toml:"level"`
	Path  string `toml:"path"`
}

func ParseConfig(path string) (*Config, error) {
	cnt, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = toml.Unmarshal(cnt, &cfg)
	if err != nil {
		return nil, err
	}

	// if len(cfg.Type) == 0 {
	// 	return nil, fmt.Errorf("type MUST configured")
	// }

	// if len(cfg.AccessKey) == 0 {
	// 	cfg.AccessKey = os.Getenv("access_key")
	// }

	// if len(cfg.AccessSecret) == 0 {
	// 	cfg.AccessSecret = os.Getenv("access_secret")
	// }

	// if cfg.AccessKey == "" ||
	// 	cfg.AccessSecret == "" {
	// 	return nil, fmt.Errorf("access_key and secrect_key MUST configured")
	// }

	return &cfg, nil
}
