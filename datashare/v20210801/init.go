// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210801

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native-sdk"
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
	case "azure-native:datashare/v20210801:ADLSGen1FileDataSet":
		r = &ADLSGen1FileDataSet{}
	case "azure-native:datashare/v20210801:ADLSGen1FolderDataSet":
		r = &ADLSGen1FolderDataSet{}
	case "azure-native:datashare/v20210801:ADLSGen2FileDataSet":
		r = &ADLSGen2FileDataSet{}
	case "azure-native:datashare/v20210801:ADLSGen2FileDataSetMapping":
		r = &ADLSGen2FileDataSetMapping{}
	case "azure-native:datashare/v20210801:ADLSGen2FileSystemDataSet":
		r = &ADLSGen2FileSystemDataSet{}
	case "azure-native:datashare/v20210801:ADLSGen2FileSystemDataSetMapping":
		r = &ADLSGen2FileSystemDataSetMapping{}
	case "azure-native:datashare/v20210801:ADLSGen2FolderDataSet":
		r = &ADLSGen2FolderDataSet{}
	case "azure-native:datashare/v20210801:ADLSGen2FolderDataSetMapping":
		r = &ADLSGen2FolderDataSetMapping{}
	case "azure-native:datashare/v20210801:Account":
		r = &Account{}
	case "azure-native:datashare/v20210801:BlobContainerDataSet":
		r = &BlobContainerDataSet{}
	case "azure-native:datashare/v20210801:BlobContainerDataSetMapping":
		r = &BlobContainerDataSetMapping{}
	case "azure-native:datashare/v20210801:BlobDataSet":
		r = &BlobDataSet{}
	case "azure-native:datashare/v20210801:BlobDataSetMapping":
		r = &BlobDataSetMapping{}
	case "azure-native:datashare/v20210801:BlobFolderDataSet":
		r = &BlobFolderDataSet{}
	case "azure-native:datashare/v20210801:BlobFolderDataSetMapping":
		r = &BlobFolderDataSetMapping{}
	case "azure-native:datashare/v20210801:DataSet":
		r = &DataSet{}
	case "azure-native:datashare/v20210801:DataSetMapping":
		r = &DataSetMapping{}
	case "azure-native:datashare/v20210801:Invitation":
		r = &Invitation{}
	case "azure-native:datashare/v20210801:KustoClusterDataSet":
		r = &KustoClusterDataSet{}
	case "azure-native:datashare/v20210801:KustoClusterDataSetMapping":
		r = &KustoClusterDataSetMapping{}
	case "azure-native:datashare/v20210801:KustoDatabaseDataSet":
		r = &KustoDatabaseDataSet{}
	case "azure-native:datashare/v20210801:KustoDatabaseDataSetMapping":
		r = &KustoDatabaseDataSetMapping{}
	case "azure-native:datashare/v20210801:KustoTableDataSet":
		r = &KustoTableDataSet{}
	case "azure-native:datashare/v20210801:KustoTableDataSetMapping":
		r = &KustoTableDataSetMapping{}
	case "azure-native:datashare/v20210801:ScheduledSynchronizationSetting":
		r = &ScheduledSynchronizationSetting{}
	case "azure-native:datashare/v20210801:ScheduledTrigger":
		r = &ScheduledTrigger{}
	case "azure-native:datashare/v20210801:Share":
		r = &Share{}
	case "azure-native:datashare/v20210801:ShareSubscription":
		r = &ShareSubscription{}
	case "azure-native:datashare/v20210801:SqlDBTableDataSet":
		r = &SqlDBTableDataSet{}
	case "azure-native:datashare/v20210801:SqlDBTableDataSetMapping":
		r = &SqlDBTableDataSetMapping{}
	case "azure-native:datashare/v20210801:SqlDWTableDataSet":
		r = &SqlDWTableDataSet{}
	case "azure-native:datashare/v20210801:SqlDWTableDataSetMapping":
		r = &SqlDWTableDataSetMapping{}
	case "azure-native:datashare/v20210801:SynapseWorkspaceSqlPoolTableDataSet":
		r = &SynapseWorkspaceSqlPoolTableDataSet{}
	case "azure-native:datashare/v20210801:SynapseWorkspaceSqlPoolTableDataSetMapping":
		r = &SynapseWorkspaceSqlPoolTableDataSetMapping{}
	case "azure-native:datashare/v20210801:SynchronizationSetting":
		r = &SynchronizationSetting{}
	case "azure-native:datashare/v20210801:Trigger":
		r = &Trigger{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := pulumiazurenativesdk.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"datashare/v20210801",
		&module{version},
	)
}