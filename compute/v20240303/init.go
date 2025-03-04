// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240303

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
	case "azure-native:compute/v20240303:Gallery":
		r = &Gallery{}
	case "azure-native:compute/v20240303:GalleryApplication":
		r = &GalleryApplication{}
	case "azure-native:compute/v20240303:GalleryApplicationVersion":
		r = &GalleryApplicationVersion{}
	case "azure-native:compute/v20240303:GalleryImage":
		r = &GalleryImage{}
	case "azure-native:compute/v20240303:GalleryImageVersion":
		r = &GalleryImageVersion{}
	case "azure-native:compute/v20240303:GalleryInVMAccessControlProfile":
		r = &GalleryInVMAccessControlProfile{}
	case "azure-native:compute/v20240303:GalleryInVMAccessControlProfileVersion":
		r = &GalleryInVMAccessControlProfileVersion{}
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
		"compute/v20240303",
		&module{version},
	)
}
