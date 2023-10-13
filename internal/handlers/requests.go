package handlers

type PutHostname struct {
	Upstream []Upstream `json:"upstream" validate:"required,dive"`
}

type Upstream struct {
	Name   string `json:"name"`
	Host   string `json:"host" validate:"required"`
	Weight uint8  `json:"weight" validate:"min=0,max=100"`
}
