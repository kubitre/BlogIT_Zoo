package Config

import (
	"log"

	"github.com/BurntSushi/toml"
)

/*Configuration - it is struct for parsing settings from .toml file*/
type Configuration struct {
	Server   string
	Database string
}

func (c *Configuration) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
