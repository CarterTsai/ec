package jet

// Config for jet template engine
type Config struct {
	Layout      string
	Vars        map[string]string
	Funcs       map[string]interface{}
	LayoutFuncs map[string]interface{}
}

// DefaultConfig returns the default configuration for the django template engine
func DefaultConfig() Config {
	return Config{
		Layout:      "",
		Vars:        make(map[string]string),
		Funcs:       make(map[string]interface{}, 0),
		LayoutFuncs: make(map[string]interface{}, 0),
	}
}
