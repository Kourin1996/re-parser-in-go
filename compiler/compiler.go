package compiler

import (
	"fmt"

	"github.com/Kourin1996/re-parser-in-go/operator"
	"github.com/Kourin1996/re-parser-in-go/re"
)

func Compile(re re.RegularExpression) operator.Program {
	compiler := &Compiler{operator.Program{}}
	re.Accept(compiler)
	compiler.Codes = append(compiler.Codes, &operator.OMatch{})
	return compiler.Codes
}

type Compiler struct {
	Codes operator.Program
}

func (c *Compiler) VisitChar(e *re.REChar) {
	fmt.Printf("[Compiler] Visit Char %#v\n", e)
	c.Codes = append(c.Codes, &operator.OChar{e.Ch})
}

func (c *Compiler) VisitConcat(e *re.REConcat) {
	fmt.Printf("[Compiler] Visit Concat %#v\n", e)
	e.First.Accept(c)
	e.Second.Accept(c)
}

func (c *Compiler) VisitChoice(e *re.REChoice) {
	fmt.Printf("[Compiler] Visit Choice %#v\n", e)

	splitOp := &operator.OSplit{}
	c.Codes = append(c.Codes, splitOp)

	firstStart := len(c.Codes)
	e.First.Accept(c)
	secondStart := len(c.Codes)
	e.Second.Accept(c)

	splitOp.Left = firstStart
	splitOp.Right = secondStart
}

func (c *Compiler) VisitOption(e *re.REOption) {
	fmt.Printf("[Compiler] Visit Option %#v\n", e)
	splitOp := &operator.OSplit{}
	c.Codes = append(c.Codes, splitOp)

	exStart := len(c.Codes)
	e.Ex.Accept(c)
	nextStart := len(c.Codes)

	splitOp.Left = exStart
	splitOp.Right = nextStart
}

func (c *Compiler) VisitZeroOrMany(e *re.REZeroOrMany) {
	fmt.Printf("[Compiler] Visit Zero Or Many %#v\n", e)

	start := len(c.Codes)
	splitOp := &operator.OSplit{}
	c.Codes = append(c.Codes, splitOp)

	e.Ex.Accept(c)
	c.Codes = append(c.Codes, &operator.OSplit{start + 1, len(c.Codes) + 1})

	splitOp.Left = len(c.Codes)
	splitOp.Right = start + 1
}

func (c *Compiler) VisitOneOrMany(e *re.REOneOrMany) {
	fmt.Printf("[Compiler] Visit One Or Many %#v\n", e)

	start := len(c.Codes)
	e.Ex.Accept(c)
	end := len(c.Codes)
	c.Codes = append(c.Codes, &operator.OSplit{end, start})
}

func (c *Compiler) VisitGroup(e *re.REGroup) {
	fmt.Printf("[Compiler] Visit Group %#v\n", e)
	c.Codes = append(c.Codes, &operator.OSaveL{e.GroupName})
	e.Ex.Accept(c)
	c.Codes = append(c.Codes, &operator.OSaveR{e.GroupName})
}
