// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20250401preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The shared dashboard resource definition.
type Dashboard struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties DashboardPropertiesWithProvisioningStateResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDashboard registers a new resource with the given unique name, arguments, and options.
func NewDashboard(ctx *pulumi.Context,
	name string, args *DashboardArgs, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:portal/v20150801preview:Dashboard"),
		},
		{
			Type: pulumi.String("azure-native:portal/v20181001preview:Dashboard"),
		},
		{
			Type: pulumi.String("azure-native:portal/v20190101preview:Dashboard"),
		},
		{
			Type: pulumi.String("azure-native:portal/v20200901preview:Dashboard"),
		},
		{
			Type: pulumi.String("azure-native:portal/v20221201preview:Dashboard"),
		},
		{
			Type: pulumi.String("azure-native:portal:Dashboard"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Dashboard
	err := ctx.RegisterResource("azure-native:portal/v20250401preview:Dashboard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboard gets an existing Dashboard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardState, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	var resource Dashboard
	err := ctx.ReadResource("azure-native:portal/v20250401preview:Dashboard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dashboard resources.
type dashboardState struct {
}

type DashboardState struct {
}

func (DashboardState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardState)(nil)).Elem()
}

type dashboardArgs struct {
	// The name of the dashboard.
	DashboardName *string `pulumi:"dashboardName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The resource-specific properties for this resource.
	Properties *DashboardPropertiesWithProvisioningState `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Dashboard resource.
type DashboardArgs struct {
	// The name of the dashboard.
	DashboardName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The resource-specific properties for this resource.
	Properties DashboardPropertiesWithProvisioningStatePtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (DashboardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardArgs)(nil)).Elem()
}

type DashboardInput interface {
	pulumi.Input

	ToDashboardOutput() DashboardOutput
	ToDashboardOutputWithContext(ctx context.Context) DashboardOutput
}

func (*Dashboard) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil)).Elem()
}

func (i *Dashboard) ToDashboardOutput() DashboardOutput {
	return i.ToDashboardOutputWithContext(context.Background())
}

func (i *Dashboard) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardOutput)
}

type DashboardOutput struct{ *pulumi.OutputState }

func (DashboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil)).Elem()
}

func (o DashboardOutput) ToDashboardOutput() DashboardOutput {
	return o
}

func (o DashboardOutput) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return o
}

// The geo-location where the resource lives
func (o DashboardOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o DashboardOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o DashboardOutput) Properties() DashboardPropertiesWithProvisioningStateResponseOutput {
	return o.ApplyT(func(v *Dashboard) DashboardPropertiesWithProvisioningStateResponseOutput { return v.Properties }).(DashboardPropertiesWithProvisioningStateResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o DashboardOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Dashboard) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o DashboardOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DashboardOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Dashboard) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DashboardOutput{})
}
