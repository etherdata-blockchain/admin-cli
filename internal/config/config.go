package config

import "log"

type Config struct {
	//DefaultNodeId is the node id for this client
	DefaultNodeId string
	//DefaultPassword is the etd client password to connect the admin
	DefaultPassword string
	//DefaultEndpoint is the etd admin endpoint url
	DefaultEndpoint   string
	DefaultTemplateId string
}

//NewConfig will create a selection object
func NewConfig() Config {
	return Config{}
}

func (c *Config) Print() {
	log.Printf("DefaultNodeId %s", c.DefaultNodeId)
	log.Printf("DefaultPassword %s", c.DefaultPassword)
	log.Printf("DefaultEndpoint %s", c.DefaultEndpoint)
	log.Printf("DefaultTemplateId %s", c.DefaultTemplateId)

}
