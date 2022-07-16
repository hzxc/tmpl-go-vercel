package mysql

import (
	"context"
	"tmpl-go-vercel/app/global"

	"github.com/planetscale/planetscale-go/planetscale"
)

type PlanetscaleMysql struct {
	token string
}

func NewPlanetscaleMysql() *PlanetscaleMysql {
	return &PlanetscaleMysql{
		token: global.PlanetscaleToken,
	}
}

func (p *PlanetscaleMysql) Mysql(ctx context.Context) (*planetscale.Client, error) {
	return planetscale.NewClient(
		planetscale.WithAccessToken(p.token),
	)
}
