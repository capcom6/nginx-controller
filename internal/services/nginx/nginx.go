package nginx

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
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

func (n *Nginx) Apply(c Context) error {
	confPath := path.Join(n.Config.ConfigPath, c.DomainName+".conf")

	file, err := os.OpenFile(confPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening config file: %w", err)
	}

	if err := n.Render(c, file); err != nil {
		file.Close()
		return fmt.Errorf("error rendering config file: %w", err)
	}

	file.Close()

	return n.reload()
}

func (n *Nginx) Remove(hostname string) error {
	confPath := path.Join(n.Config.ConfigPath, hostname+".conf")

	err := os.Remove(confPath)
	if err != nil {
		return fmt.Errorf("error removing config file: %w", err)
	}

	return n.reload()
}

func (n *Nginx) reload() error {
	cmd := exec.Command("nginx", "-s", "reload")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error reloading nginx: %w", err)
	}
	return nil
}

// func (n *Nginx) validate() error {
// 	cmd := exec.Command("nginx", "-t")
// 	stdout, err := cmd.StdoutPipe()
// 	if err != nil {
// 		return fmt.Errorf("error getting stdout pipe: %w", err)
// 	}

// 	err = cmd.Start()

// 	return nil
// }
