package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go/aws"
)

var (
	rdssvc *rds.Client

	awsCfg aws.Config
)

func init() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	rdssvc = rds.New(cfg)

}

func getrds(filter string) {
	params := &rds.DescribeDBInstancesInput{}
	req := rdssvc.DescribeDBInstancesRequest(params)

	p := rds.NewDescribeDBInstancesPaginator(req)

	for p.Next(context.TODO()) {
		page := p.CurrentPage()

		for _, v := range page.DBInstances {
			fmt.Println(aws.StringValue(v.DBInstanceIdentifier))
		}

	}

}
