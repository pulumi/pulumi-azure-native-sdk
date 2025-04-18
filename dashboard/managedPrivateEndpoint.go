// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dashboard

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The managed private endpoint resource type.
//
// Uses Azure REST API version 2024-10-01. In version 2.x of the Azure Native provider, it used API version 2022-10-01-preview.
//
// Other available API versions: 2022-10-01-preview, 2023-09-01, 2023-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dashboard [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ManagedPrivateEndpoint struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The state of managed private endpoint connection.
	ConnectionState ManagedPrivateEndpointConnectionStateResponseOutput `pulumi:"connectionState"`
	// The group Ids of the managed private endpoint.
	GroupIds pulumi.StringArrayOutput `pulumi:"groupIds"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARM resource ID of the resource for which the managed private endpoint is pointing to.
	PrivateLinkResourceId pulumi.StringPtrOutput `pulumi:"privateLinkResourceId"`
	// The region of the resource to which the managed private endpoint is pointing to.
	PrivateLinkResourceRegion pulumi.StringPtrOutput `pulumi:"privateLinkResourceRegion"`
	// The private IP of private endpoint after approval. This property is empty before connection is approved.
	PrivateLinkServicePrivateIP pulumi.StringOutput `pulumi:"privateLinkServicePrivateIP"`
	// The URL of the data store behind the private link service. It would be the URL in the Grafana data source configuration page without the protocol and port.
	PrivateLinkServiceUrl pulumi.StringPtrOutput `pulumi:"privateLinkServiceUrl"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// User input request message of the managed private endpoint.
	RequestMessage pulumi.StringPtrOutput `pulumi:"requestMessage"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedPrivateEndpoint registers a new resource with the given unique name, arguments, and options.
func NewManagedPrivateEndpoint(ctx *pulumi.Context,
	name string, args *ManagedPrivateEndpointArgs, opts ...pulumi.ResourceOption) (*ManagedPrivateEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dashboard/v20221001preview:ManagedPrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:dashboard/v20230901:ManagedPrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:dashboard/v20231001preview:ManagedPrivateEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:dashboard/v20241001:ManagedPrivateEndpoint"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ManagedPrivateEndpoint
	err := ctx.RegisterResource("azure-native:dashboard:ManagedPrivateEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedPrivateEndpoint gets an existing ManagedPrivateEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedPrivateEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedPrivateEndpointState, opts ...pulumi.ResourceOption) (*ManagedPrivateEndpoint, error) {
	var resource ManagedPrivateEndpoint
	err := ctx.ReadResource("azure-native:dashboard:ManagedPrivateEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedPrivateEndpoint resources.
type managedPrivateEndpointState struct {
}

type ManagedPrivateEndpointState struct {
}

func (ManagedPrivateEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrivateEndpointState)(nil)).Elem()
}

type managedPrivateEndpointArgs struct {
	// The group Ids of the managed private endpoint.
	GroupIds []string `pulumi:"groupIds"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The managed private endpoint name of Azure Managed Grafana.
	ManagedPrivateEndpointName *string `pulumi:"managedPrivateEndpointName"`
	// The ARM resource ID of the resource for which the managed private endpoint is pointing to.
	PrivateLinkResourceId *string `pulumi:"privateLinkResourceId"`
	// The region of the resource to which the managed private endpoint is pointing to.
	PrivateLinkResourceRegion *string `pulumi:"privateLinkResourceRegion"`
	// The URL of the data store behind the private link service. It would be the URL in the Grafana data source configuration page without the protocol and port.
	PrivateLinkServiceUrl *string `pulumi:"privateLinkServiceUrl"`
	// User input request message of the managed private endpoint.
	RequestMessage *string `pulumi:"requestMessage"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The workspace name of Azure Managed Grafana.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ManagedPrivateEndpoint resource.
type ManagedPrivateEndpointArgs struct {
	// The group Ids of the managed private endpoint.
	GroupIds pulumi.StringArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The managed private endpoint name of Azure Managed Grafana.
	ManagedPrivateEndpointName pulumi.StringPtrInput
	// The ARM resource ID of the resource for which the managed private endpoint is pointing to.
	PrivateLinkResourceId pulumi.StringPtrInput
	// The region of the resource to which the managed private endpoint is pointing to.
	PrivateLinkResourceRegion pulumi.StringPtrInput
	// The URL of the data store behind the private link service. It would be the URL in the Grafana data source configuration page without the protocol and port.
	PrivateLinkServiceUrl pulumi.StringPtrInput
	// User input request message of the managed private endpoint.
	RequestMessage pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The workspace name of Azure Managed Grafana.
	WorkspaceName pulumi.StringInput
}

func (ManagedPrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedPrivateEndpointArgs)(nil)).Elem()
}

type ManagedPrivateEndpointInput interface {
	pulumi.Input

	ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput
	ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput
}

func (*ManagedPrivateEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPrivateEndpoint)(nil)).Elem()
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput {
	return i.ToManagedPrivateEndpointOutputWithContext(context.Background())
}

func (i *ManagedPrivateEndpoint) ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedPrivateEndpointOutput)
}

type ManagedPrivateEndpointOutput struct{ *pulumi.OutputState }

func (ManagedPrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedPrivateEndpoint)(nil)).Elem()
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointOutput() ManagedPrivateEndpointOutput {
	return o
}

func (o ManagedPrivateEndpointOutput) ToManagedPrivateEndpointOutputWithContext(ctx context.Context) ManagedPrivateEndpointOutput {
	return o
}

// The Azure API version of the resource.
func (o ManagedPrivateEndpointOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The state of managed private endpoint connection.
func (o ManagedPrivateEndpointOutput) ConnectionState() ManagedPrivateEndpointConnectionStateResponseOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) ManagedPrivateEndpointConnectionStateResponseOutput {
		return v.ConnectionState
	}).(ManagedPrivateEndpointConnectionStateResponseOutput)
}

// The group Ids of the managed private endpoint.
func (o ManagedPrivateEndpointOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringArrayOutput { return v.GroupIds }).(pulumi.StringArrayOutput)
}

// The geo-location where the resource lives
func (o ManagedPrivateEndpointOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o ManagedPrivateEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ARM resource ID of the resource for which the managed private endpoint is pointing to.
func (o ManagedPrivateEndpointOutput) PrivateLinkResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringPtrOutput { return v.PrivateLinkResourceId }).(pulumi.StringPtrOutput)
}

// The region of the resource to which the managed private endpoint is pointing to.
func (o ManagedPrivateEndpointOutput) PrivateLinkResourceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringPtrOutput { return v.PrivateLinkResourceRegion }).(pulumi.StringPtrOutput)
}

// The private IP of private endpoint after approval. This property is empty before connection is approved.
func (o ManagedPrivateEndpointOutput) PrivateLinkServicePrivateIP() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringOutput { return v.PrivateLinkServicePrivateIP }).(pulumi.StringOutput)
}

// The URL of the data store behind the private link service. It would be the URL in the Grafana data source configuration page without the protocol and port.
func (o ManagedPrivateEndpointOutput) PrivateLinkServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringPtrOutput { return v.PrivateLinkServiceUrl }).(pulumi.StringPtrOutput)
}

// Provisioning state of the resource.
func (o ManagedPrivateEndpointOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// User input request message of the managed private endpoint.
func (o ManagedPrivateEndpointOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringPtrOutput { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ManagedPrivateEndpointOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o ManagedPrivateEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ManagedPrivateEndpointOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedPrivateEndpoint) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedPrivateEndpointOutput{})
}
