package rds

import (
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/hacker65536/rds/internal/aws"
)

var (
	rdssvc *rds.Client
)

func init() {
	cfgCp := cfg.Copy()
	rdssvc = rds.New(cfgCp)
}
