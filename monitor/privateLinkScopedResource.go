// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A private link scoped resource
//
// Uses Azure REST API version 2023-06-01-preview.
//
// Other available API versions: 2021-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native monitor [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type PrivateLinkScopedResource struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The kind of scoped Azure monitor resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The resource id of the scoped Azure monitor resource.
	LinkedResourceId pulumi.StringPtrOutput `pulumi:"linkedResourceId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the Azure monitor resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The location of a scoped subscription. Only needs to be specified for metric dataplane subscriptions.
	SubscriptionLocation pulumi.StringPtrOutput `pulumi:"subscriptionLocation"`
	// System data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateLinkScopedResource registers a new resource with the given unique name, arguments, and options.
func NewPrivateLinkScopedResource(ctx *pulumi.Context,
	name string, args *PrivateLinkScopedResourceArgs, opts ...pulumi.ResourceOption) (*PrivateLinkScopedResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScopeName == nil {
		return nil, errors.New("invalid value for required argument 'ScopeName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights/v20210701preview:PrivateLinkScopedResource"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20210901:PrivateLinkScopedResource"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20230601preview:PrivateLinkScopedResource"),
		},
		{
			Type: pulumi.String("azure-native:insights:PrivateLinkScopedResource"),
		},
		{
			Type: pulumi.String("azure-native:monitor/v20191017preview:PrivateLinkScopedResource"),
		},
		{
			Type: pulumi.String("azure-native:monitor/v20210701preview:PrivateLinkScopedResource"),
		},
		{
			Type: pulumi.String("azure-native:monitor/v20210901:PrivateLinkScopedResource"),
		},
		{
			Type: pulumi.String("azure-native:monitor/v20230601preview:PrivateLinkScopedResource"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateLinkScopedResource
	err := ctx.RegisterResource("azure-native:monitor:PrivateLinkScopedResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateLinkScopedResource gets an existing PrivateLinkScopedResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateLinkScopedResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkScopedResourceState, opts ...pulumi.ResourceOption) (*PrivateLinkScopedResource, error) {
	var resource PrivateLinkScopedResource
	err := ctx.ReadResource("azure-native:monitor:PrivateLinkScopedResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateLinkScopedResource resources.
type privateLinkScopedResourceState struct {
}

type PrivateLinkScopedResourceState struct {
}

func (PrivateLinkScopedResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkScopedResourceState)(nil)).Elem()
}

type privateLinkScopedResourceArgs struct {
	// The kind of scoped Azure monitor resource.
	Kind *string `pulumi:"kind"`
	// The resource id of the scoped Azure monitor resource.
	LinkedResourceId *string `pulumi:"linkedResourceId"`
	// The name of the scoped resource object.
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Azure Monitor PrivateLinkScope resource.
	ScopeName string `pulumi:"scopeName"`
	// The location of a scoped subscription. Only needs to be specified for metric dataplane subscriptions.
	SubscriptionLocation *string `pulumi:"subscriptionLocation"`
}

// The set of arguments for constructing a PrivateLinkScopedResource resource.
type PrivateLinkScopedResourceArgs struct {
	// The kind of scoped Azure monitor resource.
	Kind pulumi.StringPtrInput
	// The resource id of the scoped Azure monitor resource.
	LinkedResourceId pulumi.StringPtrInput
	// The name of the scoped resource object.
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Azure Monitor PrivateLinkScope resource.
	ScopeName pulumi.StringInput
	// The location of a scoped subscription. Only needs to be specified for metric dataplane subscriptions.
	SubscriptionLocation pulumi.StringPtrInput
}

func (PrivateLinkScopedResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkScopedResourceArgs)(nil)).Elem()
}

type PrivateLinkScopedResourceInput interface {
	pulumi.Input

	ToPrivateLinkScopedResourceOutput() PrivateLinkScopedResourceOutput
	ToPrivateLinkScopedResourceOutputWithContext(ctx context.Context) PrivateLinkScopedResourceOutput
}

func (*PrivateLinkScopedResource) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkScopedResource)(nil)).Elem()
}

func (i *PrivateLinkScopedResource) ToPrivateLinkScopedResourceOutput() PrivateLinkScopedResourceOutput {
	return i.ToPrivateLinkScopedResourceOutputWithContext(context.Background())
}

func (i *PrivateLinkScopedResource) ToPrivateLinkScopedResourceOutputWithContext(ctx context.Context) PrivateLinkScopedResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkScopedResourceOutput)
}

type PrivateLinkScopedResourceOutput struct{ *pulumi.OutputState }

func (PrivateLinkScopedResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkScopedResource)(nil)).Elem()
}

func (o PrivateLinkScopedResourceOutput) ToPrivateLinkScopedResourceOutput() PrivateLinkScopedResourceOutput {
	return o
}

func (o PrivateLinkScopedResourceOutput) ToPrivateLinkScopedResourceOutputWithContext(ctx context.Context) PrivateLinkScopedResourceOutput {
	return o
}

// The Azure API version of the resource.
func (o PrivateLinkScopedResourceOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScopedResource) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The kind of scoped Azure monitor resource.
func (o PrivateLinkScopedResourceOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkScopedResource) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The resource id of the scoped Azure monitor resource.
func (o PrivateLinkScopedResourceOutput) LinkedResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkScopedResource) pulumi.StringPtrOutput { return v.LinkedResourceId }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o PrivateLinkScopedResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScopedResource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// State of the Azure monitor resource.
func (o PrivateLinkScopedResourceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScopedResource) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The location of a scoped subscription. Only needs to be specified for metric dataplane subscriptions.
func (o PrivateLinkScopedResourceOutput) SubscriptionLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkScopedResource) pulumi.StringPtrOutput { return v.SubscriptionLocation }).(pulumi.StringPtrOutput)
}

// System data
func (o PrivateLinkScopedResourceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateLinkScopedResource) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateLinkScopedResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScopedResource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkScopedResourceOutput{})
}
