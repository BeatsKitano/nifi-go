package server

import "github.com/google/wire"

// ServerSet is biz providers.

var ServerSet = wire.NewSet(NewHTTPServer)
