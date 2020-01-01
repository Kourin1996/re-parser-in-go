package vm

import (
	"fmt"

	"github.com/Kourin1996/re-parser-in-go/operator"
)

type VM struct {
	Program     operator.Program
	Input       string
	ContextList [](*Context)
}

func NewVM(program operator.Program, input string) *VM {
	baseContext := NewContext().
		SetStartCodeCursor(0).
		SetCurrentCodeCursor(0).
		SetEndCodeCursor(len(program)).
		SetStartInputCursor(0).
		SetCurrentInputCursor(0).
		SetEndInputCursor(len(input))
	return &VM{program, input, [](*Context){baseContext}}
}

func (vm *VM) handleOChar(op *operator.OChar, context *Context) {
	d := struct {
		Input   string
		Ch      string
		Context Context
	}{vm.Input, string(op.Ch), *context}
	fmt.Printf("[VM-OChar] Read Character %+v\n", d)

	if context.CurrentInputCursor >= context.EndInputCursor {
		context.SetStatus(FAILED)
		fmt.Printf("[VM-OChar] Out of Index\n")
		return
	}
	if ch := vm.Input[context.CurrentInputCursor]; op.Ch != ch {
		context.SetStatus(FAILED)
		fmt.Printf("[VM-OChar] Mismatch expect=%c, got=%c\n", op.Ch, ch)
		return
	}
	context.SetCurrentInputCursor(context.CurrentInputCursor + 1)
	context.SetCurrentCodeCursor(context.CurrentCodeCursor + 1)
}

func (vm *VM) handleOMatch(op *operator.OMatch, context *Context) {
	context.SetStatus(SUCCEEDED)
	fmt.Printf("[VM-OMatch] %+v\n", context)
}

func (vm *VM) handleOJump(op *operator.OJmp, context *Context) {
	fmt.Printf("[VM-OJump] Jump to %d\n", op.Next)
	context.SetCurrentCodeCursor(op.Next)
}

func (vm *VM) handleOSplit(op *operator.OSplit, context *Context) {
	left := context.Clone().SetCurrentCodeCursor(op.Left).SetStatus(FORKED)
	right := context.Clone().SetCurrentCodeCursor(op.Right).SetStatus(FORKED)
	fmt.Printf("[VM-OSplit] Fork Context %+v => left:%+v, right:%+v\n", context, left, right)

	context.SetStatus(KILLED)
	vm.ContextList = append(vm.ContextList, left)
	vm.ContextList = append(vm.ContextList, right)
}

func (vm *VM) handleOSaveL(op *operator.OSaveL, context *Context) {
	if _, ok := context.IndexMap[op.Label]; ok {
		panic("Conflict in IndexMap")
	}
	context.IndexMap[op.Label] = Index{context.CurrentInputCursor, -1}
	context.SetCurrentCodeCursor(context.CurrentCodeCursor + 1)
}

func (vm *VM) handleOSaveR(op *operator.OSaveR, context *Context) {
	index, ok := context.IndexMap[op.Label]
	if !ok {
		panic("No index")
	}
	index.End = context.CurrentInputCursor
	context.IndexMap[op.Label] = index
	context.SetCurrentCodeCursor(context.CurrentCodeCursor + 1)
}

func (vm *VM) ExecuteCode(context *Context) {
	code := vm.Program[context.CurrentCodeCursor]
	fmt.Printf("[VM] ExecuteCode %+v %+v\n", context, code)
	switch op := code.(type) {
	case *operator.OChar:
		vm.handleOChar(op, context)
		break
	case *operator.OMatch:
		vm.handleOMatch(op, context)
		break
	case *operator.OJmp:
		vm.handleOJump(op, context)
		break
	case *operator.OSplit:
		vm.handleOSplit(op, context)
		break
	case *operator.OSaveL:
		vm.handleOSaveL(op, context)
		break
	case *operator.OSaveR:
		vm.handleOSaveR(op, context)
		break
	default:
		break
	}
}

func (vm *VM) IsAllContextClosed() bool {
	if len(vm.ContextList) == 0 {
		return true
	}
	for _, cxt := range vm.ContextList {
		if cxt.Status != SUCCEEDED && cxt.Status != FAILED && cxt.Status != KILLED {
			return false
		}
	}
	return true
}

func (vm *VM) GetFirstContextToDo() (*Context, int) {
	for ix, cxt := range vm.ContextList {
		if cxt.Status == INITIALIZED || cxt.Status == FORKED {
			return cxt, ix
		}
	}
	return nil, -1
}

func (vm *VM) RunCode() {
	fmt.Printf("[VM] RunCode%+v\n", vm.ContextList)
	for !vm.IsAllContextClosed() {
		context, _ := vm.GetFirstContextToDo()
		context.SetStatus(RUNNING)
		for context.Status != SUCCEEDED && context.Status != FAILED && context.Status != KILLED {
			vm.ExecuteCode(context)

			fmt.Printf("Context List\n")
			for _, cxt := range vm.ContextList {
				fmt.Printf("%+v\n", cxt)
			}
		}
		fmt.Printf("[VM] Done Thread %+v\n", context)
	}
	fmt.Printf("Result\n")
	for _, cxt := range vm.ContextList {
		fmt.Printf("%+v\n", cxt)
	}
}
