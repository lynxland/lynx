//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"log/slog"

	"github.com/google/wire"
	"github.com/lynxland/lynx"
	"github.com/lynxland/lynx/boot"
)

func wireBootstrap(app lynx.App, slogger *slog.Logger) (*boot.Bootstrap, func(), error) {
	panic(wire.Build(ProviderSet))
}
