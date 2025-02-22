package config

import (
	"github.com/rollicks-c/configcove"
	"github.com/rollicks-c/configcove/profiles"
	"github.com/rollicks-c/secretblendproviders/envvar"
)

type PortGroup struct {
	Target       Target        `yaml:"target"`
	PortForwards []PortForward `yaml:"portForwards"`
	Name         string        `yaml:"name"`
}

type Target struct {
	K8sConfigFile string `yaml:"k8sConfigFile"`
	K8sContext    string `yaml:"k8sContext"`
}

type PortForward struct {
	Namespace  string `yaml:"namespace"`
	Service    string `yaml:"service"`
	LocalPort  string `yaml:"localPort"`
	RemotePort string `yaml:"remotePort"`
}

type Config struct {
	Groups []PortGroup `yaml:"groups"`
}

func init() {
	_ = envvar.RegisterGlobally()
}

func Profiles() *profiles.Manager[Config] {
	pm := configcove.Profiles[Config](AppName)
	return pm
}
