package template

type UpstreamServer struct {
	Host    string
	Options string
}

type Context struct {
	DomainName   string
	UpstreamName string
	Servers      []UpstreamServer
}
