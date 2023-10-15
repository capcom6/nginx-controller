package nginx

type UpstreamServer struct {
	Host    string
	Port    uint16
	Options string
}

type Context struct {
	DomainName   string
	UpstreamName string
	Servers      []UpstreamServer
}
