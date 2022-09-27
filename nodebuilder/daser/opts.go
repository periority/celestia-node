package daser

import (
	"go.uber.org/fx"

	"github.com/celestiaorg/celestia-node/das"
)

func WithMetrics(enable bool) fx.Option {
	if !enable {
		return fx.Options()
	}
	return fx.Invoke(func(d *das.DASer) error {
		return d.InitMetrics()
	})
}