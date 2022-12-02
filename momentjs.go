package momentjs

import (
	_ "embed"
	"fmt"
	"sync"

	"github.com/lithdew/quickjs"
)

//go:embed asset/moment.min.js
var momentjsCode string

type MomentJS struct {
	QjsRuntime *quickjs.Runtime
	QjsCtx     *quickjs.Context
	Global     *quickjs.Value
	free       func()
}

func New() (*MomentJS, error) {
	qjsRuntime := quickjs.NewRuntime()
	qjsCtx := qjsRuntime.NewContext()
	global := qjsCtx.Globals()
	free := func() {
		qjsCtx.Free()
		qjsRuntime.Free()
	}
	momentjsResult, err := qjsCtx.Eval(momentjsCode)
	if err != nil {
		free()
		return nil, err
	}
	momentjsResult.Free()
	var freeOnce sync.Once
	return &MomentJS{
		QjsRuntime: &qjsRuntime,
		QjsCtx:     qjsCtx,
		Global:     &global,
		free: func() {
			freeOnce.Do(free)
		},
	}, nil
}

func (m *MomentJS) Free() {
	m.free()
}

// this method must be called in locked method.
func (m *MomentJS) ClearGlobal(name string) error {
	result, err := m.QjsCtx.Eval(fmt.Sprintf("delete globalThis.%s", name))
	if err != nil {
		return err
	}
	result.Free()
	return nil
}
