package service

import (
	v1 "ali/api/instance/v1"
	"context"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"os"
	"testing"
)

func TestInstanceService_ListInstanceTypes(t *testing.T) {
	accessKeyId := tea.String(os.Getenv("accessKeyId"))
	accessKeySecret := tea.String(os.Getenv("accessKeySecret"))
	endpoint := tea.String(os.Getenv("endpoint"))
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	_, err = instanceClient.ListInstanceTypes(context.Background(), &v1.ListInstanceTypesRequest{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Endpoint:        endpoint,
	})
	assert.Equal(t, err, nil)
}
