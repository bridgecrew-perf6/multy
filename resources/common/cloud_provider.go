package common

import (
	"fmt"
	"github.com/multycloud/multy/api/proto/commonpb"
	"github.com/multycloud/multy/api/proto/resourcespb"
	"github.com/multycloud/multy/resources/output"
)

const (
	AWS   = commonpb.CloudProvider_AWS
	AZURE = commonpb.CloudProvider_AZURE
	GCP   = commonpb.CloudProvider_GCP
)

var LOCATION = map[commonpb.Location]map[commonpb.CloudProvider]string{
	// Ireland
	commonpb.Location_EU_WEST_1: {
		AWS:   "eu-west-1",
		AZURE: "northeurope",
		GCP:   "europe-west1", // Belgium
	},
	// London
	commonpb.Location_EU_WEST_2: {
		AWS:   "eu-west-2",
		AZURE: "uksouth",
		GCP:   "europe-west2",
	},
	// France
	commonpb.Location_EU_WEST_3: {
		AWS:   "eu-west-3",
		AZURE: "francecentral",
		GCP:   "europe-west9",
	},
	// N. Virginia
	commonpb.Location_US_EAST_1: {
		AWS:   "us-east-1",
		AZURE: "eastus",
		GCP:   "us-east4", // Virginia
	},
	// Ohio
	commonpb.Location_US_EAST_2: {
		AWS:   "us-east-2",
		AZURE: "eastus2",
		GCP:   "us-east5",
	},
	// California
	commonpb.Location_US_WEST_1: {
		AWS:   "us-west-1",
		AZURE: "westus2",
		GCP:   "us-west2",
	},
	// Oregon
	commonpb.Location_US_WEST_2: {
		AWS:   "us-west-2",
		AZURE: "westus3",
		GCP:   "us-west1",
	},
	// Sweden
	commonpb.Location_EU_NORTH_1: {
		AWS:   "eu-north-1",
		AZURE: "swedencentral",
		GCP:   "europe-north1", // Finland
	},
}

var AVAILABILITY_ZONES = map[commonpb.Location]map[commonpb.CloudProvider][]string{
	commonpb.Location_EU_WEST_1: {
		AWS:   []string{"eu-west-1a", "eu-west-1b", "eu-west-1c"},
		AZURE: []string{"1", "2", "3"},
		GCP:   []string{"europe-west1-b", "europe-west1-c", "europe-west1-d"},
	},
	commonpb.Location_EU_WEST_2: {
		AWS:   []string{"eu-west-2a", "eu-west-2b", "eu-west-2c"},
		AZURE: []string{"1", "2", "3"},
		GCP:   []string{"europe-west2-a", "europe-west2-b", "europe-west2-c"},
	},
	commonpb.Location_EU_WEST_3: {
		AWS:   []string{"eu-west-3a", "eu-west-3b", "eu-west-3c"},
		AZURE: []string{"1", "2", "3"},
		GCP:   []string{"europe-west3-a", "europe-west3-b", "europe-west3-c"},
	},
	commonpb.Location_US_EAST_1: {
		AWS:   []string{"us-east-1a", "us-east-1b", "us-east-1c"},
		AZURE: []string{"1", "2", "3"},
		GCP:   []string{"us-east4-a", "us-east4-b", "us-east4-c"},
	},
	commonpb.Location_US_EAST_2: {
		AWS:   []string{"us-east-2a", "us-east-2b", "us-east-2c"},
		AZURE: []string{"1", "2", "3"},
		GCP:   []string{"us-east5-a", "us-east5-b", "us-east5-c"},
	},
	commonpb.Location_US_WEST_1: {
		AWS:   []string{"us-west-1a", "us-west-1c"},
		AZURE: []string{"1", "2"},
		GCP:   []string{"us-west2-a", "us-west2-b", "us-west2-c"},
	},
	commonpb.Location_US_WEST_2: {
		AWS:   []string{"us-west-2a", "us-west-2b", "us-west-2c", "us-west-2d"},
		AZURE: []string{"1", "2", "3"},
		GCP:   []string{"us-west1-a", "us-west1-b", "us-west1-c"},
	},
	commonpb.Location_EU_NORTH_1: {
		AWS:   []string{"eu-north-1a", "eu-north-1b", "eu-north-1c"},
		AZURE: []string{"1", "2", "3"},
		GCP:   []string{"europe-north1-a", "europe-north1-b", "europe-north1-c"},
	},
}

// eu-west-2 "ami-0fc15d50d39e4503c"
// amzn2-ami-hvm-2.0.20211103.0-x86_64-gp2
//https://cloud-images.ubuntu.com/locator/ec2/
var AMIMAP = map[string]string{
	"eu-west-1":  "ami-09d4a659cdd8677be",
	"eu-west-2":  "ami-0fc15d50d39e4503c",
	"eu-west-3":  "ami-0fc15d50d39e4503c",
	"us-east-1":  "ami-04ad2567c9e3d7893",
	"us-east-2":  "ami-04ad2567c9e3d7893",
	"us-west-1":  "ami-04ad2567c9e3d7893",
	"us-west-2":  "ami-04ad2567c9e3d7893",
	"eu-north-1": "ami-04ad2567c9e3d7893",
}

var AwsAmiOwners = map[resourcespb.ImageReference_OperatingSystemDistribution]string{
	resourcespb.ImageReference_UBUNTU:  "099720109477",
	resourcespb.ImageReference_CENT_OS: "125523088429",
	resourcespb.ImageReference_DEBIAN:  "136693071363",
}

func GetAvailabilityZone(location commonpb.Location, az int, cloud commonpb.CloudProvider) (string, error) {
	if AVAILABILITY_ZONES[location] == nil {
		return "", fmt.Errorf("invalid location: %s", location)
	}
	azArray := AVAILABILITY_ZONES[location][cloud]
	if az == 0 {
		return "", nil
	}
	if az <= len(azArray) {
		return azArray[az-1], nil
	}
	return "", fmt.Errorf("invalid az value: %d", az)
}

func GetCloudLocation(location commonpb.Location, provider commonpb.CloudProvider) (string, error) {
	if _, ok := LOCATION[location]; !ok {
		return "", fmt.Errorf("location %s is not defined", location)
	}
	if _, ok := LOCATION[location][provider]; !ok {
		return "", fmt.Errorf("location %s is not defined for cloud %s", location, provider)
	}
	return LOCATION[location][provider], nil
}

type AzResource struct {
	output.TerraformResource `hcl:",squash"`
	ResourceGroupName        string `hcl:"resource_group_name,expr" hcle:"omitempty"`
	Name                     string `hcl:"name" hcle:"omitempty"`
	Location                 string `hcl:"location" hcle:"omitempty"`
}

func NewAwsResource(resourceId string, name string) *AwsResource {
	return &AwsResource{
		TerraformResource: output.TerraformResource{ResourceId: resourceId},
		Tags:              map[string]string{"Name": name}}
}

func NewAwsResourceWithDeps(resourceId string, name string, deps []string) *AwsResource {
	return &AwsResource{
		TerraformResource: output.TerraformResource{ResourceId: resourceId, DependsOn: deps},
		Tags:              map[string]string{"Name": name}}
}

func NewAwsResourceWithIdOnly(resourceId string) *AwsResource {
	return &AwsResource{
		TerraformResource: output.TerraformResource{ResourceId: resourceId}}
}

func NewAzResource(resourceId string, name string, rgName string, location string) *AzResource {
	return &AzResource{
		TerraformResource: output.TerraformResource{ResourceId: resourceId},
		Name:              name,
		ResourceGroupName: rgName,
		Location:          location,
	}
}

func NewGcpResource(resourceId string, name string, project string) *GcpResource {
	return &GcpResource{
		TerraformResource: output.TerraformResource{ResourceId: resourceId},
		Name:              name,
		Project:           project,
	}
}

func (r *AwsResource) SetName(name string) {
	//r.ResourceName = name
	r.TerraformResource.ResourceName = name
}

func (r *AzResource) SetName(name string) {
	r.TerraformResource.ResourceName = name
}

func (r *GcpResource) SetName(name string) {
	r.TerraformResource.ResourceName = name
}

type AwsResource struct {
	output.TerraformResource `hcl:",squash"`
	Tags                     map[string]string `hcl:"tags" hcle:"omitempty"`
}

type GcpResource struct {
	output.TerraformResource `hcl:",squash"`
	Name                     string `hcl:"name" hcle:"omitempty"`
	Project                  string `hcl:"project"`
}
