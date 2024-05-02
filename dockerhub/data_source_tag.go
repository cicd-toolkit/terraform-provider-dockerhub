package dockerhub

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDockerHubLastTag() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDockerHubLastTagRead,
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_tag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceDockerHubLastTagRead(d *schema.ResourceData, meta interface{}) error {
	cli := meta.(*client.Client)
	repository := d.Get("repository").(string)

	ctx := context.Background()

	tags, err := cli.ImageList(ctx, repository, types.ImageListOptions{All: false})
	if err != nil {
		return err
	}

	latestTag := ""
	for _, tag := range tags {
		for _, t := range tag.RepoTags {
			parts := strings.Split(t, ":")
			tag := parts[len(parts)-1]
			if latestTag == "" || tag > latestTag {
				latestTag = tag
			}
		}
	}

	d.Set("last_tag", latestTag)
	d.SetId(repository)

	return nil
}
