// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:storage/v20230401:BlobContainer":
		r = &BlobContainer{}
	case "azure-native:storage/v20230401:BlobContainerImmutabilityPolicy":
		r = &BlobContainerImmutabilityPolicy{}
	case "azure-native:storage/v20230401:BlobInventoryPolicy":
		r = &BlobInventoryPolicy{}
	case "azure-native:storage/v20230401:BlobServiceProperties":
		r = &BlobServiceProperties{}
	case "azure-native:storage/v20230401:EncryptionScope":
		r = &EncryptionScope{}
	case "azure-native:storage/v20230401:FileServiceProperties":
		r = &FileServiceProperties{}
	case "azure-native:storage/v20230401:FileShare":
		r = &FileShare{}
	case "azure-native:storage/v20230401:LocalUser":
		r = &LocalUser{}
	case "azure-native:storage/v20230401:ManagementPolicy":
		r = &ManagementPolicy{}
	case "azure-native:storage/v20230401:ObjectReplicationPolicy":
		r = &ObjectReplicationPolicy{}
	case "azure-native:storage/v20230401:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:storage/v20230401:Queue":
		r = &Queue{}
	case "azure-native:storage/v20230401:QueueServiceProperties":
		r = &QueueServiceProperties{}
	case "azure-native:storage/v20230401:StorageAccount":
		r = &StorageAccount{}
	case "azure-native:storage/v20230401:Table":
		r = &Table{}
	case "azure-native:storage/v20230401:TableServiceProperties":
		r = &TableServiceProperties{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := utilities.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"storage/v20230401",
		&module{version},
	)
}
