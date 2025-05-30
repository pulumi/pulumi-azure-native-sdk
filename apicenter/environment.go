// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apicenter

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Environment entity.
//
// Uses Azure REST API version 2024-03-15-preview. In version 2.x of the Azure Native provider, it used API version 2024-03-01.
//
// Other available API versions: 2024-03-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apicenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type Environment struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The custom metadata defined for API catalog entities.
	CustomProperties pulumi.AnyOutput `pulumi:"customProperties"`
	// The environment description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Environment kind.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Environment onboarding information
	Onboarding OnboardingResponsePtrOutput `pulumi:"onboarding"`
	// Server information of the environment.
	Server EnvironmentServerResponsePtrOutput `pulumi:"server"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Environment title.
	Title pulumi.StringOutput `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apicenter/v20240301:Environment"),
		},
		{
			Type: pulumi.String("azure-native:apicenter/v20240315preview:Environment"),
		},
		{
			Type: pulumi.String("azure-native:apicenter/v20240601preview:Environment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Environment
	err := ctx.RegisterResource("azure-native:apicenter:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironment gets an existing Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("azure-native:apicenter:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Environment resources.
type environmentState struct {
}

type EnvironmentState struct {
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	// The custom metadata defined for API catalog entities.
	CustomProperties interface{} `pulumi:"customProperties"`
	// The environment description.
	Description *string `pulumi:"description"`
	// The name of the environment.
	EnvironmentName *string `pulumi:"environmentName"`
	// Environment kind.
	Kind string `pulumi:"kind"`
	// Environment onboarding information
	Onboarding *Onboarding `pulumi:"onboarding"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Server information of the environment.
	Server *EnvironmentServer `pulumi:"server"`
	// The name of Azure API Center service.
	ServiceName string `pulumi:"serviceName"`
	// Environment title.
	Title string `pulumi:"title"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	// The custom metadata defined for API catalog entities.
	CustomProperties pulumi.Input
	// The environment description.
	Description pulumi.StringPtrInput
	// The name of the environment.
	EnvironmentName pulumi.StringPtrInput
	// Environment kind.
	Kind pulumi.StringInput
	// Environment onboarding information
	Onboarding OnboardingPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Server information of the environment.
	Server EnvironmentServerPtrInput
	// The name of Azure API Center service.
	ServiceName pulumi.StringInput
	// Environment title.
	Title pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}

type EnvironmentInput interface {
	pulumi.Input

	ToEnvironmentOutput() EnvironmentOutput
	ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput
}

func (*Environment) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (i *Environment) ToEnvironmentOutput() EnvironmentOutput {
	return i.ToEnvironmentOutputWithContext(context.Background())
}

func (i *Environment) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentOutput)
}

type EnvironmentOutput struct{ *pulumi.OutputState }

func (EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (o EnvironmentOutput) ToEnvironmentOutput() EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return o
}

// The Azure API version of the resource.
func (o EnvironmentOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The custom metadata defined for API catalog entities.
func (o EnvironmentOutput) CustomProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Environment) pulumi.AnyOutput { return v.CustomProperties }).(pulumi.AnyOutput)
}

// The environment description.
func (o EnvironmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Environment kind.
func (o EnvironmentOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o EnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Environment onboarding information
func (o EnvironmentOutput) Onboarding() OnboardingResponsePtrOutput {
	return o.ApplyT(func(v *Environment) OnboardingResponsePtrOutput { return v.Onboarding }).(OnboardingResponsePtrOutput)
}

// Server information of the environment.
func (o EnvironmentOutput) Server() EnvironmentServerResponsePtrOutput {
	return o.ApplyT(func(v *Environment) EnvironmentServerResponsePtrOutput { return v.Server }).(EnvironmentServerResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o EnvironmentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Environment) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Environment title.
func (o EnvironmentOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o EnvironmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentOutput{})
}
