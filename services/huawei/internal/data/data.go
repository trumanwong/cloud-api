package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	ecsregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
	ims "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ims/v2"
	imsregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ims/v2/region"
	"huawei/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

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

func getAuth(accessKey, secretKey string) basic.Credentials {
	return basic.NewCredentialsBuilder().
		WithAk(accessKey).
		WithSk(secretKey).
		Build()
}

func getImsClient(accessKey, secretKey string, regionId string) *ims.ImsClient {
	return ims.NewImsClient(
		ims.ImsClientBuilder().
			WithRegion(imsregion.ValueOf(regionId)).
			WithCredential(getAuth(accessKey, secretKey)).
			Build())
}

func getEcsClient(accessKey, secretKey string, regionId string) *ecs.EcsClient {
	return ecs.NewEcsClient(
		ecs.EcsClientBuilder().
			WithRegion(ecsregion.ValueOf(regionId)).
			WithCredential(getAuth(accessKey, secretKey)).Build())
}
