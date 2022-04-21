package data

import (
	"ali/internal/conf"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewInstanceRepo, NewImageRepo, NewRegionRepo, NewInstanceTypeRepo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

func getClient(accessKeyId string, accessKeySecret, endpoint string) (result *ecs20140526.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: tea.String(accessKeyId),
		// 您的AccessKey Secret
		AccessKeySecret: tea.String(accessKeySecret),
	}
	if endpoint == "" {
		endpoint = "ecs.cn-shenzhen.aliyuncs.com"
	}
	// 访问的域名
	config.Endpoint = tea.String(endpoint)
	result = &ecs20140526.Client{}
	result, _err = ecs20140526.NewClient(config)
	return result, _err
}
