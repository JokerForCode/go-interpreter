package object

import "fmt"

const (
	INTEGER_OBJ      = "INTEGER"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
)

type Null struct{}

type Integer struct {
	Value int64
}

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}
type Boolean struct {
	Value bool
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }
