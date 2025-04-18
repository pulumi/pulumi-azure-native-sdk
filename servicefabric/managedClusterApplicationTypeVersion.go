// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicefabric

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An application type version resource for the specified application type name resource.
//
// Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01-preview.
//
// Other available API versions: 2023-03-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01-preview, 2024-02-01-preview, 2024-06-01-preview, 2024-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native servicefabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ManagedClusterApplicationTypeVersion struct {
	pulumi.CustomResourceState

	// The URL to the application package
	AppPackageUrl pulumi.StringOutput `pulumi:"appPackageUrl"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Resource location depends on the parent resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current deployment or provisioning state, which only appears in the response
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Azure resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedClusterApplicationTypeVersion registers a new resource with the given unique name, arguments, and options.
func NewManagedClusterApplicationTypeVersion(ctx *pulumi.Context,
	name string, args *ManagedClusterApplicationTypeVersionArgs, opts ...pulumi.ResourceOption) (*ManagedClusterApplicationTypeVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppPackageUrl == nil {
		return nil, errors.New("invalid value for required argument 'AppPackageUrl'")
	}
	if args.ApplicationTypeName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationTypeName'")
	}
	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicefabric/v20210101preview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210501:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210701preview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210901privatepreview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20211101preview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220101:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220201preview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220601preview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20220801preview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20221001preview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20230201preview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20230301preview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20230701preview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20230901preview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20231101preview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20231201preview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20240201preview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20240401:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20240601preview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20240901preview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20241101preview:ManagedClusterApplicationTypeVersion"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20250301preview:ManagedClusterApplicationTypeVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedClusterApplicationTypeVersion
	err := ctx.RegisterResource("azure-native:servicefabric:ManagedClusterApplicationTypeVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedClusterApplicationTypeVersion gets an existing ManagedClusterApplicationTypeVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedClusterApplicationTypeVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedClusterApplicationTypeVersionState, opts ...pulumi.ResourceOption) (*ManagedClusterApplicationTypeVersion, error) {
	var resource ManagedClusterApplicationTypeVersion
	err := ctx.ReadResource("azure-native:servicefabric:ManagedClusterApplicationTypeVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedClusterApplicationTypeVersion resources.
type managedClusterApplicationTypeVersionState struct {
}

type ManagedClusterApplicationTypeVersionState struct {
}

func (ManagedClusterApplicationTypeVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterApplicationTypeVersionState)(nil)).Elem()
}

type managedClusterApplicationTypeVersionArgs struct {
	// The URL to the application package
	AppPackageUrl string `pulumi:"appPackageUrl"`
	// The name of the application type name resource.
	ApplicationTypeName string `pulumi:"applicationTypeName"`
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// Resource location depends on the parent resource.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The application type version.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a ManagedClusterApplicationTypeVersion resource.
type ManagedClusterApplicationTypeVersionArgs struct {
	// The URL to the application package
	AppPackageUrl pulumi.StringInput
	// The name of the application type name resource.
	ApplicationTypeName pulumi.StringInput
	// The name of the cluster resource.
	ClusterName pulumi.StringInput
	// Resource location depends on the parent resource.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Azure resource tags.
	Tags pulumi.StringMapInput
	// The application type version.
	Version pulumi.StringPtrInput
}

func (ManagedClusterApplicationTypeVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterApplicationTypeVersionArgs)(nil)).Elem()
}

type ManagedClusterApplicationTypeVersionInput interface {
	pulumi.Input

	ToManagedClusterApplicationTypeVersionOutput() ManagedClusterApplicationTypeVersionOutput
	ToManagedClusterApplicationTypeVersionOutputWithContext(ctx context.Context) ManagedClusterApplicationTypeVersionOutput
}

func (*ManagedClusterApplicationTypeVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterApplicationTypeVersion)(nil)).Elem()
}

func (i *ManagedClusterApplicationTypeVersion) ToManagedClusterApplicationTypeVersionOutput() ManagedClusterApplicationTypeVersionOutput {
	return i.ToManagedClusterApplicationTypeVersionOutputWithContext(context.Background())
}

func (i *ManagedClusterApplicationTypeVersion) ToManagedClusterApplicationTypeVersionOutputWithContext(ctx context.Context) ManagedClusterApplicationTypeVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedClusterApplicationTypeVersionOutput)
}

type ManagedClusterApplicationTypeVersionOutput struct{ *pulumi.OutputState }

func (ManagedClusterApplicationTypeVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedClusterApplicationTypeVersion)(nil)).Elem()
}

func (o ManagedClusterApplicationTypeVersionOutput) ToManagedClusterApplicationTypeVersionOutput() ManagedClusterApplicationTypeVersionOutput {
	return o
}

func (o ManagedClusterApplicationTypeVersionOutput) ToManagedClusterApplicationTypeVersionOutputWithContext(ctx context.Context) ManagedClusterApplicationTypeVersionOutput {
	return o
}

// The URL to the application package
func (o ManagedClusterApplicationTypeVersionOutput) AppPackageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterApplicationTypeVersion) pulumi.StringOutput { return v.AppPackageUrl }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o ManagedClusterApplicationTypeVersionOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterApplicationTypeVersion) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Resource location depends on the parent resource.
func (o ManagedClusterApplicationTypeVersionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedClusterApplicationTypeVersion) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Azure resource name.
func (o ManagedClusterApplicationTypeVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterApplicationTypeVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current deployment or provisioning state, which only appears in the response
func (o ManagedClusterApplicationTypeVersionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterApplicationTypeVersion) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o ManagedClusterApplicationTypeVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagedClusterApplicationTypeVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Azure resource tags.
func (o ManagedClusterApplicationTypeVersionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedClusterApplicationTypeVersion) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Azure resource type.
func (o ManagedClusterApplicationTypeVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedClusterApplicationTypeVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedClusterApplicationTypeVersionOutput{})
}
