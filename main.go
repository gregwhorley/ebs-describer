package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"log"
)

func main() {
	var profile, credFile string

	flag.StringVar(&profile, "p", "default", "AWS Profile name")
	flag.StringVar(&credFile, "c", "", "Absolute path to AWS credentials file")
	flag.Parse()

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(profile),
		config.WithSharedCredentialsFiles([]string{credFile}),
	)
	if err != nil {
		log.Fatalf("Failed to load config, %v", err)
	}

	client := ec2.NewFromConfig(cfg)

	response, err := client.DescribeVolumes(context.TODO(), &ec2.DescribeVolumesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("status"),
				Values: []string{"available"},
			},
		},
		MaxResults: aws.Int32(100),
	})
	if err != nil {
		log.Fatalf("Error while requesting describe volumes, %v", err)
	}

	fmt.Println("Volume creation time:")
	for _, volume := range response.Volumes {
		fmt.Println(volume.CreateTime)
	}

}
