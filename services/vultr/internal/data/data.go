package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/vultr/govultr/v2"
	"golang.org/x/oauth2"
	"vultr/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewInstanceRepo, NewImageRepo, NewRegionRepo, NewPlanRepo)

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

func getClient(ctx context.Context, accessToken string) *govultr.Client {
	config := &oauth2.Config{}
	ts := config.TokenSource(ctx, &oauth2.Token{AccessToken: accessToken})
	return govultr.NewClient(oauth2.NewClient(ctx, ts))
}
