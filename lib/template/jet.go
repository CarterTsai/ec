package jet

import (
	"io"
	"sync"

	"fmt"
)

type (
	// Engine the jet engine
	Engine struct {
		Config        Config
		templateCache map[string]interface{}
		mu            sync.Mutex
	}
)

// New creates and returns a Pongo template engine
func New(cfg ...Config) *Engine {
	fmt.Println("[New Engine For Jet]")

	c := DefaultConfig()
	if len(cfg) > 0 {
		c = cfg[0]
	}

	return &Engine{Config: c, templateCache: make(map[string]interface{})}
}

// LoadDirectory builds the markdown templates
func (j *Engine) LoadDirectory(dir string, extension string) error {
	fmt.Println("[LoadDirectory For Jet]", dir, extension)
	return nil
}

// LoadAssets loads the templates by binary
func (j *Engine) LoadAssets(virtuaDir string, virtualExt string, assetFn func(name string) ([]byte, error), namesFn func() []string) error {
	fmt.Println("[LoadAssets For Jet]")
	fmt.Println(virtuaDir, virtualExt)
	return nil
}

func (j *Engine) fromCache(relativeName string) []byte {
	fmt.Println("[fromCache For Jet]", relativeName)
	j.mu.Lock()

	ok := true

	if ok {
		j.mu.Unlock() // defer is slow
		return nil
	}
	j.mu.Unlock() // defer is slow
	return nil
}

// ExecuteWriter executes a templates and write its results to the out writer
// layout here is useless
func (j *Engine) ExecuteWriter(out io.Writer, name string, binding interface{}, options ...map[string]interface{}) error {

	fmt.Println("[ExecuteWriter]")
	fmt.Println(out, name, binding, options)

	if tmpl := j.fromCache(name); tmpl != nil {
		_, err := out.Write(tmpl)
		return err
	}

	return fmt.Errorf("[IRIS TEMPLATES] Template with name %s doesn't exists in the dir", name)
}
