package goaws

import "github.com/rustyeddy/store"

/*
This file exists primarily to handle package wide variables and
state.
*/

var (
	C Configuration
	S *store.Store
)

func init() {
	C = DefaultConfig
}

func SetStore(st *store.Store) {
	S = st
}

func SetConfig(cfg *Configuration) {
	C = *cfg
}
