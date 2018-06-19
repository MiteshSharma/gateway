package api

type LocalRegistry struct {
}

func NewLocalRegistry() LocalRegistry {
	return LocalRegistry{}
}

func (reg LocalRegistry) GetService(service string) (string, error) {
	return "http://localhost:9000", nil
}
