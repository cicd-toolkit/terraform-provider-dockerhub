package dockerhub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DOCKER_USERNAME", nil),
				Description: "Username for authentication.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DOCKER_PASSWORD", nil),
				Description: "Password for authentication.",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"tag": dataSourceDockerHubLastTag(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	ctx = context.Background()
	cli, err := client.NewClientWithOpts(client.WithUserAgent("Terraform-DockerHub-Provider"))
	if err != nil {
		return nil, err
	}
	authConfig := types.AuthConfig{
		Username: username,
		Password: password,
	}
	_, err = cli.RegistryLogin(ctx, authConfig)
	if err != nil {
		return nil, err
	}

	return cli, nil
}
