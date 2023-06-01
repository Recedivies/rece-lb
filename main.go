package main

import (
	"fmt"
	"net/http"

	"github.com/Recedivies/rece-lb/balancer"
	"github.com/spf13/viper"
)

type Config struct {
	Port      int      `json:"port"`
	Servers   []string `json:"servers"`
	Algorithm string   `json:"algorithm"`
}

func main() {
	config := Config{}
	viper.AddConfigPath(".")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	_ = viper.ReadInConfig()
	_ = viper.Unmarshal(&config)

	fmt.Printf("LoadBalancer serving requests at port %d\n", config.Port)

	// send the parsed config to create a new load balancer
	http.ListenAndServe(fmt.Sprintf(":%v", config.Port), balancer.NewLoadBalancer(
		config.Servers,
		config.Algorithm,
	))
}
