// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the variable.
//
// Uses Azure REST API version 2023-11-01. In version 2.x of the Azure Native provider, it used API version 2022-08-08.
//
// Other available API versions: 2015-10-31, 2019-06-01, 2020-01-13-preview, 2022-08-08, 2023-05-15-preview, 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type Variable struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Gets or sets the creation time.
	CreationTime pulumi.StringPtrOutput `pulumi:"creationTime"`
	// Gets or sets the description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gets or sets the encrypted flag of the variable.
	IsEncrypted pulumi.BoolPtrOutput `pulumi:"isEncrypted"`
	// Gets or sets the last modified time.
	LastModifiedTime pulumi.StringPtrOutput `pulumi:"lastModifiedTime"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets or sets the value of the variable.
	Value pulumi.StringPtrOutput `pulumi:"value"`
}

// NewVariable registers a new resource with the given unique name, arguments, and options.
func NewVariable(ctx *pulumi.Context,
	name string, args *VariableArgs, opts ...pulumi.ResourceOption) (*Variable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation/v20151031:Variable"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:Variable"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:Variable"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20220808:Variable"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20230515preview:Variable"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20231101:Variable"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20241023:Variable"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Variable
	err := ctx.RegisterResource("azure-native:automation:Variable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVariable gets an existing Variable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VariableState, opts ...pulumi.ResourceOption) (*Variable, error) {
	var resource Variable
	err := ctx.ReadResource("azure-native:automation:Variable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Variable resources.
type variableState struct {
}

type VariableState struct {
}

func (VariableState) ElementType() reflect.Type {
	return reflect.TypeOf((*variableState)(nil)).Elem()
}

type variableArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Gets or sets the description of the variable.
	Description *string `pulumi:"description"`
	// Gets or sets the encrypted flag of the variable.
	IsEncrypted *bool `pulumi:"isEncrypted"`
	// Gets or sets the name of the variable.
	Name string `pulumi:"name"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets the value of the variable.
	Value *string `pulumi:"value"`
	// The variable name.
	VariableName *string `pulumi:"variableName"`
}

// The set of arguments for constructing a Variable resource.
type VariableArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// Gets or sets the description of the variable.
	Description pulumi.StringPtrInput
	// Gets or sets the encrypted flag of the variable.
	IsEncrypted pulumi.BoolPtrInput
	// Gets or sets the name of the variable.
	Name pulumi.StringInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// Gets or sets the value of the variable.
	Value pulumi.StringPtrInput
	// The variable name.
	VariableName pulumi.StringPtrInput
}

func (VariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*variableArgs)(nil)).Elem()
}

type VariableInput interface {
	pulumi.Input

	ToVariableOutput() VariableOutput
	ToVariableOutputWithContext(ctx context.Context) VariableOutput
}

func (*Variable) ElementType() reflect.Type {
	return reflect.TypeOf((**Variable)(nil)).Elem()
}

func (i *Variable) ToVariableOutput() VariableOutput {
	return i.ToVariableOutputWithContext(context.Background())
}

func (i *Variable) ToVariableOutputWithContext(ctx context.Context) VariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VariableOutput)
}

type VariableOutput struct{ *pulumi.OutputState }

func (VariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Variable)(nil)).Elem()
}

func (o VariableOutput) ToVariableOutput() VariableOutput {
	return o
}

func (o VariableOutput) ToVariableOutputWithContext(ctx context.Context) VariableOutput {
	return o
}

// The Azure API version of the resource.
func (o VariableOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Gets or sets the creation time.
func (o VariableOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringPtrOutput { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// Gets or sets the description.
func (o VariableOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Gets or sets the encrypted flag of the variable.
func (o VariableOutput) IsEncrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Variable) pulumi.BoolPtrOutput { return v.IsEncrypted }).(pulumi.BoolPtrOutput)
}

// Gets or sets the last modified time.
func (o VariableOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringPtrOutput { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o VariableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of the resource.
func (o VariableOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets the value of the variable.
func (o VariableOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Variable) pulumi.StringPtrOutput { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VariableOutput{})
}
