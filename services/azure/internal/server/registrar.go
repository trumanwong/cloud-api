package server

import (
	"azure/internal/conf"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	consulApi "github.com/hashicorp/consul/api"
)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulApi.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme

	client, err := consulApi.NewClient(c)
	if err != nil {
		panic(err)
	}
	return consul.New(client, consul.WithHealthCheck(false))
}
