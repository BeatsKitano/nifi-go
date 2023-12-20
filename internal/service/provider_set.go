package service

import "github.com/google/wire"

// ProviderSet is biz providers.
var ServiceSet = wire.NewSet(
	NewPageQuery,
)
