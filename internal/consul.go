package internal

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func Reg(host, name, id string, port int, tags []string) error {
	defaultConfig := api.DefaultConfig()
	h := AppConf.ConsulConfig.Host
	p := AppConf.ConsulConfig.Port
	defaultConfig.Address = fmt.Sprintf("%s:%d", h, p)
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		return err
	}
	agentServiceRegistration := new(api.AgentServiceRegistration)
	agentServiceRegistration.Address = defaultConfig.Address
	agentServiceRegistration.Port = port
	agentServiceRegistration.ID = id
	agentServiceRegistration.Name = name
	agentServiceRegistration.Tags = tags
	serverAddr := fmt.Sprintf("http://%s:%d/health", host, port)
	check := api.AgentServiceCheck{
		HTTP:                           serverAddr,
		Timeout:                        "3s",
		Interval:                       "1s",
		DeregisterCriticalServiceAfter: "5s",
	}
	agentServiceRegistration.Check = &check
	err = client.Agent().ServiceRegister(agentServiceRegistration)
	return err
}

func GetServiceList() error {
	defaultConfig := api.DefaultConfig()
	h := AppConf.ConsulConfig.Host
	p := AppConf.ConsulConfig.Port
	defaultConfig.Address = fmt.Sprintf("%s:%d", h, p)
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		return err
	}
	serviceList, err := client.Agent().Services()
	if err != nil {
		return err
	}
	for k, v := range serviceList {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("----------")
	}
	return nil
}

func FilterService() error {
	defaultConfig := api.DefaultConfig()
	h := AppConf.ConsulConfig.Host
	p := AppConf.ConsulConfig.Port
	defaultConfig.Address = fmt.Sprintf("%s:%d", h, p)
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		return err
	}
	serviceList, err := client.Agent().ServicesWithFilter("Service==account_web")
	if err != nil {
		return err
	}
	for k, v := range serviceList {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("----------")
	}
	return nil
}
