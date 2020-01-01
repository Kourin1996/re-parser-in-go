package operator

import "fmt"

const (
	CHAR  = "CHAR"
	MATCH = "MATCH"
	JMP   = "JMP"
	SPLIT = "SPLIT"
	SAVEL = "SAVEL"
	SAVER = "SAVER"
)

type Program []Operator

type Operator interface {
	GetOperator() string
	String() string
}

type OChar struct {
	Ch byte
}

func (op *OChar) GetOperator() string {
	return CHAR
}

func (op *OChar) String() string {
	return fmt.Sprintf("%s %c", op.GetOperator(), op.Ch)
}

type OMatch struct{}

func (op *OMatch) GetOperator() string {
	return MATCH
}

func (op *OMatch) String() string {
	return fmt.Sprintf("%s", op.GetOperator())
}

type OJmp struct {
	Next int
}

func (op *OJmp) GetOperator() string {
	return JMP
}

func (op *OJmp) String() string {
	return fmt.Sprintf("%s %s", op.GetOperator(), op.Next)
}

type OSplit struct {
	Left  int
	Right int
}

func (op *OSplit) GetOperator() string {
	return SPLIT
}

func (op *OSplit) String() string {
	return fmt.Sprintf("%s %s,%s", op.GetOperator(), op.Left, op.Right)
}

type OSaveL struct {
	Label string
}

func (op *OSaveL) GetOperator() string {
	return SAVEL
}

func (op *OSaveL) String() string {
	return fmt.Sprintf("%s %s", op.GetOperator(), op.Label)
}

type OSaveR struct {
	Label string
}

func (op *OSaveR) GetOperator() string {
	return SAVER
}

func (op *OSaveR) String() string {
	return fmt.Sprintf("%s %s", op.GetOperator(), op.Label)
}
