package api

import "github.com/google/wire"

// ProviderSet is biz providers.
var ApiSet = wire.NewSet(
	NewPageController,
)
