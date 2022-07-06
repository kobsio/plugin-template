package instance

import (
	"github.com/mitchellh/mapstructure"
)

type Config struct {
	Greeting string `json:"string"`
}

type Instance interface {
	GetName() string
	GetGreeting() string
}

type instance struct {
	name   string
	config Config
}

func (i *instance) GetName() string {
	return i.name
}

func (i *instance) GetGreeting() string {
	return i.config.Greeting
}

func New(name string, options map[string]any) (Instance, error) {
	var config Config
	err := mapstructure.Decode(options, &config)
	if err != nil {
		return nil, err
	}

	return &instance{
		name:   name,
		config: config,
	}, nil
}
