// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The license type to apply for this managed instance.
type ArcSqlManagedInstanceLicenseType string

const (
	ArcSqlManagedInstanceLicenseTypeBasePrice        = ArcSqlManagedInstanceLicenseType("BasePrice")
	ArcSqlManagedInstanceLicenseTypeLicenseIncluded  = ArcSqlManagedInstanceLicenseType("LicenseIncluded")
	ArcSqlManagedInstanceLicenseTypeDisasterRecovery = ArcSqlManagedInstanceLicenseType("DisasterRecovery")
)

// SQL Server license type.
type ArcSqlServerLicenseType string

const (
	ArcSqlServerLicenseTypePaid      = ArcSqlServerLicenseType("Paid")
	ArcSqlServerLicenseTypeFree      = ArcSqlServerLicenseType("Free")
	ArcSqlServerLicenseTypeHADR      = ArcSqlServerLicenseType("HADR")
	ArcSqlServerLicenseTypeUndefined = ArcSqlServerLicenseType("Undefined")
)

// The cloud connectivity status.
type ConnectionStatus string

const (
	ConnectionStatusConnected    = ConnectionStatus("Connected")
	ConnectionStatusDisconnected = ConnectionStatus("Disconnected")
	ConnectionStatusUnknown      = ConnectionStatus("Unknown")
)

// Status of Azure Defender.
type DefenderStatus string

const (
	DefenderStatusProtected   = DefenderStatus("Protected")
	DefenderStatusUnprotected = DefenderStatus("Unprotected")
	DefenderStatusUnknown     = DefenderStatus("Unknown")
)

// SQL Server edition.
type EditionType string

const (
	EditionTypeEvaluation = EditionType("Evaluation")
	EditionTypeEnterprise = EditionType("Enterprise")
	EditionTypeStandard   = EditionType("Standard")
	EditionTypeWeb        = EditionType("Web")
	EditionTypeDeveloper  = EditionType("Developer")
	EditionTypeExpress    = EditionType("Express")
)

// The type of the extended location.
type ExtendedLocationTypes string

const (
	ExtendedLocationTypesCustomLocation = ExtendedLocationTypes("CustomLocation")
)

// The infrastructure the data controller is running on.
type Infrastructure string

const (
	InfrastructureAzure      = Infrastructure("azure")
	InfrastructureGcp        = Infrastructure("gcp")
	InfrastructureAws        = Infrastructure("aws")
	InfrastructureAlibaba    = Infrastructure("alibaba")
	InfrastructureOnpremises = Infrastructure("onpremises")
	InfrastructureOther      = Infrastructure("other")
)

func (Infrastructure) ElementType() reflect.Type {
	return reflect.TypeOf((*Infrastructure)(nil)).Elem()
}

func (e Infrastructure) ToInfrastructureOutput() InfrastructureOutput {
	return pulumi.ToOutput(e).(InfrastructureOutput)
}

func (e Infrastructure) ToInfrastructureOutputWithContext(ctx context.Context) InfrastructureOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InfrastructureOutput)
}

func (e Infrastructure) ToInfrastructurePtrOutput() InfrastructurePtrOutput {
	return e.ToInfrastructurePtrOutputWithContext(context.Background())
}

func (e Infrastructure) ToInfrastructurePtrOutputWithContext(ctx context.Context) InfrastructurePtrOutput {
	return Infrastructure(e).ToInfrastructureOutputWithContext(ctx).ToInfrastructurePtrOutputWithContext(ctx)
}

func (e Infrastructure) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Infrastructure) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Infrastructure) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Infrastructure) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InfrastructureOutput struct{ *pulumi.OutputState }

func (InfrastructureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Infrastructure)(nil)).Elem()
}

func (o InfrastructureOutput) ToInfrastructureOutput() InfrastructureOutput {
	return o
}

func (o InfrastructureOutput) ToInfrastructureOutputWithContext(ctx context.Context) InfrastructureOutput {
	return o
}

func (o InfrastructureOutput) ToInfrastructurePtrOutput() InfrastructurePtrOutput {
	return o.ToInfrastructurePtrOutputWithContext(context.Background())
}

func (o InfrastructureOutput) ToInfrastructurePtrOutputWithContext(ctx context.Context) InfrastructurePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Infrastructure) *Infrastructure {
		return &v
	}).(InfrastructurePtrOutput)
}

func (o InfrastructureOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InfrastructureOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Infrastructure) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InfrastructureOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InfrastructureOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Infrastructure) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InfrastructurePtrOutput struct{ *pulumi.OutputState }

func (InfrastructurePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Infrastructure)(nil)).Elem()
}

func (o InfrastructurePtrOutput) ToInfrastructurePtrOutput() InfrastructurePtrOutput {
	return o
}

func (o InfrastructurePtrOutput) ToInfrastructurePtrOutputWithContext(ctx context.Context) InfrastructurePtrOutput {
	return o
}

func (o InfrastructurePtrOutput) Elem() InfrastructureOutput {
	return o.ApplyT(func(v *Infrastructure) Infrastructure {
		if v != nil {
			return *v
		}
		var ret Infrastructure
		return ret
	}).(InfrastructureOutput)
}

func (o InfrastructurePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InfrastructurePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Infrastructure) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// InfrastructureInput is an input type that accepts InfrastructureArgs and InfrastructureOutput values.
// You can construct a concrete instance of `InfrastructureInput` via:
//
//	InfrastructureArgs{...}
type InfrastructureInput interface {
	pulumi.Input

	ToInfrastructureOutput() InfrastructureOutput
	ToInfrastructureOutputWithContext(context.Context) InfrastructureOutput
}

var infrastructurePtrType = reflect.TypeOf((**Infrastructure)(nil)).Elem()

type InfrastructurePtrInput interface {
	pulumi.Input

	ToInfrastructurePtrOutput() InfrastructurePtrOutput
	ToInfrastructurePtrOutputWithContext(context.Context) InfrastructurePtrOutput
}

type infrastructurePtr string

func InfrastructurePtr(v string) InfrastructurePtrInput {
	return (*infrastructurePtr)(&v)
}

func (*infrastructurePtr) ElementType() reflect.Type {
	return infrastructurePtrType
}

func (in *infrastructurePtr) ToInfrastructurePtrOutput() InfrastructurePtrOutput {
	return pulumi.ToOutput(in).(InfrastructurePtrOutput)
}

func (in *infrastructurePtr) ToInfrastructurePtrOutputWithContext(ctx context.Context) InfrastructurePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InfrastructurePtrOutput)
}

// The pricing tier for the instance.
type SqlManagedInstanceSkuTier string

const (
	SqlManagedInstanceSkuTierGeneralPurpose   = SqlManagedInstanceSkuTier("GeneralPurpose")
	SqlManagedInstanceSkuTierBusinessCritical = SqlManagedInstanceSkuTier("BusinessCritical")
)

func (SqlManagedInstanceSkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceSkuTier)(nil)).Elem()
}

func (e SqlManagedInstanceSkuTier) ToSqlManagedInstanceSkuTierOutput() SqlManagedInstanceSkuTierOutput {
	return pulumi.ToOutput(e).(SqlManagedInstanceSkuTierOutput)
}

func (e SqlManagedInstanceSkuTier) ToSqlManagedInstanceSkuTierOutputWithContext(ctx context.Context) SqlManagedInstanceSkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SqlManagedInstanceSkuTierOutput)
}

func (e SqlManagedInstanceSkuTier) ToSqlManagedInstanceSkuTierPtrOutput() SqlManagedInstanceSkuTierPtrOutput {
	return e.ToSqlManagedInstanceSkuTierPtrOutputWithContext(context.Background())
}

func (e SqlManagedInstanceSkuTier) ToSqlManagedInstanceSkuTierPtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuTierPtrOutput {
	return SqlManagedInstanceSkuTier(e).ToSqlManagedInstanceSkuTierOutputWithContext(ctx).ToSqlManagedInstanceSkuTierPtrOutputWithContext(ctx)
}

func (e SqlManagedInstanceSkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SqlManagedInstanceSkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SqlManagedInstanceSkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SqlManagedInstanceSkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SqlManagedInstanceSkuTierOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceSkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlManagedInstanceSkuTier)(nil)).Elem()
}

func (o SqlManagedInstanceSkuTierOutput) ToSqlManagedInstanceSkuTierOutput() SqlManagedInstanceSkuTierOutput {
	return o
}

func (o SqlManagedInstanceSkuTierOutput) ToSqlManagedInstanceSkuTierOutputWithContext(ctx context.Context) SqlManagedInstanceSkuTierOutput {
	return o
}

func (o SqlManagedInstanceSkuTierOutput) ToSqlManagedInstanceSkuTierPtrOutput() SqlManagedInstanceSkuTierPtrOutput {
	return o.ToSqlManagedInstanceSkuTierPtrOutputWithContext(context.Background())
}

func (o SqlManagedInstanceSkuTierOutput) ToSqlManagedInstanceSkuTierPtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlManagedInstanceSkuTier) *SqlManagedInstanceSkuTier {
		return &v
	}).(SqlManagedInstanceSkuTierPtrOutput)
}

func (o SqlManagedInstanceSkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SqlManagedInstanceSkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SqlManagedInstanceSkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SqlManagedInstanceSkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SqlManagedInstanceSkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SqlManagedInstanceSkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SqlManagedInstanceSkuTierPtrOutput struct{ *pulumi.OutputState }

func (SqlManagedInstanceSkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlManagedInstanceSkuTier)(nil)).Elem()
}

func (o SqlManagedInstanceSkuTierPtrOutput) ToSqlManagedInstanceSkuTierPtrOutput() SqlManagedInstanceSkuTierPtrOutput {
	return o
}

func (o SqlManagedInstanceSkuTierPtrOutput) ToSqlManagedInstanceSkuTierPtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuTierPtrOutput {
	return o
}

func (o SqlManagedInstanceSkuTierPtrOutput) Elem() SqlManagedInstanceSkuTierOutput {
	return o.ApplyT(func(v *SqlManagedInstanceSkuTier) SqlManagedInstanceSkuTier {
		if v != nil {
			return *v
		}
		var ret SqlManagedInstanceSkuTier
		return ret
	}).(SqlManagedInstanceSkuTierOutput)
}

func (o SqlManagedInstanceSkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SqlManagedInstanceSkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SqlManagedInstanceSkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SqlManagedInstanceSkuTierInput is an input type that accepts SqlManagedInstanceSkuTierArgs and SqlManagedInstanceSkuTierOutput values.
// You can construct a concrete instance of `SqlManagedInstanceSkuTierInput` via:
//
//	SqlManagedInstanceSkuTierArgs{...}
type SqlManagedInstanceSkuTierInput interface {
	pulumi.Input

	ToSqlManagedInstanceSkuTierOutput() SqlManagedInstanceSkuTierOutput
	ToSqlManagedInstanceSkuTierOutputWithContext(context.Context) SqlManagedInstanceSkuTierOutput
}

var sqlManagedInstanceSkuTierPtrType = reflect.TypeOf((**SqlManagedInstanceSkuTier)(nil)).Elem()

type SqlManagedInstanceSkuTierPtrInput interface {
	pulumi.Input

	ToSqlManagedInstanceSkuTierPtrOutput() SqlManagedInstanceSkuTierPtrOutput
	ToSqlManagedInstanceSkuTierPtrOutputWithContext(context.Context) SqlManagedInstanceSkuTierPtrOutput
}

type sqlManagedInstanceSkuTierPtr string

func SqlManagedInstanceSkuTierPtr(v string) SqlManagedInstanceSkuTierPtrInput {
	return (*sqlManagedInstanceSkuTierPtr)(&v)
}

func (*sqlManagedInstanceSkuTierPtr) ElementType() reflect.Type {
	return sqlManagedInstanceSkuTierPtrType
}

func (in *sqlManagedInstanceSkuTierPtr) ToSqlManagedInstanceSkuTierPtrOutput() SqlManagedInstanceSkuTierPtrOutput {
	return pulumi.ToOutput(in).(SqlManagedInstanceSkuTierPtrOutput)
}

func (in *sqlManagedInstanceSkuTierPtr) ToSqlManagedInstanceSkuTierPtrOutputWithContext(ctx context.Context) SqlManagedInstanceSkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SqlManagedInstanceSkuTierPtrOutput)
}

// SQL Server version.
type SqlVersion string

const (
	SqlVersion_SQL_Server_2019 = SqlVersion("SQL Server 2019")
	SqlVersion_SQL_Server_2017 = SqlVersion("SQL Server 2017")
	SqlVersion_SQL_Server_2016 = SqlVersion("SQL Server 2016")
)

func init() {
	pulumi.RegisterOutputType(InfrastructureOutput{})
	pulumi.RegisterOutputType(InfrastructurePtrOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceSkuTierOutput{})
	pulumi.RegisterOutputType(SqlManagedInstanceSkuTierPtrOutput{})
}