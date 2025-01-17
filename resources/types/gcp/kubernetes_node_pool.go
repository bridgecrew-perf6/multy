package gcp_resources

import (
	"fmt"
	"github.com/multycloud/multy/api/proto/commonpb"
	"github.com/multycloud/multy/api/proto/resourcespb"
	"github.com/multycloud/multy/resources"
	"github.com/multycloud/multy/resources/common"
	"github.com/multycloud/multy/resources/output"
	"github.com/multycloud/multy/resources/output/iam"
	"github.com/multycloud/multy/resources/output/kubernetes_node_pool"
	"github.com/multycloud/multy/resources/types"
)

type GcpKubernetesNodePool struct {
	*types.KubernetesNodePool
}

func InitKubernetesNodePool(r *types.KubernetesNodePool) resources.ResourceTranslator[*resourcespb.KubernetesNodePoolResource] {
	return GcpKubernetesNodePool{r}
}

func (r GcpKubernetesNodePool) FromState(state *output.TfState) (*resourcespb.KubernetesNodePoolResource, error) {
	return &resourcespb.KubernetesNodePoolResource{
		CommonParameters: &commonpb.CommonChildResourceParameters{
			ResourceId:  r.ResourceId,
			NeedsUpdate: false,
		},
		Name:              r.Args.Name,
		SubnetId:          r.Args.SubnetId,
		ClusterId:         r.Args.ClusterId,
		StartingNodeCount: r.Args.StartingNodeCount,
		MinNodeCount:      r.Args.MinNodeCount,
		MaxNodeCount:      r.Args.MaxNodeCount,
		VmSize:            r.Args.VmSize,
		DiskSizeGb:        r.Args.DiskSizeGb,
		AvailabilityZone:  r.Args.AvailabilityZone,
		AwsOverride:       r.Args.AwsOverride,
		AzureOverride:     r.Args.AzureOverride,
		Labels:            r.Args.Labels,
	}, nil
}

func (r GcpKubernetesNodePool) Translate(_ resources.MultyContext) ([]output.TfBlock, error) {
	clusterId, err := resources.GetMainOutputId(GcpKubernetesCluster{r.KubernetesCluster})
	if err != nil {
		return nil, err
	}
	size, err := common.GetVmSize(r.Args.VmSize, r.GetCloud())
	if err != nil {
		return nil, err
	}

	numZones := 3
	if len(r.Args.AvailabilityZone) > 0 {
		numZones = len(r.Args.AvailabilityZone)
	}

	var zones []string
	for _, zone := range r.Args.AvailabilityZone {
		availabilityZone, err := common.GetAvailabilityZone(r.KubernetesCluster.GetLocation(), int(zone), r.GetCloud())
		if err != nil {
			return nil, err
		}
		zones = append(zones, availabilityZone)
	}

	nodePool := &kubernetes_node_pool.GoogleContainerNodePool{
		GcpResource:      common.NewGcpResource(r.ResourceId, r.Args.Name, r.KubernetesCluster.Args.GetGcpOverride().GetProject()),
		Cluster:          clusterId,
		NodeLocations:    zones,
		InitialNodeCount: int(r.Args.StartingNodeCount) / numZones,
		Autoscaling: kubernetes_node_pool.GoogleContainerNodePoolAutoScaling{
			MinNodeCount: int(r.Args.MinNodeCount) / numZones,
			MaxNodeCount: int(r.Args.MaxNodeCount) / numZones,
		},
		NodeConfig: kubernetes_node_pool.GoogleContainerNodeConfig{
			DiskSizeGb:  int(r.Args.DiskSizeGb),
			Labels:      r.Args.Labels,
			MachineType: size,
			Tags:        []string{GcpSubnet{r.Subnet}.getNetworkTag()},
			// Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
			ServiceAccount: fmt.Sprintf("%s.%s.email", output.GetResourceName(iam.GoogleServiceAccount{}), GcpKubernetesCluster{r.KubernetesCluster}.getServiceAccountId()),
			OAuthScopes:    []string{"https://www.googleapis.com/auth/cloud-platform"},
		},
	}

	return []output.TfBlock{nodePool}, nil

}

func (r GcpKubernetesNodePool) GetMainResourceName() (string, error) {
	return output.GetResourceName(kubernetes_node_pool.GoogleContainerNodePool{}), nil
}
