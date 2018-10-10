package scope

type DataType int32

const (
	NULL = iota
	ANY
	BOOL
	INT
	FLOAT
	DECIMAL
	STRING
	FUNCTION
)
