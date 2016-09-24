package jet

import (
	"io"
	"os"
	"path/filepath"
	"sync"

	"fmt"

	"github.com/CloudyKit/jet"
)

type (
	// Engine the jet engine
	Engine struct {
		Config        Config
		view          *jet.Set
		templateCache map[string]*jet.Template
		mu            sync.Mutex
	}
)

// New creates and returns a Pongo template engine
func New(cfg ...Config) *Engine {
	fmt.Printf("[New Engine For Jet]")

	c := DefaultConfig()
	if len(cfg) > 0 {
		c = cfg[0]
	}

	return &Engine{Config: c, templateCache: make(map[string]*jet.Template)}
}

// LoadDirectory builds the jet templates
func (j *Engine) LoadDirectory(dir string, extension string) error {
	fmt.Println("[LoadDirectory For Jet]")
	var templateErr error
	// Walk the supplied directory and compile any files that match our extension list.
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if info == nil || info.IsDir() {
		} else {

			// returns a relative path
			relPath, err := filepath.Rel(dir, path)
			if err != nil {
				templateErr = err
				return err
			}

			if j.view == nil {
				j.view = jet.NewHTMLSet(dir)
				j.view.SetDevelopmentMode(true)
			}
			// get file extension ex. '.html'
			ext := filepath.Ext(relPath)

			if ext == extension {
				templateView, err := j.view.GetTemplate(relPath)
				if err != nil {
					fmt.Println(err)
					templateErr = err
					return err
				}
				name := filepath.ToSlash(relPath)
				j.templateCache[name] = templateView
			}
		}
		return nil
	})

	return templateErr
}

// LoadAssets loads the templates by binary
func (j *Engine) LoadAssets(virtuaDir string, virtualExt string, assetFn func(name string) ([]byte, error), namesFn func() []string) error {
	fmt.Printf("[LoadAssets For Jet]")
	return nil
}

func (j *Engine) fromCache(relativeName string) *jet.Template {
	fmt.Println("[fromCache For Jet]", relativeName)
	j.mu.Lock()

	tmpl, ok := j.templateCache[relativeName]

	if ok {
		j.mu.Unlock() // defer is slow
		return tmpl
	}
	j.mu.Unlock() // defer is slow
	return nil
}

// ExecuteWriter executes a templates and write its results to the out writer
// layout here is useless
func (j *Engine) ExecuteWriter(out io.Writer, name string, binding interface{}, options ...map[string]interface{}) error {
	fmt.Println("[ExecuteWriter]")
	if tmpl := j.fromCache(name); tmpl != nil {

		err := tmpl.Execute(out, binding.(jet.VarMap), nil)

		fmt.Println(err)
		return err
	}

	return fmt.Errorf("[JET TEMPLATES] Template with name %s doesn't exists in the dir", name)
}
