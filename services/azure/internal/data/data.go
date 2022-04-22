package data

import (
	"azure/internal/conf"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/subscriptions"
	"github.com/Azure/azure-sdk-for-go/profiles/preview/marketplaceordering/mgmt/marketplaceordering"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
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

func getEnvironment(name string) (*azure.Environment, error) {
	env, err := azure.EnvironmentFromName(name)
	if err != nil {
		return nil, err
	}
	return &env, nil
}

func getAuthorizerForResource(tenantID, clientID, clientSecret string) (*autorest.BearerAuthorizer, error) {
	env, err := getEnvironment("AzurePublicCloud")
	if err != nil {
		return nil, err
	}
	oauthConfig, err := adal.NewOAuthConfig(env.ActiveDirectoryEndpoint, tenantID)
	if err != nil {
		return nil, err
	}
	token, err := adal.NewServicePrincipalToken(
		*oauthConfig, clientID, clientSecret, env.ResourceManagerEndpoint,
	)
	if err != nil {
		return nil, err
	}
	authorizer := autorest.NewBearerAuthorizer(token)
	return authorizer, nil
}

func getVMClient(tenantID, clientID, clientSecret, subscriptionId string) (*compute.VirtualMachinesClient, error) {
	vmClient := compute.NewVirtualMachinesClient(subscriptionId)
	authorizer, err := getAuthorizerForResource(tenantID, clientID, clientSecret)
	if err != nil {
		return nil, err
	}
	vmClient.Authorizer = authorizer
	return &vmClient, nil
}

func getImagesClient(tenantID, clientID, clientSecret, subscriptionId string) (*compute.ImagesClient, error) {
	imageClient := compute.NewImagesClient(subscriptionId)
	authorizer, err := getAuthorizerForResource(tenantID, clientID, clientSecret)
	if err != nil {
		return nil, err
	}
	imageClient.Authorizer = authorizer
	return &imageClient, nil
}

func getMarketplaceAgreementsClient(tenantID, clientID, clientSecret, subscriptionId string) (*marketplaceordering.MarketplaceAgreementsClient, error) {
	marketplaceAgreementsClient := marketplaceordering.NewMarketplaceAgreementsClient(subscriptionId)
	authorizer, err := getAuthorizerForResource(tenantID, clientID, clientSecret)
	if err != nil {
		return nil, err
	}
	marketplaceAgreementsClient.Authorizer = authorizer
	return &marketplaceAgreementsClient, nil
}

func getSubscriptionsClient(tenantID, clientID, clientSecret, subscriptionId string) (*subscriptions.Client, error) {
	client := subscriptions.NewClient()
	authorizer, err := getAuthorizerForResource(tenantID, clientID, clientSecret)
	if err != nil {
		return nil, err
	}
	client.Authorizer = authorizer
	return &client, nil
}
