package command

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	data map[string][]byte
	err  int
}

func (conf Config) Get(key string) []byte {
	return conf.data[key]
}

func (conf Config) Dig(key string) (Config, error) {
	var out Config

	if err := json.Unmarshal(conf.Get(key), out); err != nil {
		return out, err
	} else {
		return out, nil
	}
}

var ConfigInstance Config

func init() {
	config, err := load()

	if err != nil {
		logrus.Errorln(err)

		config = Config{
			err: 1,
		}
	}

	ConfigInstance = config
}

func load() (config Config, err error) {
	ex, err := os.Executable()

	if err != nil {
		return config, err
	}

	if bytes, err := ioutil.ReadFile(filepath.Join(filepath.Dir(ex), ".dpg", "config")); err != nil {
		return config, err
	} else if err := json.Unmarshal(bytes, config); err != nil {
		return config, err
	} else {
		return config, nil
	}
}
