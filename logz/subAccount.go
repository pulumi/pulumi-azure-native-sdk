// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logz

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// API Version: 2020-10-01.
type SubAccount struct {
	pulumi.CustomResourceState

	Identity IdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	Location pulumi.StringOutput                 `pulumi:"location"`
	// Name of the monitor resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties specific to the monitor resource.
	Properties MonitorPropertiesResponseOutput `pulumi:"properties"`
	// The system metadata relating to this resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Tags       pulumi.StringMapOutput   `pulumi:"tags"`
	// The type of the monitor resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSubAccount registers a new resource with the given unique name, arguments, and options.
func NewSubAccount(ctx *pulumi.Context,
	name string, args *SubAccountArgs, opts ...pulumi.ResourceOption) (*SubAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MonitorName == nil {
		return nil, errors.New("invalid value for required argument 'MonitorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logz/v20201001:SubAccount"),
		},
		{
			Type: pulumi.String("azure-native:logz/v20201001preview:SubAccount"),
		},
		{
			Type: pulumi.String("azure-native:logz/v20220101preview:SubAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource SubAccount
	err := ctx.RegisterResource("azure-native:logz:SubAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubAccount gets an existing SubAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubAccountState, opts ...pulumi.ResourceOption) (*SubAccount, error) {
	var resource SubAccount
	err := ctx.ReadResource("azure-native:logz:SubAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubAccount resources.
type subAccountState struct {
}

type SubAccountState struct {
}

func (SubAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*subAccountState)(nil)).Elem()
}

type subAccountArgs struct {
	Identity *IdentityProperties `pulumi:"identity"`
	Location *string             `pulumi:"location"`
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// Properties specific to the monitor resource.
	Properties *MonitorProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sub Account resource name
	SubAccountName *string           `pulumi:"subAccountName"`
	Tags           map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SubAccount resource.
type SubAccountArgs struct {
	Identity IdentityPropertiesPtrInput
	Location pulumi.StringPtrInput
	// Monitor resource name
	MonitorName pulumi.StringInput
	// Properties specific to the monitor resource.
	Properties MonitorPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Sub Account resource name
	SubAccountName pulumi.StringPtrInput
	Tags           pulumi.StringMapInput
}

func (SubAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subAccountArgs)(nil)).Elem()
}

type SubAccountInput interface {
	pulumi.Input

	ToSubAccountOutput() SubAccountOutput
	ToSubAccountOutputWithContext(ctx context.Context) SubAccountOutput
}

func (*SubAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**SubAccount)(nil)).Elem()
}

func (i *SubAccount) ToSubAccountOutput() SubAccountOutput {
	return i.ToSubAccountOutputWithContext(context.Background())
}

func (i *SubAccount) ToSubAccountOutputWithContext(ctx context.Context) SubAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubAccountOutput)
}

type SubAccountOutput struct{ *pulumi.OutputState }

func (SubAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubAccount)(nil)).Elem()
}

func (o SubAccountOutput) ToSubAccountOutput() SubAccountOutput {
	return o
}

func (o SubAccountOutput) ToSubAccountOutputWithContext(ctx context.Context) SubAccountOutput {
	return o
}

func (o SubAccountOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *SubAccount) IdentityPropertiesResponsePtrOutput { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o SubAccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SubAccount) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the monitor resource.
func (o SubAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SubAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties specific to the monitor resource.
func (o SubAccountOutput) Properties() MonitorPropertiesResponseOutput {
	return o.ApplyT(func(v *SubAccount) MonitorPropertiesResponseOutput { return v.Properties }).(MonitorPropertiesResponseOutput)
}

// The system metadata relating to this resource
func (o SubAccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SubAccount) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SubAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SubAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the monitor resource.
func (o SubAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SubAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SubAccountOutput{})
}