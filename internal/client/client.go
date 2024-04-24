package client

import (
	"context"
	"encoding/json"
	"fmt"

	ldapi "github.com/launchdarkly/api-client-go/v14"
)

type Client interface {
	ListFlags(ctx context.Context, projKey string) ([]byte, error)
	GetFlag(ctx context.Context, projKey string, envKey string, key string) ([]byte, error)
}

type LDClient struct {
	client *ldapi.APIClient
}

type ClientFn func(accessToken string, baseURI string, version string) Client

func LDClientFn() ClientFn {
	return func(accessToken string, baseURI string, version string) Client {
		return NewLDClient(accessToken, baseURI, version)
	}
}

func NewLDClient(accessToken string, baseURI string, version string) LDClient {
	config := ldapi.NewConfiguration()
	config.AddDefaultHeader("Authorization", accessToken)
	config.UserAgent = fmt.Sprintf("launchdarkly-cli/v%s", version)
	config.Servers[0].URL = baseURI
	client := ldapi.NewAPIClient(config)

	return LDClient{
		client: client,
	}
}

func (c LDClient) ListFlags(ctx context.Context, projKey string) ([]byte, error) {
	flags, _, err := c.client.
		FeatureFlagsApi.
		GetFeatureFlags(ctx, projKey).
		Execute()
	if err != nil {
		return nil, err
	}

	responseJSON, err := json.Marshal(flags)
	if err != nil {
		return nil, err
	}

	return responseJSON, nil
}

func (c LDClient) GetFlag(ctx context.Context, projKey string, envKey string, key string) ([]byte, error) {
	flag, _, err := c.client.
		FeatureFlagsApi.
		GetFeatureFlag(ctx, projKey, key).
		Env(envKey).
		Execute()
	if err != nil {
		return nil, err
	}

	responseJSON, err := json.Marshal(flag)
	if err != nil {
		return nil, err
	}

	return responseJSON, nil
}
