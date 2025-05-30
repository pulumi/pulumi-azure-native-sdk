// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dashboard

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The integration fabric resource type.
//
// Uses Azure REST API version 2024-10-01. In version 2.x of the Azure Native provider, it used API version 2023-10-01-preview.
//
// Other available API versions: 2023-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dashboard [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type IntegrationFabric struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name       pulumi.StringOutput                       `pulumi:"name"`
	Properties IntegrationFabricPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIntegrationFabric registers a new resource with the given unique name, arguments, and options.
func NewIntegrationFabric(ctx *pulumi.Context,
	name string, args *IntegrationFabricArgs, opts ...pulumi.ResourceOption) (*IntegrationFabric, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dashboard/v20231001preview:IntegrationFabric"),
		},
		{
			Type: pulumi.String("azure-native:dashboard/v20241001:IntegrationFabric"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IntegrationFabric
	err := ctx.RegisterResource("azure-native:dashboard:IntegrationFabric", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationFabric gets an existing IntegrationFabric resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationFabric(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationFabricState, opts ...pulumi.ResourceOption) (*IntegrationFabric, error) {
	var resource IntegrationFabric
	err := ctx.ReadResource("azure-native:dashboard:IntegrationFabric", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationFabric resources.
type integrationFabricState struct {
}

type IntegrationFabricState struct {
}

func (IntegrationFabricState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationFabricState)(nil)).Elem()
}

type integrationFabricArgs struct {
	// The integration fabric name of Azure Managed Grafana.
	IntegrationFabricName *string `pulumi:"integrationFabricName"`
	// The geo-location where the resource lives
	Location   *string                      `pulumi:"location"`
	Properties *IntegrationFabricProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The workspace name of Azure Managed Grafana.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a IntegrationFabric resource.
type IntegrationFabricArgs struct {
	// The integration fabric name of Azure Managed Grafana.
	IntegrationFabricName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location   pulumi.StringPtrInput
	Properties IntegrationFabricPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The workspace name of Azure Managed Grafana.
	WorkspaceName pulumi.StringInput
}

func (IntegrationFabricArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationFabricArgs)(nil)).Elem()
}

type IntegrationFabricInput interface {
	pulumi.Input

	ToIntegrationFabricOutput() IntegrationFabricOutput
	ToIntegrationFabricOutputWithContext(ctx context.Context) IntegrationFabricOutput
}

func (*IntegrationFabric) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationFabric)(nil)).Elem()
}

func (i *IntegrationFabric) ToIntegrationFabricOutput() IntegrationFabricOutput {
	return i.ToIntegrationFabricOutputWithContext(context.Background())
}

func (i *IntegrationFabric) ToIntegrationFabricOutputWithContext(ctx context.Context) IntegrationFabricOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationFabricOutput)
}

type IntegrationFabricOutput struct{ *pulumi.OutputState }

func (IntegrationFabricOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationFabric)(nil)).Elem()
}

func (o IntegrationFabricOutput) ToIntegrationFabricOutput() IntegrationFabricOutput {
	return o
}

func (o IntegrationFabricOutput) ToIntegrationFabricOutputWithContext(ctx context.Context) IntegrationFabricOutput {
	return o
}

// The Azure API version of the resource.
func (o IntegrationFabricOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationFabric) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o IntegrationFabricOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationFabric) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o IntegrationFabricOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationFabric) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IntegrationFabricOutput) Properties() IntegrationFabricPropertiesResponseOutput {
	return o.ApplyT(func(v *IntegrationFabric) IntegrationFabricPropertiesResponseOutput { return v.Properties }).(IntegrationFabricPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o IntegrationFabricOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *IntegrationFabric) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o IntegrationFabricOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationFabric) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o IntegrationFabricOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationFabric) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationFabricOutput{})
}
