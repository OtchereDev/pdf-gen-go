package generator

type GenerationParam struct {
	TemplateName  string
	Data          map[string]interface{}
	RemoveMargins bool
	WithHeader    bool
}

type Generator interface {
	CompileTemplate(name string, data map[string]interface{}) (string, error)
	GeneratePDF(p GenerationParam) (string, error)
}
