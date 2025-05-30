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

// Definition of the Runtime Environment type.
//
// Uses Azure REST API version 2023-05-15-preview. In version 2.x of the Azure Native provider, it used API version 2023-05-15-preview.
//
// Other available API versions: 2024-10-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native automation [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type RuntimeEnvironment struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// List of Default packages for Environment
	DefaultPackages pulumi.StringMapOutput `pulumi:"defaultPackages"`
	// Gets or sets the description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Language of Runtime Environment
	Language pulumi.StringPtrOutput `pulumi:"language"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Version of Language
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewRuntimeEnvironment registers a new resource with the given unique name, arguments, and options.
func NewRuntimeEnvironment(ctx *pulumi.Context,
	name string, args *RuntimeEnvironmentArgs, opts ...pulumi.ResourceOption) (*RuntimeEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation/v20230515preview:RuntimeEnvironment"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20241023:RuntimeEnvironment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RuntimeEnvironment
	err := ctx.RegisterResource("azure-native:automation:RuntimeEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuntimeEnvironment gets an existing RuntimeEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuntimeEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuntimeEnvironmentState, opts ...pulumi.ResourceOption) (*RuntimeEnvironment, error) {
	var resource RuntimeEnvironment
	err := ctx.ReadResource("azure-native:automation:RuntimeEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuntimeEnvironment resources.
type runtimeEnvironmentState struct {
}

type RuntimeEnvironmentState struct {
}

func (RuntimeEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeEnvironmentState)(nil)).Elem()
}

type runtimeEnvironmentArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// List of Default packages for Environment
	DefaultPackages map[string]string `pulumi:"defaultPackages"`
	// Gets or sets the description.
	Description *string `pulumi:"description"`
	// Language of Runtime Environment
	Language *string `pulumi:"language"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Runtime Environment.
	RuntimeEnvironmentName *string `pulumi:"runtimeEnvironmentName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Version of Language
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a RuntimeEnvironment resource.
type RuntimeEnvironmentArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// List of Default packages for Environment
	DefaultPackages pulumi.StringMapInput
	// Gets or sets the description.
	Description pulumi.StringPtrInput
	// Language of Runtime Environment
	Language pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the Runtime Environment.
	RuntimeEnvironmentName pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Version of Language
	Version pulumi.StringPtrInput
}

func (RuntimeEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeEnvironmentArgs)(nil)).Elem()
}

type RuntimeEnvironmentInput interface {
	pulumi.Input

	ToRuntimeEnvironmentOutput() RuntimeEnvironmentOutput
	ToRuntimeEnvironmentOutputWithContext(ctx context.Context) RuntimeEnvironmentOutput
}

func (*RuntimeEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeEnvironment)(nil)).Elem()
}

func (i *RuntimeEnvironment) ToRuntimeEnvironmentOutput() RuntimeEnvironmentOutput {
	return i.ToRuntimeEnvironmentOutputWithContext(context.Background())
}

func (i *RuntimeEnvironment) ToRuntimeEnvironmentOutputWithContext(ctx context.Context) RuntimeEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeEnvironmentOutput)
}

type RuntimeEnvironmentOutput struct{ *pulumi.OutputState }

func (RuntimeEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeEnvironment)(nil)).Elem()
}

func (o RuntimeEnvironmentOutput) ToRuntimeEnvironmentOutput() RuntimeEnvironmentOutput {
	return o
}

func (o RuntimeEnvironmentOutput) ToRuntimeEnvironmentOutputWithContext(ctx context.Context) RuntimeEnvironmentOutput {
	return o
}

// The Azure API version of the resource.
func (o RuntimeEnvironmentOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeEnvironment) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// List of Default packages for Environment
func (o RuntimeEnvironmentOutput) DefaultPackages() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RuntimeEnvironment) pulumi.StringMapOutput { return v.DefaultPackages }).(pulumi.StringMapOutput)
}

// Gets or sets the description.
func (o RuntimeEnvironmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuntimeEnvironment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Language of Runtime Environment
func (o RuntimeEnvironmentOutput) Language() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuntimeEnvironment) pulumi.StringPtrOutput { return v.Language }).(pulumi.StringPtrOutput)
}

// The geo-location where the resource lives
func (o RuntimeEnvironmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeEnvironment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o RuntimeEnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeEnvironment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o RuntimeEnvironmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RuntimeEnvironment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o RuntimeEnvironmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RuntimeEnvironment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o RuntimeEnvironmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeEnvironment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Version of Language
func (o RuntimeEnvironmentOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RuntimeEnvironment) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(RuntimeEnvironmentOutput{})
}
