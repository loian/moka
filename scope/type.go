package scope

type DataType int32

const (
	NULL = iota
	BOOL
	INT
	FLOAT
	DECIMAL
	STRING
	FUNCTION
)