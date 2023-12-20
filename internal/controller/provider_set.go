package controller

import "github.com/google/wire"

// ProviderSet is biz providers.
var ControllerSet = wire.NewSet(
	NewPageController,
)
