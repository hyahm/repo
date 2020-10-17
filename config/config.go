package config

import "time"

//

type Config struct {
	Depend []string          `yaml:"depend"`
	Env    map[string]string `yaml:"env"`
	Script Script            `yaml:"script"`
}

type Script struct {
	Name               string            `yaml:"name" json:"name"`
	Command            string            `yaml:"command" json:"command"`
	Dir                string            `yaml:"dir" json:"dir"`
	Replicate          int               `yaml:"replicate" json:"replicate"`
	Always             bool              `yaml:"always" json:"always"`
	DisableAlert       bool              `yaml:"disableAlert" json:"disableAlert"`
	Env                map[string]string `yaml:"env" json:"env"`
	ContinuityInterval time.Duration     `yaml:"continuityInterval" json:"continuityInterval"`
	Port               int               `yaml:"port" json:"port"`
	AT                 AlertTo           `yaml:"alert" json:"alert"`
	KillTime           time.Duration     `yaml:"killTime" json:"killTime"`
	Version            string            `yaml:"version" json:"version"`
}

type AlertTo struct {
	Email  []string `yaml:"email"`
	Rocket []string `yaml:"rocket"`
}
