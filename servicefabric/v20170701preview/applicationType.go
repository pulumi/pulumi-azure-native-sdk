// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The application type name resource
//
// Deprecated: Version 2017-07-01-preview will be removed in v2 of the provider.
type ApplicationType struct {
	pulumi.CustomResourceState

	// Azure resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Azure resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplicationType registers a new resource with the given unique name, arguments, and options.
func NewApplicationType(ctx *pulumi.Context,
	name string, args *ApplicationTypeArgs, opts ...pulumi.ResourceOption) (*ApplicationType, error) {
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
			Type: pulumi.String("azure-native:servicefabric:ApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190301:ApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190301preview:ApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20190601preview:ApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20191101preview:ApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20200301:ApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20201201preview:ApplicationType"),
		},
		{
			Type: pulumi.String("azure-native:servicefabric/v20210601:ApplicationType"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationType
	err := ctx.RegisterResource("azure-native:servicefabric/v20170701preview:ApplicationType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationType gets an existing ApplicationType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationTypeState, opts ...pulumi.ResourceOption) (*ApplicationType, error) {
	var resource ApplicationType
	err := ctx.ReadResource("azure-native:servicefabric/v20170701preview:ApplicationType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationType resources.
type applicationTypeState struct {
}

type ApplicationTypeState struct {
}

func (ApplicationTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationTypeState)(nil)).Elem()
}

type applicationTypeArgs struct {
	// The name of the application type name resource.
	ApplicationTypeName *string `pulumi:"applicationTypeName"`
	// The name of the cluster resource.
	ClusterName string `pulumi:"clusterName"`
	// Azure resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ApplicationType resource.
type ApplicationTypeArgs struct {
	// The name of the application type name resource.
	ApplicationTypeName pulumi.StringPtrInput
	// The name of the cluster resource.
	ClusterName pulumi.StringInput
	// Azure resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (ApplicationTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationTypeArgs)(nil)).Elem()
}

type ApplicationTypeInput interface {
	pulumi.Input

	ToApplicationTypeOutput() ApplicationTypeOutput
	ToApplicationTypeOutputWithContext(ctx context.Context) ApplicationTypeOutput
}

func (*ApplicationType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationType)(nil)).Elem()
}

func (i *ApplicationType) ToApplicationTypeOutput() ApplicationTypeOutput {
	return i.ToApplicationTypeOutputWithContext(context.Background())
}

func (i *ApplicationType) ToApplicationTypeOutputWithContext(ctx context.Context) ApplicationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTypeOutput)
}

type ApplicationTypeOutput struct{ *pulumi.OutputState }

func (ApplicationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationType)(nil)).Elem()
}

func (o ApplicationTypeOutput) ToApplicationTypeOutput() ApplicationTypeOutput {
	return o
}

func (o ApplicationTypeOutput) ToApplicationTypeOutputWithContext(ctx context.Context) ApplicationTypeOutput {
	return o
}

// Azure resource location.
func (o ApplicationTypeOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationType) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Azure resource name.
func (o ApplicationTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The current deployment or provisioning state, which only appears in the response.
func (o ApplicationTypeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationType) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure resource type.
func (o ApplicationTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationType) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationTypeOutput{})
}