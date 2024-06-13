package awsgo

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func InitializeAWS() {
	Ctx = context.TODO() // There is not limitation at execution time
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))
	if err != nil {
		panic(fmt.Errorf("error loading config(./aws/config): %s", err.Error()))
	}
}
