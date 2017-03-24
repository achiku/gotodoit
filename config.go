package gotodoit

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/achiku/gotodoit/estc"
	"github.com/pkg/errors"
)

// Config app global config
type Config struct {
	Environment    string       `toml:"environment"`
	ServerPort     string       `toml:"server_port"`
	DBName         string       `toml:"db_name"`
	DBUser         string       `toml:"db_user"`
	DBHost         string       `toml:"db_host"`
	DBPass         string       `toml:"db_pass"`
	DBPort         string       `toml:"db_port"`
	DBSSLMode      string       `toml:"db_ssl_mode"`
	PasswordPepper string       `toml:"password_pepper"`
	EstcConfig     *estc.Config `toml:"estc"`
}

// NewConfig creates config
func NewConfig(cfgPath string) (*Config, error) {
	f, err := os.Open(cfgPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open config file: %s", cfgPath)
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config file")
	}
	var config Config
	if err := toml.Unmarshal(buf, &config); err != nil {
		return nil, errors.Wrap(err, "failed to create Config from file")
	}
	return &config, nil
}
