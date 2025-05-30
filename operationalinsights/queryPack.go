// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package operationalinsights

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Log Analytics QueryPack definition.
//
// Uses Azure REST API version 2023-09-01. In version 2.x of the Azure Native provider, it used API version 2019-09-01.
//
// Other available API versions: 2019-09-01, 2019-09-01-preview, 2025-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native operationalinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type QueryPack struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Current state of this QueryPack: whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The unique ID of your application. This field cannot be changed.
	QueryPackId pulumi.StringOutput `pulumi:"queryPackId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Creation Date for the Log Analytics QueryPack, in ISO 8601 format.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// Last modified date of the Log Analytics QueryPack, in ISO 8601 format.
	TimeModified pulumi.StringOutput `pulumi:"timeModified"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
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
		{
			Type: pulumi.String("azure-native:operationalinsights/v20230901:QueryPack"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20250201:QueryPack"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
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
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the Log Analytics QueryPack resource.
	QueryPackName *string `pulumi:"queryPackName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a QueryPack resource.
type QueryPackArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the Log Analytics QueryPack resource.
	QueryPackName pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
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

// The Azure API version of the resource.
func (o QueryPackOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o QueryPackOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
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

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o QueryPackOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *QueryPack) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
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

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o QueryPackOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *QueryPack) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(QueryPackOutput{})
}
