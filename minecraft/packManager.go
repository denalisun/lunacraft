package minecraft

type MCFunction struct {
	Content []string
	Name    string
}

type PackManager struct {
	Functions       []*MCFunction
	FocusedFunction *MCFunction
}

func NewPackManager() *PackManager {
	mgr := &PackManager{}
	mgr.Functions = []*MCFunction{}
	mgr.FocusedFunction = nil
	return mgr
}

func (p *PackManager) AddNewFunction(name string) *MCFunction {
	fnc := &MCFunction{Name: name}
	p.Functions = append(p.Functions, fnc)
	return fnc
}

var GlobalPackManager *PackManager = NewPackManager()
