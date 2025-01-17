package loadbalancer

import (
	"errors"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ukfast/sdk-go/pkg/client"
	"github.com/ukfast/sdk-go/pkg/connection"
	loadbalancerservice "github.com/ukfast/sdk-go/pkg/service/loadbalancer"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				DefaultFunc: func() (interface{}, error) {
					key := os.Getenv("UKF_API_KEY")
					if key != "" {
						return key, nil
					}

					return "", errors.New("api_key required")
				},
				Description: "API token required to authenticate with UKFast APIs. See https://developers.ukfast.io for more details",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"loadbalancer_accessip":    dataSourceAccessIP(),
			"loadbalancer_acl":         dataSourceACL(),
			"loadbalancer_bind":        dataSourceBind(),
			"loadbalancer_certificate": dataSourceCertificate(),
			"loadbalancer_cluster":     dataSourceCluster(),
			"loadbalancer_listener":    dataSourceListener(),
			"loadbalancer_target":      dataSourceTarget(),
			"loadbalancer_targetgroup": dataSourceTargetGroup(),
			"loadbalancer_vip":         dataSourceVip(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"loadbalancer_accessip":    resourceAccessIP(),
			"loadbalancer_acl":         resourceACL(),
			"loadbalancer_bind":        resourceBind(),
			"loadbalancer_certificate": resourceCertificate(),
			"loadbalancer_cluster":     resourceCluster(),
			"loadbalancer_listener":    resourceListener(),
			"loadbalancer_target":      resourceTarget(),
			"loadbalancer_targetgroup": resourceTargetGroup(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return getService(d.Get("api_key").(string)), nil
}

func getClient(apiKey string) client.Client {
	return client.NewClient(connection.NewAPIKeyCredentialsAPIConnection(apiKey))
}

func getService(apiKey string) loadbalancerservice.LoadBalancerService {
	return getClient(apiKey).LoadBalancerService()
}
