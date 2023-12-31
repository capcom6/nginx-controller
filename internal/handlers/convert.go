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

		port := uint16(80)
		if upstream.Port != 0 {
			port = upstream.Port
		}

		upstreams[i] = nginx.UpstreamServer{
			Host:    upstream.Host,
			Port:    port,
			Options: options,
		}
	}

	return nginx.Context{
		DomainName:   hostname,
		UpstreamName: hostname,
		Servers:      upstreams,
	}
}
