// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing a read only following database.
//
// Uses Azure REST API version 2024-04-13. In version 2.x of the Azure Native provider, it used API version 2022-12-29.
type ReadOnlyFollowingDatabase struct {
	pulumi.CustomResourceState

	// The name of the attached database configuration cluster
	AttachedDatabaseConfigurationName pulumi.StringOutput `pulumi:"attachedDatabaseConfigurationName"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The origin of the following setup.
	DatabaseShareOrigin pulumi.StringOutput `pulumi:"databaseShareOrigin"`
	// The time the data should be kept in cache for fast queries in TimeSpan.
	HotCachePeriod pulumi.StringPtrOutput `pulumi:"hotCachePeriod"`
	// Kind of the database
	// Expected value is 'ReadOnlyFollowing'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the leader cluster
	LeaderClusterResourceId pulumi.StringOutput `pulumi:"leaderClusterResourceId"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The original database name, before databaseNameOverride or databaseNamePrefix where applied.
	OriginalDatabaseName pulumi.StringOutput `pulumi:"originalDatabaseName"`
	// The principals modification kind of the database
	PrincipalsModificationKind pulumi.StringOutput `pulumi:"principalsModificationKind"`
	// The provisioned state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The time the data should be kept before it stops being accessible to queries in TimeSpan.
	SoftDeletePeriod pulumi.StringOutput `pulumi:"softDeletePeriod"`
	// The statistics of the database.
	Statistics DatabaseStatisticsResponseOutput `pulumi:"statistics"`
	// The database suspension details. If the database is suspended, this object contains information related to the database's suspension state.
	SuspensionDetails SuspensionDetailsResponseOutput `pulumi:"suspensionDetails"`
	// Table level sharing specifications
	TableLevelSharingProperties TableLevelSharingPropertiesResponseOutput `pulumi:"tableLevelSharingProperties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReadOnlyFollowingDatabase registers a new resource with the given unique name, arguments, and options.
func NewReadOnlyFollowingDatabase(ctx *pulumi.Context,
	name string, args *ReadOnlyFollowingDatabaseArgs, opts ...pulumi.ResourceOption) (*ReadOnlyFollowingDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.Kind = pulumi.String("ReadOnlyFollowing")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:kusto/v20170907privatepreview:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20180907preview:Database"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20180907preview:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190121:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190515:Database"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190515:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190907:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20190907:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20191109:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200215:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200614:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20200918:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210101:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20210827:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220201:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20220707:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221111:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221229:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20221229:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230502:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230502:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230815:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20230815:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20240413:ReadOnlyFollowingDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto/v20240413:ReadWriteDatabase"),
		},
		{
			Type: pulumi.String("azure-native:kusto:ReadWriteDatabase"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ReadOnlyFollowingDatabase
	err := ctx.RegisterResource("azure-native:kusto:ReadOnlyFollowingDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReadOnlyFollowingDatabase gets an existing ReadOnlyFollowingDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReadOnlyFollowingDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReadOnlyFollowingDatabaseState, opts ...pulumi.ResourceOption) (*ReadOnlyFollowingDatabase, error) {
	var resource ReadOnlyFollowingDatabase
	err := ctx.ReadResource("azure-native:kusto:ReadOnlyFollowingDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReadOnlyFollowingDatabase resources.
type readOnlyFollowingDatabaseState struct {
}

type ReadOnlyFollowingDatabaseState struct {
}

func (ReadOnlyFollowingDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*readOnlyFollowingDatabaseState)(nil)).Elem()
}

type readOnlyFollowingDatabaseArgs struct {
	// By default, any user who run operation on a database become an Admin on it. This property allows the caller to exclude the caller from Admins list.
	CallerRole *string `pulumi:"callerRole"`
	// The name of the Kusto cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the database in the Kusto cluster.
	DatabaseName *string `pulumi:"databaseName"`
	// The time the data should be kept in cache for fast queries in TimeSpan.
	HotCachePeriod *string `pulumi:"hotCachePeriod"`
	// Kind of the database
	// Expected value is 'ReadOnlyFollowing'.
	Kind string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ReadOnlyFollowingDatabase resource.
type ReadOnlyFollowingDatabaseArgs struct {
	// By default, any user who run operation on a database become an Admin on it. This property allows the caller to exclude the caller from Admins list.
	CallerRole pulumi.StringPtrInput
	// The name of the Kusto cluster.
	ClusterName pulumi.StringInput
	// The name of the database in the Kusto cluster.
	DatabaseName pulumi.StringPtrInput
	// The time the data should be kept in cache for fast queries in TimeSpan.
	HotCachePeriod pulumi.StringPtrInput
	// Kind of the database
	// Expected value is 'ReadOnlyFollowing'.
	Kind pulumi.StringInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (ReadOnlyFollowingDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*readOnlyFollowingDatabaseArgs)(nil)).Elem()
}

type ReadOnlyFollowingDatabaseInput interface {
	pulumi.Input

	ToReadOnlyFollowingDatabaseOutput() ReadOnlyFollowingDatabaseOutput
	ToReadOnlyFollowingDatabaseOutputWithContext(ctx context.Context) ReadOnlyFollowingDatabaseOutput
}

func (*ReadOnlyFollowingDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadOnlyFollowingDatabase)(nil)).Elem()
}

func (i *ReadOnlyFollowingDatabase) ToReadOnlyFollowingDatabaseOutput() ReadOnlyFollowingDatabaseOutput {
	return i.ToReadOnlyFollowingDatabaseOutputWithContext(context.Background())
}

func (i *ReadOnlyFollowingDatabase) ToReadOnlyFollowingDatabaseOutputWithContext(ctx context.Context) ReadOnlyFollowingDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadOnlyFollowingDatabaseOutput)
}

type ReadOnlyFollowingDatabaseOutput struct{ *pulumi.OutputState }

func (ReadOnlyFollowingDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadOnlyFollowingDatabase)(nil)).Elem()
}

func (o ReadOnlyFollowingDatabaseOutput) ToReadOnlyFollowingDatabaseOutput() ReadOnlyFollowingDatabaseOutput {
	return o
}

func (o ReadOnlyFollowingDatabaseOutput) ToReadOnlyFollowingDatabaseOutputWithContext(ctx context.Context) ReadOnlyFollowingDatabaseOutput {
	return o
}

// The name of the attached database configuration cluster
func (o ReadOnlyFollowingDatabaseOutput) AttachedDatabaseConfigurationName() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.AttachedDatabaseConfigurationName }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o ReadOnlyFollowingDatabaseOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The origin of the following setup.
func (o ReadOnlyFollowingDatabaseOutput) DatabaseShareOrigin() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.DatabaseShareOrigin }).(pulumi.StringOutput)
}

// The time the data should be kept in cache for fast queries in TimeSpan.
func (o ReadOnlyFollowingDatabaseOutput) HotCachePeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringPtrOutput { return v.HotCachePeriod }).(pulumi.StringPtrOutput)
}

// Kind of the database
// Expected value is 'ReadOnlyFollowing'.
func (o ReadOnlyFollowingDatabaseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the leader cluster
func (o ReadOnlyFollowingDatabaseOutput) LeaderClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.LeaderClusterResourceId }).(pulumi.StringOutput)
}

// Resource location.
func (o ReadOnlyFollowingDatabaseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o ReadOnlyFollowingDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The original database name, before databaseNameOverride or databaseNamePrefix where applied.
func (o ReadOnlyFollowingDatabaseOutput) OriginalDatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.OriginalDatabaseName }).(pulumi.StringOutput)
}

// The principals modification kind of the database
func (o ReadOnlyFollowingDatabaseOutput) PrincipalsModificationKind() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.PrincipalsModificationKind }).(pulumi.StringOutput)
}

// The provisioned state of the resource.
func (o ReadOnlyFollowingDatabaseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The time the data should be kept before it stops being accessible to queries in TimeSpan.
func (o ReadOnlyFollowingDatabaseOutput) SoftDeletePeriod() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.SoftDeletePeriod }).(pulumi.StringOutput)
}

// The statistics of the database.
func (o ReadOnlyFollowingDatabaseOutput) Statistics() DatabaseStatisticsResponseOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) DatabaseStatisticsResponseOutput { return v.Statistics }).(DatabaseStatisticsResponseOutput)
}

// The database suspension details. If the database is suspended, this object contains information related to the database's suspension state.
func (o ReadOnlyFollowingDatabaseOutput) SuspensionDetails() SuspensionDetailsResponseOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) SuspensionDetailsResponseOutput { return v.SuspensionDetails }).(SuspensionDetailsResponseOutput)
}

// Table level sharing specifications
func (o ReadOnlyFollowingDatabaseOutput) TableLevelSharingProperties() TableLevelSharingPropertiesResponseOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) TableLevelSharingPropertiesResponseOutput {
		return v.TableLevelSharingProperties
	}).(TableLevelSharingPropertiesResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ReadOnlyFollowingDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadOnlyFollowingDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReadOnlyFollowingDatabaseOutput{})
}
