package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Instance 定义返回的数据结构体
type Instance struct {
	// 阿里云请求id
	RequestId string
	// 订单id，该参数只有创建包年包月ECS实例（请求参数InstanceChargeType=PrePaid）时有返回值
	OrderId string
	// 订单成交价
	TradePrice string
	// 实例id列表
	InstanceIdSets []string
}

// InstanceRepo is a Greater repo.
type InstanceRepo interface {
	Save(context.Context, *Instance) (*Instance, error)
	Update(context.Context, *Instance) (*Instance, error)
	FindByID(context.Context, int64) (*Instance, error)
	ListByHello(context.Context, string) ([]*Instance, error)
	ListAll(context.Context) ([]*Instance, error)
}

// InstanceUsecase is a Instance usecase.
type InstanceUsecase struct {
	repo InstanceRepo
	log  *log.Helper
}

// NewInstanceUsecase new a Instance usecase.
func NewInstanceUsecase(repo InstanceRepo, logger log.Logger) *InstanceUsecase {
	return &InstanceUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateInstance creates a Instance, and returns the new Instance.
func (uc *InstanceUsecase) CreateInstance(ctx context.Context, g *Instance) (*Instance, error) {
	uc.log.WithContext(ctx).Infof("CreateInstance: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
