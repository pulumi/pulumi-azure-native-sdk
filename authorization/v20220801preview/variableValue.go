// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The variable value.
type VariableValue struct {
	pulumi.CustomResourceState

	// The name of the variable.
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource (Microsoft.Authorization/variables/values).
	Type pulumi.StringOutput `pulumi:"type"`
	// Variable value column value array.
	Values PolicyVariableValueColumnValueResponseArrayOutput `pulumi:"values"`
}

// NewVariableValue registers a new resource with the given unique name, arguments, and options.
func NewVariableValue(ctx *pulumi.Context,
	name string, args *VariableValueArgs, opts ...pulumi.ResourceOption) (*VariableValue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Values == nil {
		return nil, errors.New("invalid value for required argument 'Values'")
	}
	if args.VariableName == nil {
		return nil, errors.New("invalid value for required argument 'VariableName'")
	}
	var resource VariableValue
	err := ctx.RegisterResource("azure-native:authorization/v20220801preview:VariableValue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVariableValue gets an existing VariableValue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVariableValue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VariableValueState, opts ...pulumi.ResourceOption) (*VariableValue, error) {
	var resource VariableValue
	err := ctx.ReadResource("azure-native:authorization/v20220801preview:VariableValue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VariableValue resources.
type variableValueState struct {
}

type VariableValueState struct {
}

func (VariableValueState) ElementType() reflect.Type {
	return reflect.TypeOf((*variableValueState)(nil)).Elem()
}

type variableValueArgs struct {
	// Variable value column value array.
	Values []PolicyVariableValueColumnValue `pulumi:"values"`
	// The name of the variable to operate on.
	VariableName string `pulumi:"variableName"`
	// The name of the variable value to operate on.
	VariableValueName *string `pulumi:"variableValueName"`
}

// The set of arguments for constructing a VariableValue resource.
type VariableValueArgs struct {
	// Variable value column value array.
	Values PolicyVariableValueColumnValueArrayInput
	// The name of the variable to operate on.
	VariableName pulumi.StringInput
	// The name of the variable value to operate on.
	VariableValueName pulumi.StringPtrInput
}

func (VariableValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*variableValueArgs)(nil)).Elem()
}

type VariableValueInput interface {
	pulumi.Input

	ToVariableValueOutput() VariableValueOutput
	ToVariableValueOutputWithContext(ctx context.Context) VariableValueOutput
}

func (*VariableValue) ElementType() reflect.Type {
	return reflect.TypeOf((**VariableValue)(nil)).Elem()
}

func (i *VariableValue) ToVariableValueOutput() VariableValueOutput {
	return i.ToVariableValueOutputWithContext(context.Background())
}

func (i *VariableValue) ToVariableValueOutputWithContext(ctx context.Context) VariableValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableValueOutput)
}

type VariableValueOutput struct{ *pulumi.OutputState }

func (VariableValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VariableValue)(nil)).Elem()
}

func (o VariableValueOutput) ToVariableValueOutput() VariableValueOutput {
	return o
}

func (o VariableValueOutput) ToVariableValueOutputWithContext(ctx context.Context) VariableValueOutput {
	return o
}

// The name of the variable.
func (o VariableValueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VariableValue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o VariableValueOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VariableValue) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource (Microsoft.Authorization/variables/values).
func (o VariableValueOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VariableValue) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Variable value column value array.
func (o VariableValueOutput) Values() PolicyVariableValueColumnValueResponseArrayOutput {
	return o.ApplyT(func(v *VariableValue) PolicyVariableValueColumnValueResponseArrayOutput { return v.Values }).(PolicyVariableValueColumnValueResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VariableValueOutput{})
}