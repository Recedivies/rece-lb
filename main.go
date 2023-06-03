package main

import (
	"fmt"
	"net/http"

	"github.com/Recedivies/rece-lb/balancer"
	"github.com/spf13/viper"
)

type Config struct {
	Port                int      `mapstructure:"port"`
	Servers             []string `mapstructure:"servers"`
	Algorithm           string   `mapstructure:"algorithm"`
	HealthCheckType     string   `mapstructure:"health_check_type"`
	HealthCheckInterval int      `mapstructure:"health_check_interval"`
}

func main() {
	config := Config{}
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	_ = viper.ReadInConfig()
	_ = viper.Unmarshal(&config)

	fmt.Printf("LoadBalancer serving requests at port %d\n", config.Port)

	// send the parsed config to create a new load balancer
	http.ListenAndServe(fmt.Sprintf(":%v", config.Port), balancer.NewLoadBalancer(
		config.Servers,
		config.Algorithm,
		config.HealthCheckType,
		config.HealthCheckInterval,
	))
}
