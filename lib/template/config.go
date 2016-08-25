package jet

// Config for jet template engine
type Config struct {
	Layout      string
	Funcs       map[string]interface{}
	LayoutFuncs map[string]interface{}
}

// DefaultConfig returns the default configuration for the django template engine
func DefaultConfig() Config {
	return Config{
		Layout:      "",
		Funcs:       make(map[string]interface{}, 0),
		LayoutFuncs: make(map[string]interface{}, 0),
	}
}
