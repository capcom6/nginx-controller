package handlers

import (
	"fmt"

	"github.com/capcom6/nginx-controller/internal/services/nginx"
)

func putHostnameToContext(hostname string, i PutHostname) nginx.Context {
	upstreams := make([]nginx.UpstreamServer, len(i.Upstream))
	for i, upstream := range i.Upstream {
		options := ""
		if upstream.Weight < 0 {
			options = "down"
		} else if upstream.Weight > 0 {
			options = fmt.Sprintf("weight=%d", upstream.Weight)
		}

		upstreams[i] = nginx.UpstreamServer{
			Host:    upstream.Host,
			Options: options,
		}
	}

	return nginx.Context{
		DomainName:   hostname,
		UpstreamName: "test",
		Servers:      upstreams,
	}
}
