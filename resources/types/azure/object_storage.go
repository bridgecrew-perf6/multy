package azure_resources

import (
	"fmt"
	"github.com/multycloud/multy/api/proto/commonpb"
	"github.com/multycloud/multy/api/proto/resourcespb"
	"github.com/multycloud/multy/resources"
	"github.com/multycloud/multy/resources/common"
	"github.com/multycloud/multy/resources/output"
	"github.com/multycloud/multy/resources/output/object_storage"
	"github.com/multycloud/multy/resources/output/object_storage_object"
	"github.com/multycloud/multy/resources/types"
)

type AzureObjectStorage struct {
	*types.ObjectStorage
}

func InitObjectStorage(vn *types.ObjectStorage) resources.ResourceTranslator[*resourcespb.ObjectStorageResource] {
	return AzureObjectStorage{vn}
}

func (r AzureObjectStorage) FromState(state *output.TfState) (*resourcespb.ObjectStorageResource, error) {
	return &resourcespb.ObjectStorageResource{
		CommonParameters: &commonpb.CommonResourceParameters{
			ResourceId:      r.ResourceId,
			ResourceGroupId: r.Args.CommonParameters.ResourceGroupId,
			Location:        r.Args.CommonParameters.Location,
			CloudProvider:   r.Args.CommonParameters.CloudProvider,
			NeedsUpdate:     false,
		},
		Name:       r.Args.Name,
		Versioning: r.Args.Versioning,
	}, nil
}

func (r AzureObjectStorage) Translate(resources.MultyContext) ([]output.TfBlock, error) {
	rgName := GetResourceGroupName(r.Args.CommonParameters.ResourceGroupId)

	storageAccount := object_storage.AzureStorageAccount{
		AzResource: common.NewAzResource(
			r.ResourceId, common.RemoveSpecialChars(r.Args.Name), rgName,
			r.GetCloudSpecificLocation(),
		),
		AccountTier:                "Standard",
		AccountReplicationType:     "GZRS",
		AllowNestedItemsToBePublic: true,
		BlobProperties: object_storage.BlobProperties{
			VersioningEnabled: r.Args.Versioning,
		},
	}

	return []output.TfBlock{
		storageAccount,
		object_storage_object.AzureStorageContainer{
			AzResource: &common.AzResource{
				TerraformResource: output.TerraformResource{
					ResourceId: fmt.Sprintf("%s_%s", r.ResourceId, "public"),
				},
				Name: "public",
			},
			StorageAccountName:  storageAccount.GetResourceName(),
			ContainerAccessType: "blob",
		}, object_storage_object.AzureStorageContainer{
			AzResource: &common.AzResource{
				TerraformResource: output.TerraformResource{
					ResourceId: fmt.Sprintf("%s_%s", r.ResourceId, "private"),
				},
				Name: "private",
			},
			StorageAccountName:  storageAccount.GetResourceName(),
			ContainerAccessType: "private",
		}}, nil
}

func (r AzureObjectStorage) GetMainResourceName() (string, error) {
	return object_storage.AzureResourceName, nil
}
