// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package operationalinsights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Log Analytics QueryPack definition.
// API Version: 2019-09-01.
type QueryPack struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Current state of this QueryPack: whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The unique ID of your application. This field cannot be changed.
	QueryPackId pulumi.StringOutput `pulumi:"queryPackId"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Creation Date for the Log Analytics QueryPack, in ISO 8601 format.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// Last modified date of the Log Analytics QueryPack, in ISO 8601 format.
	TimeModified pulumi.StringOutput `pulumi:"timeModified"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewQueryPack registers a new resource with the given unique name, arguments, and options.
func NewQueryPack(ctx *pulumi.Context,
	name string, args *QueryPackArgs, opts ...pulumi.ResourceOption) (*QueryPack, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights/v20190901:QueryPack"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20190901preview:QueryPack"),
		},
	})
	opts = append(opts, aliases)
	var resource QueryPack
	err := ctx.RegisterResource("azure-native:operationalinsights:QueryPack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueryPack gets an existing QueryPack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueryPack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueryPackState, opts ...pulumi.ResourceOption) (*QueryPack, error) {
	var resource QueryPack
	err := ctx.ReadResource("azure-native:operationalinsights:QueryPack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueryPack resources.
type queryPackState struct {
}

type QueryPackState struct {
}

func (QueryPackState) ElementType() reflect.Type {
	return reflect.TypeOf((*queryPackState)(nil)).Elem()
}

type queryPackArgs struct {
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the Log Analytics QueryPack resource.
	QueryPackName *string `pulumi:"queryPackName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a QueryPack resource.
type QueryPackArgs struct {
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the Log Analytics QueryPack resource.
	QueryPackName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (QueryPackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queryPackArgs)(nil)).Elem()
}

type QueryPackInput interface {
	pulumi.Input

	ToQueryPackOutput() QueryPackOutput
	ToQueryPackOutputWithContext(ctx context.Context) QueryPackOutput
}

func (*QueryPack) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryPack)(nil)).Elem()
}

func (i *QueryPack) ToQueryPackOutput() QueryPackOutput {
	return i.ToQueryPackOutputWithContext(context.Background())
}

func (i *QueryPack) ToQueryPackOutputWithContext(ctx context.Context) QueryPackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryPackOutput)
}

type QueryPackOutput struct{ *pulumi.OutputState }

func (QueryPackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryPack)(nil)).Elem()
}

func (o QueryPackOutput) ToQueryPackOutput() QueryPackOutput {
	return o
}

func (o QueryPackOutput) ToQueryPackOutputWithContext(ctx context.Context) QueryPackOutput {
	return o
}

// Resource location
func (o QueryPackOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name
func (o QueryPackOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Current state of this QueryPack: whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
func (o QueryPackOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The unique ID of your application. This field cannot be changed.
func (o QueryPackOutput) QueryPackId() pulumi.StringOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringOutput { return v.QueryPackId }).(pulumi.StringOutput)
}

// Resource tags
func (o QueryPackOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Creation Date for the Log Analytics QueryPack, in ISO 8601 format.
func (o QueryPackOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringOutput { return v.TimeCreated }).(pulumi.StringOutput)
}

// Last modified date of the Log Analytics QueryPack, in ISO 8601 format.
func (o QueryPackOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringOutput { return v.TimeModified }).(pulumi.StringOutput)
}

// Azure resource type
func (o QueryPackOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(QueryPackOutput{})
}