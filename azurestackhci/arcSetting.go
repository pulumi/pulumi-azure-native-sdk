// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azurestackhci

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ArcSetting details.
//
// Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01.
//
// Other available API versions: 2022-12-15-preview, 2023-02-01, 2023-03-01, 2023-06-01, 2023-08-01, 2023-08-01-preview, 2023-11-01-preview, 2024-01-01, 2024-02-15-preview, 2024-09-01-preview, 2024-12-01-preview, 2025-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurestackhci [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ArcSetting struct {
	pulumi.CustomResourceState

	// Aggregate state of Arc agent across the nodes in this HCI cluster.
	AggregateState pulumi.StringOutput `pulumi:"aggregateState"`
	// App id of arc AAD identity.
	ArcApplicationClientId pulumi.StringPtrOutput `pulumi:"arcApplicationClientId"`
	// Object id of arc AAD identity.
	ArcApplicationObjectId pulumi.StringPtrOutput `pulumi:"arcApplicationObjectId"`
	// Tenant id of arc AAD identity.
	ArcApplicationTenantId pulumi.StringPtrOutput `pulumi:"arcApplicationTenantId"`
	// The resource group that hosts the Arc agents, ie. Hybrid Compute Machine resources.
	ArcInstanceResourceGroup pulumi.StringPtrOutput `pulumi:"arcInstanceResourceGroup"`
	// Object id of arc AAD service principal.
	ArcServicePrincipalObjectId pulumi.StringPtrOutput `pulumi:"arcServicePrincipalObjectId"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// contains connectivity related configuration for ARC resources
	ConnectivityProperties ArcConnectivityPropertiesResponseArrayOutput `pulumi:"connectivityProperties"`
	// Properties for each of the default extensions category
	DefaultExtensions DefaultExtensionDetailsResponseArrayOutput `pulumi:"defaultExtensions"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// State of Arc agent in each of the nodes.
	PerNodeDetails PerNodeStateResponseArrayOutput `pulumi:"perNodeDetails"`
	// Provisioning state of the ArcSetting proxy resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewArcSetting registers a new resource with the given unique name, arguments, and options.
func NewArcSetting(ctx *pulumi.Context,
	name string, args *ArcSettingArgs, opts ...pulumi.ResourceOption) (*ArcSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210101preview:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20210901preview:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220101:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220301:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220501:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20220901:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221001:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221201:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20221215preview:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230201:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230301:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230601:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230801:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20230801preview:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20231101preview:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20240101:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20240215preview:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20240401:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20240901preview:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20241201preview:ArcSetting"),
		},
		{
			Type: pulumi.String("azure-native:azurestackhci/v20250201preview:ArcSetting"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ArcSetting
	err := ctx.RegisterResource("azure-native:azurestackhci:ArcSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetArcSetting gets an existing ArcSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetArcSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArcSettingState, opts ...pulumi.ResourceOption) (*ArcSetting, error) {
	var resource ArcSetting
	err := ctx.ReadResource("azure-native:azurestackhci:ArcSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ArcSetting resources.
type arcSettingState struct {
}

type ArcSettingState struct {
}

func (ArcSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*arcSettingState)(nil)).Elem()
}

type arcSettingArgs struct {
	// App id of arc AAD identity.
	ArcApplicationClientId *string `pulumi:"arcApplicationClientId"`
	// Object id of arc AAD identity.
	ArcApplicationObjectId *string `pulumi:"arcApplicationObjectId"`
	// Tenant id of arc AAD identity.
	ArcApplicationTenantId *string `pulumi:"arcApplicationTenantId"`
	// The resource group that hosts the Arc agents, ie. Hybrid Compute Machine resources.
	ArcInstanceResourceGroup *string `pulumi:"arcInstanceResourceGroup"`
	// Object id of arc AAD service principal.
	ArcServicePrincipalObjectId *string `pulumi:"arcServicePrincipalObjectId"`
	// The name of the proxy resource holding details of HCI ArcSetting information.
	ArcSettingName *string `pulumi:"arcSettingName"`
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// contains connectivity related configuration for ARC resources
	ConnectivityProperties []ArcConnectivityProperties `pulumi:"connectivityProperties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ArcSetting resource.
type ArcSettingArgs struct {
	// App id of arc AAD identity.
	ArcApplicationClientId pulumi.StringPtrInput
	// Object id of arc AAD identity.
	ArcApplicationObjectId pulumi.StringPtrInput
	// Tenant id of arc AAD identity.
	ArcApplicationTenantId pulumi.StringPtrInput
	// The resource group that hosts the Arc agents, ie. Hybrid Compute Machine resources.
	ArcInstanceResourceGroup pulumi.StringPtrInput
	// Object id of arc AAD service principal.
	ArcServicePrincipalObjectId pulumi.StringPtrInput
	// The name of the proxy resource holding details of HCI ArcSetting information.
	ArcSettingName pulumi.StringPtrInput
	// The name of the cluster.
	ClusterName pulumi.StringInput
	// contains connectivity related configuration for ARC resources
	ConnectivityProperties ArcConnectivityPropertiesArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (ArcSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*arcSettingArgs)(nil)).Elem()
}

type ArcSettingInput interface {
	pulumi.Input

	ToArcSettingOutput() ArcSettingOutput
	ToArcSettingOutputWithContext(ctx context.Context) ArcSettingOutput
}

func (*ArcSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcSetting)(nil)).Elem()
}

func (i *ArcSetting) ToArcSettingOutput() ArcSettingOutput {
	return i.ToArcSettingOutputWithContext(context.Background())
}

func (i *ArcSetting) ToArcSettingOutputWithContext(ctx context.Context) ArcSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArcSettingOutput)
}

type ArcSettingOutput struct{ *pulumi.OutputState }

func (ArcSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcSetting)(nil)).Elem()
}

func (o ArcSettingOutput) ToArcSettingOutput() ArcSettingOutput {
	return o
}

func (o ArcSettingOutput) ToArcSettingOutputWithContext(ctx context.Context) ArcSettingOutput {
	return o
}

// Aggregate state of Arc agent across the nodes in this HCI cluster.
func (o ArcSettingOutput) AggregateState() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringOutput { return v.AggregateState }).(pulumi.StringOutput)
}

// App id of arc AAD identity.
func (o ArcSettingOutput) ArcApplicationClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringPtrOutput { return v.ArcApplicationClientId }).(pulumi.StringPtrOutput)
}

// Object id of arc AAD identity.
func (o ArcSettingOutput) ArcApplicationObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringPtrOutput { return v.ArcApplicationObjectId }).(pulumi.StringPtrOutput)
}

// Tenant id of arc AAD identity.
func (o ArcSettingOutput) ArcApplicationTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringPtrOutput { return v.ArcApplicationTenantId }).(pulumi.StringPtrOutput)
}

// The resource group that hosts the Arc agents, ie. Hybrid Compute Machine resources.
func (o ArcSettingOutput) ArcInstanceResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringPtrOutput { return v.ArcInstanceResourceGroup }).(pulumi.StringPtrOutput)
}

// Object id of arc AAD service principal.
func (o ArcSettingOutput) ArcServicePrincipalObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringPtrOutput { return v.ArcServicePrincipalObjectId }).(pulumi.StringPtrOutput)
}

// The Azure API version of the resource.
func (o ArcSettingOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// contains connectivity related configuration for ARC resources
func (o ArcSettingOutput) ConnectivityProperties() ArcConnectivityPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *ArcSetting) ArcConnectivityPropertiesResponseArrayOutput { return v.ConnectivityProperties }).(ArcConnectivityPropertiesResponseArrayOutput)
}

// Properties for each of the default extensions category
func (o ArcSettingOutput) DefaultExtensions() DefaultExtensionDetailsResponseArrayOutput {
	return o.ApplyT(func(v *ArcSetting) DefaultExtensionDetailsResponseArrayOutput { return v.DefaultExtensions }).(DefaultExtensionDetailsResponseArrayOutput)
}

// The name of the resource
func (o ArcSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// State of Arc agent in each of the nodes.
func (o ArcSettingOutput) PerNodeDetails() PerNodeStateResponseArrayOutput {
	return o.ApplyT(func(v *ArcSetting) PerNodeStateResponseArrayOutput { return v.PerNodeDetails }).(PerNodeStateResponseArrayOutput)
}

// Provisioning state of the ArcSetting proxy resource.
func (o ArcSettingOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ArcSettingOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ArcSetting) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ArcSettingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ArcSetting) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ArcSettingOutput{})
}
