package commands

import (
	"dao/shadow/config"
)


var Cfg *config.Config

func Setup(c *config.Config) {
	Cfg = c
}