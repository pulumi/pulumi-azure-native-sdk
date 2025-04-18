// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package signalrservice

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes a Shared Private Link Resource
//
// Uses Azure REST API version 2024-03-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
//
// Other available API versions: 2023-02-01, 2023-03-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-08-01-preview, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native signalrservice [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type SignalRSharedPrivateLinkResource struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The group id from the provider of resource the shared private link resource is for
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource id of the resource the shared private link resource is for
	PrivateLinkResourceId pulumi.StringOutput `pulumi:"privateLinkResourceId"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The request message for requesting approval of the shared private link resource
	RequestMessage pulumi.StringPtrOutput `pulumi:"requestMessage"`
	// Status of the shared private link resource
	Status pulumi.StringOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSignalRSharedPrivateLinkResource registers a new resource with the given unique name, arguments, and options.
func NewSignalRSharedPrivateLinkResource(ctx *pulumi.Context,
	name string, args *SignalRSharedPrivateLinkResourceArgs, opts ...pulumi.ResourceOption) (*SignalRSharedPrivateLinkResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.PrivateLinkResourceId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:signalrservice/v20210401preview:SignalRSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210601preview:SignalRSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210901preview:SignalRSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20211001:SignalRSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20220201:SignalRSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20220801preview:SignalRSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230201:SignalRSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230301preview:SignalRSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230601preview:SignalRSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20230801preview:SignalRSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20240101preview:SignalRSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20240301:SignalRSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20240401preview:SignalRSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20240801preview:SignalRSharedPrivateLinkResource"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20241001preview:SignalRSharedPrivateLinkResource"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SignalRSharedPrivateLinkResource
	err := ctx.RegisterResource("azure-native:signalrservice:SignalRSharedPrivateLinkResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSignalRSharedPrivateLinkResource gets an existing SignalRSharedPrivateLinkResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSignalRSharedPrivateLinkResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SignalRSharedPrivateLinkResourceState, opts ...pulumi.ResourceOption) (*SignalRSharedPrivateLinkResource, error) {
	var resource SignalRSharedPrivateLinkResource
	err := ctx.ReadResource("azure-native:signalrservice:SignalRSharedPrivateLinkResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SignalRSharedPrivateLinkResource resources.
type signalRSharedPrivateLinkResourceState struct {
}

type SignalRSharedPrivateLinkResourceState struct {
}

func (SignalRSharedPrivateLinkResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRSharedPrivateLinkResourceState)(nil)).Elem()
}

type signalRSharedPrivateLinkResourceArgs struct {
	// The group id from the provider of resource the shared private link resource is for
	GroupId string `pulumi:"groupId"`
	// The resource id of the resource the shared private link resource is for
	PrivateLinkResourceId string `pulumi:"privateLinkResourceId"`
	// The request message for requesting approval of the shared private link resource
	RequestMessage *string `pulumi:"requestMessage"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource.
	ResourceName string `pulumi:"resourceName"`
	// The name of the shared private link resource.
	SharedPrivateLinkResourceName *string `pulumi:"sharedPrivateLinkResourceName"`
}

// The set of arguments for constructing a SignalRSharedPrivateLinkResource resource.
type SignalRSharedPrivateLinkResourceArgs struct {
	// The group id from the provider of resource the shared private link resource is for
	GroupId pulumi.StringInput
	// The resource id of the resource the shared private link resource is for
	PrivateLinkResourceId pulumi.StringInput
	// The request message for requesting approval of the shared private link resource
	RequestMessage pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the resource.
	ResourceName pulumi.StringInput
	// The name of the shared private link resource.
	SharedPrivateLinkResourceName pulumi.StringPtrInput
}

func (SignalRSharedPrivateLinkResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRSharedPrivateLinkResourceArgs)(nil)).Elem()
}

type SignalRSharedPrivateLinkResourceInput interface {
	pulumi.Input

	ToSignalRSharedPrivateLinkResourceOutput() SignalRSharedPrivateLinkResourceOutput
	ToSignalRSharedPrivateLinkResourceOutputWithContext(ctx context.Context) SignalRSharedPrivateLinkResourceOutput
}

func (*SignalRSharedPrivateLinkResource) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRSharedPrivateLinkResource)(nil)).Elem()
}

func (i *SignalRSharedPrivateLinkResource) ToSignalRSharedPrivateLinkResourceOutput() SignalRSharedPrivateLinkResourceOutput {
	return i.ToSignalRSharedPrivateLinkResourceOutputWithContext(context.Background())
}

func (i *SignalRSharedPrivateLinkResource) ToSignalRSharedPrivateLinkResourceOutputWithContext(ctx context.Context) SignalRSharedPrivateLinkResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRSharedPrivateLinkResourceOutput)
}

type SignalRSharedPrivateLinkResourceOutput struct{ *pulumi.OutputState }

func (SignalRSharedPrivateLinkResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRSharedPrivateLinkResource)(nil)).Elem()
}

func (o SignalRSharedPrivateLinkResourceOutput) ToSignalRSharedPrivateLinkResourceOutput() SignalRSharedPrivateLinkResourceOutput {
	return o
}

func (o SignalRSharedPrivateLinkResourceOutput) ToSignalRSharedPrivateLinkResourceOutputWithContext(ctx context.Context) SignalRSharedPrivateLinkResourceOutput {
	return o
}

// The Azure API version of the resource.
func (o SignalRSharedPrivateLinkResourceOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRSharedPrivateLinkResource) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The group id from the provider of resource the shared private link resource is for
func (o SignalRSharedPrivateLinkResourceOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRSharedPrivateLinkResource) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// The name of the resource
func (o SignalRSharedPrivateLinkResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRSharedPrivateLinkResource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource id of the resource the shared private link resource is for
func (o SignalRSharedPrivateLinkResourceOutput) PrivateLinkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRSharedPrivateLinkResource) pulumi.StringOutput { return v.PrivateLinkResourceId }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o SignalRSharedPrivateLinkResourceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRSharedPrivateLinkResource) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The request message for requesting approval of the shared private link resource
func (o SignalRSharedPrivateLinkResourceOutput) RequestMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalRSharedPrivateLinkResource) pulumi.StringPtrOutput { return v.RequestMessage }).(pulumi.StringPtrOutput)
}

// Status of the shared private link resource
func (o SignalRSharedPrivateLinkResourceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRSharedPrivateLinkResource) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SignalRSharedPrivateLinkResourceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SignalRSharedPrivateLinkResource) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SignalRSharedPrivateLinkResourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRSharedPrivateLinkResource) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SignalRSharedPrivateLinkResourceOutput{})
}
