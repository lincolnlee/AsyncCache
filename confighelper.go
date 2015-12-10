package AsyncCache

import (
	"github.com/astaxie/beego/config"
	_ "github.com/astaxie/beego/config/yaml"
)

type configHelper struct {
	Configurator config.ConfigContainer
}

func newConfig() config.ConfigContainer {
	cnf, err := config.NewConfig("yaml", "config.yaml")
	if err != nil {
		panic(err)
	}
	return cnf
}
