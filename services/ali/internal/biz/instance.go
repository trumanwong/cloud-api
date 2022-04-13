package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Instance 定义返回的数据结构体
type Instance struct {
	// 阿里云请求id
	RequestId string
	// 实例ID
	InstanceId string
	// 实例名称
	InstanceName string
	// 实例规格
	InstanceType string
	// 实例的计费方式
	InstanceChargeType string
	// 实例所属地域ID
	RegionId string
	// vCPU数
	CPU uint32
	// 实例操作系统的英文名称
	OsNameEn string
	// 实例状态
	Status string
	// 实例创建时间，UTC+0时间
	CreationTime int
	// 实例网络类型
	InstanceNetworkType string
	// 内存大小，单位为MiB
	Memory uint32
	// 实例规格附带的GPU类型
	GpuSpec string
	// 按量付费实例的自动释放时间
	AutoReleaseTime string
	// 实例规格附带的GPU数量
	GpuAmount uint32
	// 实例的操作系统类型
	OsType string
	// 过期时间
	ExpiredTime string
	// 实例最近一次的启动时间
	StartTime string
	// 实例公网IP地址
	PublicIpAddress []string
	// 经典网络类型实例的内网IP地址
	InnerIpAddress []string
	// HPC实例的Rdma网络IP
	RdmaIpAddress []string
	// 实例所属安全组集合
	SecurityGroupIds []string
	EipAddress       EipAddress
	// 实例是否可以挂载数据盘
	DeviceAvailable bool
	// 实例挂载的本地存储容量
	LocalStorageCapacity uint64
	// 实例挂载的本地存储数量
	LocalStorageAmount uint32
	// 实例运行的镜像ID
	ImageId string
}

// EipAddress 弹性公网IP绑定信息
type EipAddress struct {
	// 弹性公网IP的ID
	AllocationId string
	// 是否可以解绑弹性公网IP
	IsSupportUnAssociate bool
	// 弹性公网IP的计费方式
	InternalChargeType string
	// 弹性公网IP
	IpAddress string
	// 弹性公网IP
	Bandwidth uint32
}

type CreateInstanceRequest struct {
	// 区域id
	RegionId string
	// 镜像id
	ImageId string
	// 实例名称
	Name string
	// 实例规格
	InstanceType string
	// 系统盘大小
	SystemDiskSize uint32
	// 是否为instance_name自动添加有序后缀（默认false）
	UniqueSuffix bool
	// 创建数量(1-100，默认1)
	Amount uint32
	// 实例密码
	Password string
}

type CreateInstanceResponse struct {
	// 阿里云请求id
	RequestId string
	// 订单id，该参数只有创建包年包月ECS实例（请求参数InstanceChargeType=PrePaid）时有返回值
	OrderId string
	// 订单成交价
	TradePrice float32
	// 实例id列表
	InstanceIdSets []string
}

// InstanceResponse is a Greater repo.
type InstanceResponse interface {
	Create(context.Context, *CreateInstanceRequest) (*CreateInstanceResponse, error)
	ListAll(context.Context) ([]*Instance, error)
}

// InstanceUseCase is a Instance UseCase.
type InstanceUseCase struct {
	repo InstanceResponse
	log  *log.Helper
}

// NewInstanceUseCase new a Instance UseCase.
func NewInstanceUseCase(repo InstanceResponse, logger log.Logger) *InstanceUseCase {
	return &InstanceUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreateInstance creates a Instance, and returns the new Instance.
func (uc *InstanceUseCase) CreateInstance(ctx context.Context, g *CreateInstanceRequest) (*CreateInstanceResponse, error) {
	uc.log.WithContext(ctx).Infof("CreateInstance: %v", g.Name)
	return uc.repo.Create(ctx, g)
}
