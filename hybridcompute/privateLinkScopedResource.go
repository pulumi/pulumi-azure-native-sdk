// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hybridcompute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A private link scoped resource
//
// Uses Azure REST API version 2020-08-15-preview. In version 2.x of the Azure Native provider, it used API version 2020-08-15-preview.
type PrivateLinkScopedResource struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The resource id of the scoped Azure monitor resource.
	LinkedResourceId pulumi.StringPtrOutput `pulumi:"linkedResourceId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the private endpoint connection.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
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
			Type: pulumi.String("azure-native:hybridcompute/v20200815preview:PrivateLinkScopedResource"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource PrivateLinkScopedResource
	err := ctx.RegisterResource("azure-native:hybridcompute:PrivateLinkScopedResource", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:hybridcompute:PrivateLinkScopedResource", name, id, state, &resource, opts...)
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
	// The resource id of the scoped Azure monitor resource.
	LinkedResourceId *string `pulumi:"linkedResourceId"`
	// The name of the scoped resource object.
	Name *string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Azure Arc PrivateLinkScope resource.
	ScopeName string `pulumi:"scopeName"`
}

// The set of arguments for constructing a PrivateLinkScopedResource resource.
type PrivateLinkScopedResourceArgs struct {
	// The resource id of the scoped Azure monitor resource.
	LinkedResourceId pulumi.StringPtrInput
	// The name of the scoped resource object.
	Name pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the Azure Arc PrivateLinkScope resource.
	ScopeName pulumi.StringInput
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

// The resource id of the scoped Azure monitor resource.
func (o PrivateLinkScopedResourceOutput) LinkedResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkScopedResource) pulumi.StringPtrOutput { return v.LinkedResourceId }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o PrivateLinkScopedResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScopedResource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// State of the private endpoint connection.
func (o PrivateLinkScopedResourceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScopedResource) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PrivateLinkScopedResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateLinkScopedResource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkScopedResourceOutput{})
}
