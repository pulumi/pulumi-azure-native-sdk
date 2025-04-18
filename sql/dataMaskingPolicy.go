// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A database data masking policy.
//
// Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2021-11-01.
//
// Other available API versions: 2014-04-01, 2021-11-01, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type DataMaskingPolicy struct {
	pulumi.CustomResourceState

	// The list of the application principals. This is a legacy parameter and is no longer used.
	ApplicationPrincipals pulumi.StringOutput `pulumi:"applicationPrincipals"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The state of the data masking policy.
	DataMaskingState pulumi.StringOutput `pulumi:"dataMaskingState"`
	// The list of the exempt principals. Specifies the semicolon-separated list of database users for which the data masking policy does not apply. The specified users receive data results without masking for all of the database queries.
	ExemptPrincipals pulumi.StringPtrOutput `pulumi:"exemptPrincipals"`
	// The kind of Data Masking Policy. Metadata, used for Azure portal.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The location of the data masking policy.
	Location pulumi.StringOutput `pulumi:"location"`
	// The masking level. This is a legacy parameter and is no longer used.
	MaskingLevel pulumi.StringOutput `pulumi:"maskingLevel"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataMaskingPolicy registers a new resource with the given unique name, arguments, and options.
func NewDataMaskingPolicy(ctx *pulumi.Context,
	name string, args *DataMaskingPolicyArgs, opts ...pulumi.ResourceOption) (*DataMaskingPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataMaskingState == nil {
		return nil, errors.New("invalid value for required argument 'DataMaskingState'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql/v20140401:DataMaskingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:DataMaskingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:DataMaskingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220501preview:DataMaskingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220801preview:DataMaskingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20221101preview:DataMaskingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230201preview:DataMaskingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230501preview:DataMaskingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230801:DataMaskingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230801preview:DataMaskingPolicy"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20240501preview:DataMaskingPolicy"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource DataMaskingPolicy
	err := ctx.RegisterResource("azure-native:sql:DataMaskingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataMaskingPolicy gets an existing DataMaskingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataMaskingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataMaskingPolicyState, opts ...pulumi.ResourceOption) (*DataMaskingPolicy, error) {
	var resource DataMaskingPolicy
	err := ctx.ReadResource("azure-native:sql:DataMaskingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataMaskingPolicy resources.
type dataMaskingPolicyState struct {
}

type DataMaskingPolicyState struct {
}

func (DataMaskingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataMaskingPolicyState)(nil)).Elem()
}

type dataMaskingPolicyArgs struct {
	// The name of the database for which the data masking policy applies.
	DataMaskingPolicyName *string `pulumi:"dataMaskingPolicyName"`
	// The state of the data masking policy.
	DataMaskingState DataMaskingState `pulumi:"dataMaskingState"`
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The list of the exempt principals. Specifies the semicolon-separated list of database users for which the data masking policy does not apply. The specified users receive data results without masking for all of the database queries.
	ExemptPrincipals *string `pulumi:"exemptPrincipals"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// The set of arguments for constructing a DataMaskingPolicy resource.
type DataMaskingPolicyArgs struct {
	// The name of the database for which the data masking policy applies.
	DataMaskingPolicyName pulumi.StringPtrInput
	// The state of the data masking policy.
	DataMaskingState DataMaskingStateInput
	// The name of the database.
	DatabaseName pulumi.StringInput
	// The list of the exempt principals. Specifies the semicolon-separated list of database users for which the data masking policy does not apply. The specified users receive data results without masking for all of the database queries.
	ExemptPrincipals pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
}

func (DataMaskingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataMaskingPolicyArgs)(nil)).Elem()
}

type DataMaskingPolicyInput interface {
	pulumi.Input

	ToDataMaskingPolicyOutput() DataMaskingPolicyOutput
	ToDataMaskingPolicyOutputWithContext(ctx context.Context) DataMaskingPolicyOutput
}

func (*DataMaskingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DataMaskingPolicy)(nil)).Elem()
}

func (i *DataMaskingPolicy) ToDataMaskingPolicyOutput() DataMaskingPolicyOutput {
	return i.ToDataMaskingPolicyOutputWithContext(context.Background())
}

func (i *DataMaskingPolicy) ToDataMaskingPolicyOutputWithContext(ctx context.Context) DataMaskingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataMaskingPolicyOutput)
}

type DataMaskingPolicyOutput struct{ *pulumi.OutputState }

func (DataMaskingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataMaskingPolicy)(nil)).Elem()
}

func (o DataMaskingPolicyOutput) ToDataMaskingPolicyOutput() DataMaskingPolicyOutput {
	return o
}

func (o DataMaskingPolicyOutput) ToDataMaskingPolicyOutputWithContext(ctx context.Context) DataMaskingPolicyOutput {
	return o
}

// The list of the application principals. This is a legacy parameter and is no longer used.
func (o DataMaskingPolicyOutput) ApplicationPrincipals() pulumi.StringOutput {
	return o.ApplyT(func(v *DataMaskingPolicy) pulumi.StringOutput { return v.ApplicationPrincipals }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o DataMaskingPolicyOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DataMaskingPolicy) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The state of the data masking policy.
func (o DataMaskingPolicyOutput) DataMaskingState() pulumi.StringOutput {
	return o.ApplyT(func(v *DataMaskingPolicy) pulumi.StringOutput { return v.DataMaskingState }).(pulumi.StringOutput)
}

// The list of the exempt principals. Specifies the semicolon-separated list of database users for which the data masking policy does not apply. The specified users receive data results without masking for all of the database queries.
func (o DataMaskingPolicyOutput) ExemptPrincipals() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataMaskingPolicy) pulumi.StringPtrOutput { return v.ExemptPrincipals }).(pulumi.StringPtrOutput)
}

// The kind of Data Masking Policy. Metadata, used for Azure portal.
func (o DataMaskingPolicyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *DataMaskingPolicy) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The location of the data masking policy.
func (o DataMaskingPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataMaskingPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The masking level. This is a legacy parameter and is no longer used.
func (o DataMaskingPolicyOutput) MaskingLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *DataMaskingPolicy) pulumi.StringOutput { return v.MaskingLevel }).(pulumi.StringOutput)
}

// Resource name.
func (o DataMaskingPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataMaskingPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource type.
func (o DataMaskingPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataMaskingPolicy) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataMaskingPolicyOutput{})
}
