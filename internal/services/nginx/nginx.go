package nginx

import (
	"io"
	"sync"
	"text/template"
)

type Config struct {
	ConfigPath     string
	ConfigTemplate string
}

type Nginx struct {
	Config Config

	template *template.Template
	once     sync.Once
}

func New(cfg Config) *Nginx {
	return &Nginx{
		Config: cfg,
	}
}

func (n *Nginx) init() (err error) {
	n.once.Do(func() {
		n.template, err = template.New("config").Parse(n.Config.ConfigTemplate)
	})
	return
}

func (n *Nginx) Render(c Context, w io.Writer) error {
	err := n.init()
	if err != nil {
		return err
	}

	return n.template.Execute(w, c)
}
