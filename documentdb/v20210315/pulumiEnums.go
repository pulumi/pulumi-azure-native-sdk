// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210315

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Describes the mode of backups.
type BackupPolicyType string

const (
	BackupPolicyTypePeriodic   = BackupPolicyType("Periodic")
	BackupPolicyTypeContinuous = BackupPolicyType("Continuous")
)

// Sort order for composite paths.
type CompositePathSortOrder string

const (
	CompositePathSortOrderAscending  = CompositePathSortOrder("ascending")
	CompositePathSortOrderDescending = CompositePathSortOrder("descending")
)

// Indicates the conflict resolution mode.
type ConflictResolutionMode string

const (
	ConflictResolutionModeLastWriterWins = ConflictResolutionMode("LastWriterWins")
	ConflictResolutionModeCustom         = ConflictResolutionMode("Custom")
)

// The cassandra connector offer type for the Cosmos DB database C* account.
type ConnectorOffer string

const (
	ConnectorOfferSmall = ConnectorOffer("Small")
)

// The datatype for which the indexing behavior is applied to.
type DataType string

const (
	DataTypeString       = DataType("String")
	DataTypeNumber       = DataType("Number")
	DataTypePoint        = DataType("Point")
	DataTypePolygon      = DataType("Polygon")
	DataTypeLineString   = DataType("LineString")
	DataTypeMultiPolygon = DataType("MultiPolygon")
)

// Indicates the type of database account. This can only be set at database account creation.
type DatabaseAccountKind string

const (
	DatabaseAccountKindGlobalDocumentDB = DatabaseAccountKind("GlobalDocumentDB")
	DatabaseAccountKindMongoDB          = DatabaseAccountKind("MongoDB")
	DatabaseAccountKindParse            = DatabaseAccountKind("Parse")
)

// The offer type for the database
type DatabaseAccountOfferType string

const (
	DatabaseAccountOfferTypeStandard = DatabaseAccountOfferType("Standard")
)

func (DatabaseAccountOfferType) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseAccountOfferType)(nil)).Elem()
}

func (e DatabaseAccountOfferType) ToDatabaseAccountOfferTypeOutput() DatabaseAccountOfferTypeOutput {
	return pulumi.ToOutput(e).(DatabaseAccountOfferTypeOutput)
}

func (e DatabaseAccountOfferType) ToDatabaseAccountOfferTypeOutputWithContext(ctx context.Context) DatabaseAccountOfferTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatabaseAccountOfferTypeOutput)
}

func (e DatabaseAccountOfferType) ToDatabaseAccountOfferTypePtrOutput() DatabaseAccountOfferTypePtrOutput {
	return e.ToDatabaseAccountOfferTypePtrOutputWithContext(context.Background())
}

func (e DatabaseAccountOfferType) ToDatabaseAccountOfferTypePtrOutputWithContext(ctx context.Context) DatabaseAccountOfferTypePtrOutput {
	return DatabaseAccountOfferType(e).ToDatabaseAccountOfferTypeOutputWithContext(ctx).ToDatabaseAccountOfferTypePtrOutputWithContext(ctx)
}

func (e DatabaseAccountOfferType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseAccountOfferType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabaseAccountOfferType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatabaseAccountOfferType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatabaseAccountOfferTypeOutput struct{ *pulumi.OutputState }

func (DatabaseAccountOfferTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseAccountOfferType)(nil)).Elem()
}

func (o DatabaseAccountOfferTypeOutput) ToDatabaseAccountOfferTypeOutput() DatabaseAccountOfferTypeOutput {
	return o
}

func (o DatabaseAccountOfferTypeOutput) ToDatabaseAccountOfferTypeOutputWithContext(ctx context.Context) DatabaseAccountOfferTypeOutput {
	return o
}

func (o DatabaseAccountOfferTypeOutput) ToDatabaseAccountOfferTypePtrOutput() DatabaseAccountOfferTypePtrOutput {
	return o.ToDatabaseAccountOfferTypePtrOutputWithContext(context.Background())
}

func (o DatabaseAccountOfferTypeOutput) ToDatabaseAccountOfferTypePtrOutputWithContext(ctx context.Context) DatabaseAccountOfferTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabaseAccountOfferType) *DatabaseAccountOfferType {
		return &v
	}).(DatabaseAccountOfferTypePtrOutput)
}

func (o DatabaseAccountOfferTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatabaseAccountOfferTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseAccountOfferType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatabaseAccountOfferTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseAccountOfferTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabaseAccountOfferType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatabaseAccountOfferTypePtrOutput struct{ *pulumi.OutputState }

func (DatabaseAccountOfferTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseAccountOfferType)(nil)).Elem()
}

func (o DatabaseAccountOfferTypePtrOutput) ToDatabaseAccountOfferTypePtrOutput() DatabaseAccountOfferTypePtrOutput {
	return o
}

func (o DatabaseAccountOfferTypePtrOutput) ToDatabaseAccountOfferTypePtrOutputWithContext(ctx context.Context) DatabaseAccountOfferTypePtrOutput {
	return o
}

func (o DatabaseAccountOfferTypePtrOutput) Elem() DatabaseAccountOfferTypeOutput {
	return o.ApplyT(func(v *DatabaseAccountOfferType) DatabaseAccountOfferType {
		if v != nil {
			return *v
		}
		var ret DatabaseAccountOfferType
		return ret
	}).(DatabaseAccountOfferTypeOutput)
}

func (o DatabaseAccountOfferTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabaseAccountOfferTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatabaseAccountOfferType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DatabaseAccountOfferTypeInput is an input type that accepts DatabaseAccountOfferTypeArgs and DatabaseAccountOfferTypeOutput values.
// You can construct a concrete instance of `DatabaseAccountOfferTypeInput` via:
//
//	DatabaseAccountOfferTypeArgs{...}
type DatabaseAccountOfferTypeInput interface {
	pulumi.Input

	ToDatabaseAccountOfferTypeOutput() DatabaseAccountOfferTypeOutput
	ToDatabaseAccountOfferTypeOutputWithContext(context.Context) DatabaseAccountOfferTypeOutput
}

var databaseAccountOfferTypePtrType = reflect.TypeOf((**DatabaseAccountOfferType)(nil)).Elem()

type DatabaseAccountOfferTypePtrInput interface {
	pulumi.Input

	ToDatabaseAccountOfferTypePtrOutput() DatabaseAccountOfferTypePtrOutput
	ToDatabaseAccountOfferTypePtrOutputWithContext(context.Context) DatabaseAccountOfferTypePtrOutput
}

type databaseAccountOfferTypePtr string

func DatabaseAccountOfferTypePtr(v string) DatabaseAccountOfferTypePtrInput {
	return (*databaseAccountOfferTypePtr)(&v)
}

func (*databaseAccountOfferTypePtr) ElementType() reflect.Type {
	return databaseAccountOfferTypePtrType
}

func (in *databaseAccountOfferTypePtr) ToDatabaseAccountOfferTypePtrOutput() DatabaseAccountOfferTypePtrOutput {
	return pulumi.ToOutput(in).(DatabaseAccountOfferTypePtrOutput)
}

func (in *databaseAccountOfferTypePtr) ToDatabaseAccountOfferTypePtrOutputWithContext(ctx context.Context) DatabaseAccountOfferTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatabaseAccountOfferTypePtrOutput)
}

// The default consistency level and configuration settings of the Cosmos DB account.
type DefaultConsistencyLevel string

const (
	DefaultConsistencyLevelEventual         = DefaultConsistencyLevel("Eventual")
	DefaultConsistencyLevelSession          = DefaultConsistencyLevel("Session")
	DefaultConsistencyLevelBoundedStaleness = DefaultConsistencyLevel("BoundedStaleness")
	DefaultConsistencyLevelStrong           = DefaultConsistencyLevel("Strong")
	DefaultConsistencyLevelConsistentPrefix = DefaultConsistencyLevel("ConsistentPrefix")
)

func (DefaultConsistencyLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultConsistencyLevel)(nil)).Elem()
}

func (e DefaultConsistencyLevel) ToDefaultConsistencyLevelOutput() DefaultConsistencyLevelOutput {
	return pulumi.ToOutput(e).(DefaultConsistencyLevelOutput)
}

func (e DefaultConsistencyLevel) ToDefaultConsistencyLevelOutputWithContext(ctx context.Context) DefaultConsistencyLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DefaultConsistencyLevelOutput)
}

func (e DefaultConsistencyLevel) ToDefaultConsistencyLevelPtrOutput() DefaultConsistencyLevelPtrOutput {
	return e.ToDefaultConsistencyLevelPtrOutputWithContext(context.Background())
}

func (e DefaultConsistencyLevel) ToDefaultConsistencyLevelPtrOutputWithContext(ctx context.Context) DefaultConsistencyLevelPtrOutput {
	return DefaultConsistencyLevel(e).ToDefaultConsistencyLevelOutputWithContext(ctx).ToDefaultConsistencyLevelPtrOutputWithContext(ctx)
}

func (e DefaultConsistencyLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultConsistencyLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultConsistencyLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DefaultConsistencyLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DefaultConsistencyLevelOutput struct{ *pulumi.OutputState }

func (DefaultConsistencyLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultConsistencyLevel)(nil)).Elem()
}

func (o DefaultConsistencyLevelOutput) ToDefaultConsistencyLevelOutput() DefaultConsistencyLevelOutput {
	return o
}

func (o DefaultConsistencyLevelOutput) ToDefaultConsistencyLevelOutputWithContext(ctx context.Context) DefaultConsistencyLevelOutput {
	return o
}

func (o DefaultConsistencyLevelOutput) ToDefaultConsistencyLevelPtrOutput() DefaultConsistencyLevelPtrOutput {
	return o.ToDefaultConsistencyLevelPtrOutputWithContext(context.Background())
}

func (o DefaultConsistencyLevelOutput) ToDefaultConsistencyLevelPtrOutputWithContext(ctx context.Context) DefaultConsistencyLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultConsistencyLevel) *DefaultConsistencyLevel {
		return &v
	}).(DefaultConsistencyLevelPtrOutput)
}

func (o DefaultConsistencyLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DefaultConsistencyLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultConsistencyLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DefaultConsistencyLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultConsistencyLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultConsistencyLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DefaultConsistencyLevelPtrOutput struct{ *pulumi.OutputState }

func (DefaultConsistencyLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultConsistencyLevel)(nil)).Elem()
}

func (o DefaultConsistencyLevelPtrOutput) ToDefaultConsistencyLevelPtrOutput() DefaultConsistencyLevelPtrOutput {
	return o
}

func (o DefaultConsistencyLevelPtrOutput) ToDefaultConsistencyLevelPtrOutputWithContext(ctx context.Context) DefaultConsistencyLevelPtrOutput {
	return o
}

func (o DefaultConsistencyLevelPtrOutput) Elem() DefaultConsistencyLevelOutput {
	return o.ApplyT(func(v *DefaultConsistencyLevel) DefaultConsistencyLevel {
		if v != nil {
			return *v
		}
		var ret DefaultConsistencyLevel
		return ret
	}).(DefaultConsistencyLevelOutput)
}

func (o DefaultConsistencyLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultConsistencyLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DefaultConsistencyLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DefaultConsistencyLevelInput is an input type that accepts DefaultConsistencyLevelArgs and DefaultConsistencyLevelOutput values.
// You can construct a concrete instance of `DefaultConsistencyLevelInput` via:
//
//	DefaultConsistencyLevelArgs{...}
type DefaultConsistencyLevelInput interface {
	pulumi.Input

	ToDefaultConsistencyLevelOutput() DefaultConsistencyLevelOutput
	ToDefaultConsistencyLevelOutputWithContext(context.Context) DefaultConsistencyLevelOutput
}

var defaultConsistencyLevelPtrType = reflect.TypeOf((**DefaultConsistencyLevel)(nil)).Elem()

type DefaultConsistencyLevelPtrInput interface {
	pulumi.Input

	ToDefaultConsistencyLevelPtrOutput() DefaultConsistencyLevelPtrOutput
	ToDefaultConsistencyLevelPtrOutputWithContext(context.Context) DefaultConsistencyLevelPtrOutput
}

type defaultConsistencyLevelPtr string

func DefaultConsistencyLevelPtr(v string) DefaultConsistencyLevelPtrInput {
	return (*defaultConsistencyLevelPtr)(&v)
}

func (*defaultConsistencyLevelPtr) ElementType() reflect.Type {
	return defaultConsistencyLevelPtrType
}

func (in *defaultConsistencyLevelPtr) ToDefaultConsistencyLevelPtrOutput() DefaultConsistencyLevelPtrOutput {
	return pulumi.ToOutput(in).(DefaultConsistencyLevelPtrOutput)
}

func (in *defaultConsistencyLevelPtr) ToDefaultConsistencyLevelPtrOutputWithContext(ctx context.Context) DefaultConsistencyLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DefaultConsistencyLevelPtrOutput)
}

// Indicates the type of index.
type IndexKind string

const (
	IndexKindHash    = IndexKind("Hash")
	IndexKindRange   = IndexKind("Range")
	IndexKindSpatial = IndexKind("Spatial")
)

// Indicates the indexing mode.
type IndexingMode string

const (
	IndexingModeConsistent = IndexingMode("consistent")
	IndexingModeLazy       = IndexingMode("lazy")
	IndexingModeNone       = IndexingMode("none")
)

// Indicates what services are allowed to bypass firewall checks.
type NetworkAclBypass string

const (
	NetworkAclBypassNone          = NetworkAclBypass("None")
	NetworkAclBypassAzureServices = NetworkAclBypass("AzureServices")
)

func (NetworkAclBypass) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAclBypass)(nil)).Elem()
}

func (e NetworkAclBypass) ToNetworkAclBypassOutput() NetworkAclBypassOutput {
	return pulumi.ToOutput(e).(NetworkAclBypassOutput)
}

func (e NetworkAclBypass) ToNetworkAclBypassOutputWithContext(ctx context.Context) NetworkAclBypassOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkAclBypassOutput)
}

func (e NetworkAclBypass) ToNetworkAclBypassPtrOutput() NetworkAclBypassPtrOutput {
	return e.ToNetworkAclBypassPtrOutputWithContext(context.Background())
}

func (e NetworkAclBypass) ToNetworkAclBypassPtrOutputWithContext(ctx context.Context) NetworkAclBypassPtrOutput {
	return NetworkAclBypass(e).ToNetworkAclBypassOutputWithContext(ctx).ToNetworkAclBypassPtrOutputWithContext(ctx)
}

func (e NetworkAclBypass) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkAclBypass) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkAclBypass) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkAclBypass) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkAclBypassOutput struct{ *pulumi.OutputState }

func (NetworkAclBypassOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAclBypass)(nil)).Elem()
}

func (o NetworkAclBypassOutput) ToNetworkAclBypassOutput() NetworkAclBypassOutput {
	return o
}

func (o NetworkAclBypassOutput) ToNetworkAclBypassOutputWithContext(ctx context.Context) NetworkAclBypassOutput {
	return o
}

func (o NetworkAclBypassOutput) ToNetworkAclBypassPtrOutput() NetworkAclBypassPtrOutput {
	return o.ToNetworkAclBypassPtrOutputWithContext(context.Background())
}

func (o NetworkAclBypassOutput) ToNetworkAclBypassPtrOutputWithContext(ctx context.Context) NetworkAclBypassPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkAclBypass) *NetworkAclBypass {
		return &v
	}).(NetworkAclBypassPtrOutput)
}

func (o NetworkAclBypassOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkAclBypassOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkAclBypass) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkAclBypassOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkAclBypassOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkAclBypass) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkAclBypassPtrOutput struct{ *pulumi.OutputState }

func (NetworkAclBypassPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAclBypass)(nil)).Elem()
}

func (o NetworkAclBypassPtrOutput) ToNetworkAclBypassPtrOutput() NetworkAclBypassPtrOutput {
	return o
}

func (o NetworkAclBypassPtrOutput) ToNetworkAclBypassPtrOutputWithContext(ctx context.Context) NetworkAclBypassPtrOutput {
	return o
}

func (o NetworkAclBypassPtrOutput) Elem() NetworkAclBypassOutput {
	return o.ApplyT(func(v *NetworkAclBypass) NetworkAclBypass {
		if v != nil {
			return *v
		}
		var ret NetworkAclBypass
		return ret
	}).(NetworkAclBypassOutput)
}

func (o NetworkAclBypassPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkAclBypassPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkAclBypass) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// NetworkAclBypassInput is an input type that accepts NetworkAclBypassArgs and NetworkAclBypassOutput values.
// You can construct a concrete instance of `NetworkAclBypassInput` via:
//
//	NetworkAclBypassArgs{...}
type NetworkAclBypassInput interface {
	pulumi.Input

	ToNetworkAclBypassOutput() NetworkAclBypassOutput
	ToNetworkAclBypassOutputWithContext(context.Context) NetworkAclBypassOutput
}

var networkAclBypassPtrType = reflect.TypeOf((**NetworkAclBypass)(nil)).Elem()

type NetworkAclBypassPtrInput interface {
	pulumi.Input

	ToNetworkAclBypassPtrOutput() NetworkAclBypassPtrOutput
	ToNetworkAclBypassPtrOutputWithContext(context.Context) NetworkAclBypassPtrOutput
}

type networkAclBypassPtr string

func NetworkAclBypassPtr(v string) NetworkAclBypassPtrInput {
	return (*networkAclBypassPtr)(&v)
}

func (*networkAclBypassPtr) ElementType() reflect.Type {
	return networkAclBypassPtrType
}

func (in *networkAclBypassPtr) ToNetworkAclBypassPtrOutput() NetworkAclBypassPtrOutput {
	return pulumi.ToOutput(in).(NetworkAclBypassPtrOutput)
}

func (in *networkAclBypassPtr) ToNetworkAclBypassPtrOutputWithContext(ctx context.Context) NetworkAclBypassPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkAclBypassPtrOutput)
}

// Indicates the kind of algorithm used for partitioning. For MultiHash, multiple partition keys (upto three maximum) are supported for container create
type PartitionKind string

const (
	PartitionKindHash      = PartitionKind("Hash")
	PartitionKindRange     = PartitionKind("Range")
	PartitionKindMultiHash = PartitionKind("MultiHash")
)

// Whether requests from Public Network are allowed
type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

// The type of identity used for the resource. The type 'SystemAssigned,UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the service.
type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned,UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ResourceIdentityTypeInput is an input type that accepts ResourceIdentityTypeArgs and ResourceIdentityTypeOutput values.
// You can construct a concrete instance of `ResourceIdentityTypeInput` via:
//
//	ResourceIdentityTypeArgs{...}
type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

// Describes the ServerVersion of an a MongoDB account.
type ServerVersion string

const (
	ServerVersion_3_2 = ServerVersion("3.2")
	ServerVersion_3_6 = ServerVersion("3.6")
	ServerVersion_4_0 = ServerVersion("4.0")
)

// Indicates the spatial type of index.
type SpatialType string

const (
	SpatialTypePoint        = SpatialType("Point")
	SpatialTypeLineString   = SpatialType("LineString")
	SpatialTypePolygon      = SpatialType("Polygon")
	SpatialTypeMultiPolygon = SpatialType("MultiPolygon")
)

// The operation the trigger is associated with
type TriggerOperation string

const (
	TriggerOperationAll     = TriggerOperation("All")
	TriggerOperationCreate  = TriggerOperation("Create")
	TriggerOperationUpdate  = TriggerOperation("Update")
	TriggerOperationDelete  = TriggerOperation("Delete")
	TriggerOperationReplace = TriggerOperation("Replace")
)

// Type of the Trigger
type TriggerType string

const (
	TriggerTypePre  = TriggerType("Pre")
	TriggerTypePost = TriggerType("Post")
)

func init() {
	pulumi.RegisterOutputType(DatabaseAccountOfferTypeOutput{})
	pulumi.RegisterOutputType(DatabaseAccountOfferTypePtrOutput{})
	pulumi.RegisterOutputType(DefaultConsistencyLevelOutput{})
	pulumi.RegisterOutputType(DefaultConsistencyLevelPtrOutput{})
	pulumi.RegisterOutputType(NetworkAclBypassOutput{})
	pulumi.RegisterOutputType(NetworkAclBypassPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}