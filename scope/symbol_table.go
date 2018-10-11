package scope

import (
	"fmt"
	"sync"
)

const CUSTOM_TYPE_BASE = 31

var scopeId int64 = 0
var customTypeId = DataType(CUSTOM_TYPE_BASE)
var customTypeIdMutex = &sync.Mutex{}
var scopeIdMutex = &sync.Mutex{}

var GlobalScope = &Scope{
	Id:      0,
	Parent:  nil,
	Symbols: SymbolTable{},
	Types:   CustomTypeTable{},
}

type CustomTypeTable map[string]DataType

type SymbolTable map[string]int64

type Scope struct {
	Id      int64
	Parent  *Scope
	Symbols SymbolTable
	Types   CustomTypeTable
}

func (s *Scope) AddSymbol(symbolName string, symbolType int64) error {
	if _, ok := s.Symbols[symbolName]; !ok {
		s.Symbols[symbolName] = symbolType
		return nil
	}
	return fmt.Errorf("symbol '%s' is already defined in this scope", symbolName)
}

func (s *Scope) AddType(typeName string) error {
	if _, ok := s.Types[typeName]; !ok {
		customTypeIdMutex.Lock()
		customTypeId++
		newTypeId := customTypeId
		customTypeIdMutex.Unlock()
		s.Types[typeName] = newTypeId
		return nil
	}

	return fmt.Errorf("type '%s' is already defined in this scope", typeName)
}

func NewScope(parent *Scope) *Scope {
	scopeIdMutex.Lock()
	scopeId++
	currentScopeId := scopeId
	scopeIdMutex.Unlock()

	return &Scope{
		Id:      currentScopeId,
		Parent:  parent,
		Symbols: SymbolTable{},
		Types:   CustomTypeTable{},
	}
}
