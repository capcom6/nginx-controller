package handlers

type PutHostname struct {
	Upstream []Upstream `json:"upstream" validate:"required,dive"`
}

type Upstream struct {
	Name   string `json:"name"`
	Host   string `json:"host" validate:"hostname,required"`
	Port   uint16 `json:"port" validate:"min=1,max=65535"`
	Weight int8   `json:"weight" validate:"min=0,max=100"`
}
