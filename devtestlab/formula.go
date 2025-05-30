// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devtestlab

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A formula for creating a VM, specifying an image base and other parameters
//
// Uses Azure REST API version 2018-09-15. In version 2.x of the Azure Native provider, it used API version 2018-09-15.
type Formula struct {
	pulumi.CustomResourceState

	// The author of the formula.
	Author pulumi.StringOutput `pulumi:"author"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The creation date of the formula.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The description of the formula.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The content of the formula.
	FormulaContent LabVirtualMachineCreationParameterResponsePtrOutput `pulumi:"formulaContent"`
	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The OS type of the formula.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringOutput `pulumi:"uniqueIdentifier"`
	// Information about a VM from which a formula is to be created.
	Vm FormulaPropertiesFromVmResponsePtrOutput `pulumi:"vm"`
}

// NewFormula registers a new resource with the given unique name, arguments, and options.
func NewFormula(ctx *pulumi.Context,
	name string, args *FormulaArgs, opts ...pulumi.ResourceOption) (*Formula, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.FormulaContent != nil {
		args.FormulaContent = args.FormulaContent.ToLabVirtualMachineCreationParameterPtrOutput().ApplyT(func(v *LabVirtualMachineCreationParameter) *LabVirtualMachineCreationParameter { return v.Defaults() }).(LabVirtualMachineCreationParameterPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:Formula"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:Formula"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:Formula"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Formula
	err := ctx.RegisterResource("azure-native:devtestlab:Formula", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFormula gets an existing Formula resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFormula(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FormulaState, opts ...pulumi.ResourceOption) (*Formula, error) {
	var resource Formula
	err := ctx.ReadResource("azure-native:devtestlab:Formula", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Formula resources.
type formulaState struct {
}

type FormulaState struct {
}

func (FormulaState) ElementType() reflect.Type {
	return reflect.TypeOf((*formulaState)(nil)).Elem()
}

type formulaArgs struct {
	// The description of the formula.
	Description *string `pulumi:"description"`
	// The content of the formula.
	FormulaContent *LabVirtualMachineCreationParameter `pulumi:"formulaContent"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the formula.
	Name *string `pulumi:"name"`
	// The OS type of the formula.
	OsType *string `pulumi:"osType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Information about a VM from which a formula is to be created.
	Vm *FormulaPropertiesFromVm `pulumi:"vm"`
}

// The set of arguments for constructing a Formula resource.
type FormulaArgs struct {
	// The description of the formula.
	Description pulumi.StringPtrInput
	// The content of the formula.
	FormulaContent LabVirtualMachineCreationParameterPtrInput
	// The name of the lab.
	LabName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the formula.
	Name pulumi.StringPtrInput
	// The OS type of the formula.
	OsType pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// Information about a VM from which a formula is to be created.
	Vm FormulaPropertiesFromVmPtrInput
}

func (FormulaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*formulaArgs)(nil)).Elem()
}

type FormulaInput interface {
	pulumi.Input

	ToFormulaOutput() FormulaOutput
	ToFormulaOutputWithContext(ctx context.Context) FormulaOutput
}

func (*Formula) ElementType() reflect.Type {
	return reflect.TypeOf((**Formula)(nil)).Elem()
}

func (i *Formula) ToFormulaOutput() FormulaOutput {
	return i.ToFormulaOutputWithContext(context.Background())
}

func (i *Formula) ToFormulaOutputWithContext(ctx context.Context) FormulaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FormulaOutput)
}

type FormulaOutput struct{ *pulumi.OutputState }

func (FormulaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Formula)(nil)).Elem()
}

func (o FormulaOutput) ToFormulaOutput() FormulaOutput {
	return o
}

func (o FormulaOutput) ToFormulaOutputWithContext(ctx context.Context) FormulaOutput {
	return o
}

// The author of the formula.
func (o FormulaOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringOutput { return v.Author }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o FormulaOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The creation date of the formula.
func (o FormulaOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// The description of the formula.
func (o FormulaOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The content of the formula.
func (o FormulaOutput) FormulaContent() LabVirtualMachineCreationParameterResponsePtrOutput {
	return o.ApplyT(func(v *Formula) LabVirtualMachineCreationParameterResponsePtrOutput { return v.FormulaContent }).(LabVirtualMachineCreationParameterResponsePtrOutput)
}

// The location of the resource.
func (o FormulaOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource.
func (o FormulaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The OS type of the formula.
func (o FormulaOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

// The provisioning status of the resource.
func (o FormulaOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The tags of the resource.
func (o FormulaOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o FormulaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The unique immutable identifier of a resource (Guid).
func (o FormulaOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Formula) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

// Information about a VM from which a formula is to be created.
func (o FormulaOutput) Vm() FormulaPropertiesFromVmResponsePtrOutput {
	return o.ApplyT(func(v *Formula) FormulaPropertiesFromVmResponsePtrOutput { return v.Vm }).(FormulaPropertiesFromVmResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(FormulaOutput{})
}
