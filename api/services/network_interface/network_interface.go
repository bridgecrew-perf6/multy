package network_interface

import (
	"github.com/multycloud/multy/api/proto/common"
	"github.com/multycloud/multy/api/proto/resources"
	"github.com/multycloud/multy/api/services"
	"github.com/multycloud/multy/api/util"
	"github.com/multycloud/multy/db"
	"github.com/multycloud/multy/resources/output"
)

type NetworkInterfaceService struct {
	Service services.Service[*resources.CloudSpecificNetworkInterfaceArgs, *resources.NetworkInterfaceResource]
}

func (s NetworkInterfaceService) Convert(resourceId string, args []*resources.CloudSpecificNetworkInterfaceArgs, state *output.TfState) (*resources.NetworkInterfaceResource, error) {
	var result []*resources.CloudSpecificNetworkInterfaceResource
	for _, r := range args {
		result = append(result, &resources.CloudSpecificNetworkInterfaceResource{
			CommonParameters: util.ConvertCommonParams(r.CommonParameters),
			Name:             r.Name,
			SubnetId:         r.SubnetId,
		})
	}

	return &resources.NetworkInterfaceResource{
		CommonParameters: &common.CommonResourceParameters{ResourceId: resourceId},
		Resources:        result,
	}, nil
}

func NewNetworkInterfaceService(database *db.Database) NetworkInterfaceService {
	ni := NetworkInterfaceService{
		Service: services.Service[*resources.CloudSpecificNetworkInterfaceArgs, *resources.NetworkInterfaceResource]{
			Db:         database,
			Converters: nil,
		},
	}
	ni.Service.Converters = &ni
	return ni
}