package modules

import (
	"github.com/kujtimiihoxha/plis/api"
	"github.com/kujtimiihoxha/plis/logger"
	"github.com/robertkrimen/otto"
)

type PlisModule struct {
	plisAPI *api.PlisAPI
	flags   *otto.Object
	args    *otto.Object
}

func (p *PlisModule) ModuleLoader(vm *otto.Otto) *otto.Object {
	obj, _ := vm.Call("new Object", nil)
	v := obj.Object()
	v.Set("help", p.help)
	v.Set("runPlisTool", p.runPlisTool)
	v.Set("flags", p.flags)
	v.Set("args", p.args)
	return v
}
func (p *PlisModule) help(call otto.FunctionCall) otto.Value {
	p.plisAPI.Help()
	return otto.Value{}
}
func (p *PlisModule) forceOverride(call otto.FunctionCall) otto.Value {
	b, _ := call.Argument(0).ToBoolean()
	p.plisAPI.ForceOverride(b)
	return otto.TrueValue()
}
func (p *PlisModule) runPlisTool(call otto.FunctionCall) otto.Value {
	pCmd := call.Argument(0).String()
	args := call.Argument(1)
	if !args.IsObject() {
		logger.GetLogger().Error("The arguments must be an array")
		return otto.FalseValue()
	}
	s := []string{}
	for _, a := range args.Object().Keys() {
		vl, _ := args.Object().Get(a)
		s = append(s, vl.String())
	}
	if err := p.plisAPI.RunPlisTool(pCmd, s); err != nil {
		logger.GetLogger().Errorf("Error while exectuing plis command: %s", err.Error())
		return otto.FalseValue()
	}
	return otto.TrueValue()
}

func NewPlisModule(flags *otto.Object, args *otto.Object, api *api.PlisAPI) *PlisModule {
	return &PlisModule{
		plisAPI: api,
		flags:   flags,
		args:    args,
	}
}
