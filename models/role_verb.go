package models

import (
	"fmt"
	"strings"
)

const (
	RoleVerbUnknown RoleVerb = iota
	RoleVerbCreate
	RoleVerbUpdate
	RoleVerbDelete
	RoleVerbGet
	RoleVerbList
	RoleVerbAll
)

const (
	roleVerbNameUnknown = "UNKNOWN"
	roleVerbNameCreate  = "CREATE"
	roleVerbNameUpdate  = "UPDATE"
	roleVerbNameDelete  = "DELETE"
	roleVerbNameGet     = "GET"
	roleVerbNameList    = "LIST"
	roleVerbAll         = "ALL"
)

var resolveVerbToString = map[RoleVerb]string{
	RoleVerbUnknown: roleVerbNameUnknown,
	RoleVerbCreate:  roleVerbNameCreate,
	RoleVerbUpdate:  roleVerbNameUpdate,
	RoleVerbDelete:  roleVerbNameDelete,
	RoleVerbGet:     roleVerbNameGet,
	RoleVerbList:    roleVerbNameList,
	RoleVerbAll:     roleVerbAll,
}

var resolveStringToVerb = map[string]RoleVerb{
	roleVerbNameUnknown: RoleVerbUnknown,
	roleVerbNameCreate:  RoleVerbCreate,
	roleVerbNameUpdate:  RoleVerbUpdate,
	roleVerbNameDelete:  RoleVerbDelete,
	roleVerbNameGet:     RoleVerbGet,
	roleVerbNameList:    RoleVerbList,
	roleVerbAll:         RoleVerbAll,
}

type RoleVerb uint

func (v RoleVerb) String() string {
	if s, exists := resolveVerbToString[v]; exists {
		return s
	}

	return resolveVerbToString[RoleVerbUnknown]
}

func NewRoleVerbByString(s string) (RoleVerb, error) {
	s = strings.ToUpper(s)

	if v, exists := resolveStringToVerb[s]; exists {
		return v, nil
	}

	return RoleVerbUnknown, fmt.Errorf("unknown verb")
}
