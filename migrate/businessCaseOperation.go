// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Business case resource.
//
// Uses Azure REST API version 2024-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-04-01-preview.
//
// Other available API versions: 2023-04-01-preview, 2023-05-01-preview, 2023-09-09-preview, 2024-03-03-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type BusinessCaseOperation struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the last operation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Gets the state of business case reports.
	ReportStatusDetails ReportDetailsResponseArrayOutput `pulumi:"reportStatusDetails"`
	// Business case settings.
	Settings SettingsResponsePtrOutput `pulumi:"settings"`
	// Business case state.
	State pulumi.StringOutput `pulumi:"state"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBusinessCaseOperation registers a new resource with the given unique name, arguments, and options.
func NewBusinessCaseOperation(ctx *pulumi.Context,
	name string, args *BusinessCaseOperationArgs, opts ...pulumi.ResourceOption) (*BusinessCaseOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Settings != nil {
		args.Settings = args.Settings.ToSettingsPtrOutput().ApplyT(func(v *Settings) *Settings { return v.Defaults() }).(SettingsPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20230401preview:BusinessCaseOperation"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20230501preview:BusinessCaseOperation"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20230909preview:BusinessCaseOperation"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20240101preview:BusinessCaseOperation"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20240303preview:BusinessCaseOperation"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BusinessCaseOperation
	err := ctx.RegisterResource("azure-native:migrate:BusinessCaseOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBusinessCaseOperation gets an existing BusinessCaseOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBusinessCaseOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BusinessCaseOperationState, opts ...pulumi.ResourceOption) (*BusinessCaseOperation, error) {
	var resource BusinessCaseOperation
	err := ctx.ReadResource("azure-native:migrate:BusinessCaseOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BusinessCaseOperation resources.
type businessCaseOperationState struct {
}

type BusinessCaseOperationState struct {
}

func (BusinessCaseOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*businessCaseOperationState)(nil)).Elem()
}

type businessCaseOperationArgs struct {
	// Business case ARM name
	BusinessCaseName *string `pulumi:"businessCaseName"`
	// Assessment Project Name
	ProjectName string `pulumi:"projectName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Business case settings.
	Settings *Settings `pulumi:"settings"`
}

// The set of arguments for constructing a BusinessCaseOperation resource.
type BusinessCaseOperationArgs struct {
	// Business case ARM name
	BusinessCaseName pulumi.StringPtrInput
	// Assessment Project Name
	ProjectName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Business case settings.
	Settings SettingsPtrInput
}

func (BusinessCaseOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*businessCaseOperationArgs)(nil)).Elem()
}

type BusinessCaseOperationInput interface {
	pulumi.Input

	ToBusinessCaseOperationOutput() BusinessCaseOperationOutput
	ToBusinessCaseOperationOutputWithContext(ctx context.Context) BusinessCaseOperationOutput
}

func (*BusinessCaseOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**BusinessCaseOperation)(nil)).Elem()
}

func (i *BusinessCaseOperation) ToBusinessCaseOperationOutput() BusinessCaseOperationOutput {
	return i.ToBusinessCaseOperationOutputWithContext(context.Background())
}

func (i *BusinessCaseOperation) ToBusinessCaseOperationOutputWithContext(ctx context.Context) BusinessCaseOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BusinessCaseOperationOutput)
}

type BusinessCaseOperationOutput struct{ *pulumi.OutputState }

func (BusinessCaseOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BusinessCaseOperation)(nil)).Elem()
}

func (o BusinessCaseOperationOutput) ToBusinessCaseOperationOutput() BusinessCaseOperationOutput {
	return o
}

func (o BusinessCaseOperationOutput) ToBusinessCaseOperationOutputWithContext(ctx context.Context) BusinessCaseOperationOutput {
	return o
}

// The Azure API version of the resource.
func (o BusinessCaseOperationOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *BusinessCaseOperation) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The name of the resource
func (o BusinessCaseOperationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BusinessCaseOperation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the last operation.
func (o BusinessCaseOperationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *BusinessCaseOperation) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets the state of business case reports.
func (o BusinessCaseOperationOutput) ReportStatusDetails() ReportDetailsResponseArrayOutput {
	return o.ApplyT(func(v *BusinessCaseOperation) ReportDetailsResponseArrayOutput { return v.ReportStatusDetails }).(ReportDetailsResponseArrayOutput)
}

// Business case settings.
func (o BusinessCaseOperationOutput) Settings() SettingsResponsePtrOutput {
	return o.ApplyT(func(v *BusinessCaseOperation) SettingsResponsePtrOutput { return v.Settings }).(SettingsResponsePtrOutput)
}

// Business case state.
func (o BusinessCaseOperationOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *BusinessCaseOperation) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o BusinessCaseOperationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BusinessCaseOperation) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o BusinessCaseOperationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BusinessCaseOperation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BusinessCaseOperationOutput{})
}
