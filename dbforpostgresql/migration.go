// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbforpostgresql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a migration resource.
//
// Uses Azure REST API version 2024-08-01. In version 2.x of the Azure Native provider, it used API version 2023-03-01-preview.
//
// Other available API versions: 2023-03-01-preview, 2023-06-01-preview, 2023-12-01-preview, 2024-03-01-preview, 2024-11-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native dbforpostgresql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type Migration struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// To trigger cancel for entire migration we need to send this flag as True
	Cancel pulumi.StringPtrOutput `pulumi:"cancel"`
	// Current status of migration
	CurrentStatus MigrationStatusResponseOutput `pulumi:"currentStatus"`
	// When you want to trigger cancel for specific databases send cancel flag as True and database names in this array
	DbsToCancelMigrationOn pulumi.StringArrayOutput `pulumi:"dbsToCancelMigrationOn"`
	// Number of databases to migrate
	DbsToMigrate pulumi.StringArrayOutput `pulumi:"dbsToMigrate"`
	// When you want to trigger cutover for specific databases send triggerCutover flag as True and database names in this array
	DbsToTriggerCutoverOn pulumi.StringArrayOutput `pulumi:"dbsToTriggerCutoverOn"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// To migrate roles and permissions we need to send this flag as True
	MigrateRoles pulumi.StringPtrOutput `pulumi:"migrateRoles"`
	// ID for migration, a GUID.
	MigrationId pulumi.StringOutput `pulumi:"migrationId"`
	// ResourceId of the private endpoint migration instance
	MigrationInstanceResourceId pulumi.StringPtrOutput `pulumi:"migrationInstanceResourceId"`
	// There are two types of migration modes Online and Offline
	MigrationMode pulumi.StringPtrOutput `pulumi:"migrationMode"`
	// This indicates the supported Migration option for the migration
	MigrationOption pulumi.StringPtrOutput `pulumi:"migrationOption"`
	// End time in UTC for migration window
	MigrationWindowEndTimeInUtc pulumi.StringPtrOutput `pulumi:"migrationWindowEndTimeInUtc"`
	// Start time in UTC for migration window
	MigrationWindowStartTimeInUtc pulumi.StringPtrOutput `pulumi:"migrationWindowStartTimeInUtc"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates whether the databases on the target server can be overwritten, if already present. If set to False, the migration workflow will wait for a confirmation, if it detects that the database already exists.
	OverwriteDbsInTarget pulumi.StringPtrOutput `pulumi:"overwriteDbsInTarget"`
	// Indicates whether to setup LogicalReplicationOnSourceDb, if needed
	SetupLogicalReplicationOnSourceDbIfNeeded pulumi.StringPtrOutput `pulumi:"setupLogicalReplicationOnSourceDbIfNeeded"`
	// Source server fully qualified domain name (FQDN) or IP address. It is a optional value, if customer provide it, migration service will always use it for connection
	SourceDbServerFullyQualifiedDomainName pulumi.StringPtrOutput `pulumi:"sourceDbServerFullyQualifiedDomainName"`
	// Metadata of the source database server
	SourceDbServerMetadata DbServerMetadataResponseOutput `pulumi:"sourceDbServerMetadata"`
	// ResourceId of the source database server in case the sourceType is PostgreSQLSingleServer. For other source types this should be ipaddress:port@username or hostname:port@username
	SourceDbServerResourceId pulumi.StringPtrOutput `pulumi:"sourceDbServerResourceId"`
	// migration source server type : OnPremises, AWS, GCP, AzureVM, PostgreSQLSingleServer, AWS_RDS, AWS_AURORA, AWS_EC2, GCP_CloudSQL, GCP_AlloyDB, GCP_Compute, or EDB
	SourceType pulumi.StringPtrOutput `pulumi:"sourceType"`
	// SSL modes for migration. Default SSL mode for PostgreSQLSingleServer is VerifyFull and Prefer for other source types
	SslMode pulumi.StringPtrOutput `pulumi:"sslMode"`
	// Indicates whether the data migration should start right away
	StartDataMigration pulumi.StringPtrOutput `pulumi:"startDataMigration"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Target server fully qualified domain name (FQDN) or IP address. It is a optional value, if customer provide it, migration service will always use it for connection
	TargetDbServerFullyQualifiedDomainName pulumi.StringPtrOutput `pulumi:"targetDbServerFullyQualifiedDomainName"`
	// Metadata of the target database server
	TargetDbServerMetadata DbServerMetadataResponseOutput `pulumi:"targetDbServerMetadata"`
	// ResourceId of the source database server
	TargetDbServerResourceId pulumi.StringOutput `pulumi:"targetDbServerResourceId"`
	// To trigger cutover for entire migration we need to send this flag as True
	TriggerCutover pulumi.StringPtrOutput `pulumi:"triggerCutover"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMigration registers a new resource with the given unique name, arguments, and options.
func NewMigration(ctx *pulumi.Context,
	name string, args *MigrationArgs, opts ...pulumi.ResourceOption) (*Migration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TargetDbServerName == nil {
		return nil, errors.New("invalid value for required argument 'TargetDbServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20210615privatepreview:Migration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20220501preview:Migration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20230301preview:Migration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20230601preview:Migration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20231201preview:Migration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20240301preview:Migration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20240801:Migration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20241101preview:Migration"),
		},
		{
			Type: pulumi.String("azure-native:dbforpostgresql/v20250101preview:Migration"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Migration
	err := ctx.RegisterResource("azure-native:dbforpostgresql:Migration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMigration gets an existing Migration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMigration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MigrationState, opts ...pulumi.ResourceOption) (*Migration, error) {
	var resource Migration
	err := ctx.ReadResource("azure-native:dbforpostgresql:Migration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Migration resources.
type migrationState struct {
}

type MigrationState struct {
}

func (MigrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*migrationState)(nil)).Elem()
}

type migrationArgs struct {
	// To trigger cancel for entire migration we need to send this flag as True
	Cancel *string `pulumi:"cancel"`
	// When you want to trigger cancel for specific databases send cancel flag as True and database names in this array
	DbsToCancelMigrationOn []string `pulumi:"dbsToCancelMigrationOn"`
	// Number of databases to migrate
	DbsToMigrate []string `pulumi:"dbsToMigrate"`
	// When you want to trigger cutover for specific databases send triggerCutover flag as True and database names in this array
	DbsToTriggerCutoverOn []string `pulumi:"dbsToTriggerCutoverOn"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// To migrate roles and permissions we need to send this flag as True
	MigrateRoles *string `pulumi:"migrateRoles"`
	// ResourceId of the private endpoint migration instance
	MigrationInstanceResourceId *string `pulumi:"migrationInstanceResourceId"`
	// There are two types of migration modes Online and Offline
	MigrationMode *string `pulumi:"migrationMode"`
	// The name of the migration.
	MigrationName *string `pulumi:"migrationName"`
	// This indicates the supported Migration option for the migration
	MigrationOption *string `pulumi:"migrationOption"`
	// End time in UTC for migration window
	MigrationWindowEndTimeInUtc *string `pulumi:"migrationWindowEndTimeInUtc"`
	// Start time in UTC for migration window
	MigrationWindowStartTimeInUtc *string `pulumi:"migrationWindowStartTimeInUtc"`
	// Indicates whether the databases on the target server can be overwritten, if already present. If set to False, the migration workflow will wait for a confirmation, if it detects that the database already exists.
	OverwriteDbsInTarget *string `pulumi:"overwriteDbsInTarget"`
	// The resource group name of the target database server.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Migration secret parameters
	SecretParameters *MigrationSecretParameters `pulumi:"secretParameters"`
	// Indicates whether to setup LogicalReplicationOnSourceDb, if needed
	SetupLogicalReplicationOnSourceDbIfNeeded *string `pulumi:"setupLogicalReplicationOnSourceDbIfNeeded"`
	// Source server fully qualified domain name (FQDN) or IP address. It is a optional value, if customer provide it, migration service will always use it for connection
	SourceDbServerFullyQualifiedDomainName *string `pulumi:"sourceDbServerFullyQualifiedDomainName"`
	// ResourceId of the source database server in case the sourceType is PostgreSQLSingleServer. For other source types this should be ipaddress:port@username or hostname:port@username
	SourceDbServerResourceId *string `pulumi:"sourceDbServerResourceId"`
	// migration source server type : OnPremises, AWS, GCP, AzureVM, PostgreSQLSingleServer, AWS_RDS, AWS_AURORA, AWS_EC2, GCP_CloudSQL, GCP_AlloyDB, GCP_Compute, or EDB
	SourceType *string `pulumi:"sourceType"`
	// SSL modes for migration. Default SSL mode for PostgreSQLSingleServer is VerifyFull and Prefer for other source types
	SslMode *string `pulumi:"sslMode"`
	// Indicates whether the data migration should start right away
	StartDataMigration *string `pulumi:"startDataMigration"`
	// The subscription ID of the target database server.
	SubscriptionId *string `pulumi:"subscriptionId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Target server fully qualified domain name (FQDN) or IP address. It is a optional value, if customer provide it, migration service will always use it for connection
	TargetDbServerFullyQualifiedDomainName *string `pulumi:"targetDbServerFullyQualifiedDomainName"`
	// The name of the target database server.
	TargetDbServerName string `pulumi:"targetDbServerName"`
	// To trigger cutover for entire migration we need to send this flag as True
	TriggerCutover *string `pulumi:"triggerCutover"`
}

// The set of arguments for constructing a Migration resource.
type MigrationArgs struct {
	// To trigger cancel for entire migration we need to send this flag as True
	Cancel pulumi.StringPtrInput
	// When you want to trigger cancel for specific databases send cancel flag as True and database names in this array
	DbsToCancelMigrationOn pulumi.StringArrayInput
	// Number of databases to migrate
	DbsToMigrate pulumi.StringArrayInput
	// When you want to trigger cutover for specific databases send triggerCutover flag as True and database names in this array
	DbsToTriggerCutoverOn pulumi.StringArrayInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// To migrate roles and permissions we need to send this flag as True
	MigrateRoles pulumi.StringPtrInput
	// ResourceId of the private endpoint migration instance
	MigrationInstanceResourceId pulumi.StringPtrInput
	// There are two types of migration modes Online and Offline
	MigrationMode pulumi.StringPtrInput
	// The name of the migration.
	MigrationName pulumi.StringPtrInput
	// This indicates the supported Migration option for the migration
	MigrationOption pulumi.StringPtrInput
	// End time in UTC for migration window
	MigrationWindowEndTimeInUtc pulumi.StringPtrInput
	// Start time in UTC for migration window
	MigrationWindowStartTimeInUtc pulumi.StringPtrInput
	// Indicates whether the databases on the target server can be overwritten, if already present. If set to False, the migration workflow will wait for a confirmation, if it detects that the database already exists.
	OverwriteDbsInTarget pulumi.StringPtrInput
	// The resource group name of the target database server.
	ResourceGroupName pulumi.StringInput
	// Migration secret parameters
	SecretParameters MigrationSecretParametersPtrInput
	// Indicates whether to setup LogicalReplicationOnSourceDb, if needed
	SetupLogicalReplicationOnSourceDbIfNeeded pulumi.StringPtrInput
	// Source server fully qualified domain name (FQDN) or IP address. It is a optional value, if customer provide it, migration service will always use it for connection
	SourceDbServerFullyQualifiedDomainName pulumi.StringPtrInput
	// ResourceId of the source database server in case the sourceType is PostgreSQLSingleServer. For other source types this should be ipaddress:port@username or hostname:port@username
	SourceDbServerResourceId pulumi.StringPtrInput
	// migration source server type : OnPremises, AWS, GCP, AzureVM, PostgreSQLSingleServer, AWS_RDS, AWS_AURORA, AWS_EC2, GCP_CloudSQL, GCP_AlloyDB, GCP_Compute, or EDB
	SourceType pulumi.StringPtrInput
	// SSL modes for migration. Default SSL mode for PostgreSQLSingleServer is VerifyFull and Prefer for other source types
	SslMode pulumi.StringPtrInput
	// Indicates whether the data migration should start right away
	StartDataMigration pulumi.StringPtrInput
	// The subscription ID of the target database server.
	SubscriptionId pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Target server fully qualified domain name (FQDN) or IP address. It is a optional value, if customer provide it, migration service will always use it for connection
	TargetDbServerFullyQualifiedDomainName pulumi.StringPtrInput
	// The name of the target database server.
	TargetDbServerName pulumi.StringInput
	// To trigger cutover for entire migration we need to send this flag as True
	TriggerCutover pulumi.StringPtrInput
}

func (MigrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*migrationArgs)(nil)).Elem()
}

type MigrationInput interface {
	pulumi.Input

	ToMigrationOutput() MigrationOutput
	ToMigrationOutputWithContext(ctx context.Context) MigrationOutput
}

func (*Migration) ElementType() reflect.Type {
	return reflect.TypeOf((**Migration)(nil)).Elem()
}

func (i *Migration) ToMigrationOutput() MigrationOutput {
	return i.ToMigrationOutputWithContext(context.Background())
}

func (i *Migration) ToMigrationOutputWithContext(ctx context.Context) MigrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationOutput)
}

type MigrationOutput struct{ *pulumi.OutputState }

func (MigrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Migration)(nil)).Elem()
}

func (o MigrationOutput) ToMigrationOutput() MigrationOutput {
	return o
}

func (o MigrationOutput) ToMigrationOutputWithContext(ctx context.Context) MigrationOutput {
	return o
}

// The Azure API version of the resource.
func (o MigrationOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// To trigger cancel for entire migration we need to send this flag as True
func (o MigrationOutput) Cancel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.Cancel }).(pulumi.StringPtrOutput)
}

// Current status of migration
func (o MigrationOutput) CurrentStatus() MigrationStatusResponseOutput {
	return o.ApplyT(func(v *Migration) MigrationStatusResponseOutput { return v.CurrentStatus }).(MigrationStatusResponseOutput)
}

// When you want to trigger cancel for specific databases send cancel flag as True and database names in this array
func (o MigrationOutput) DbsToCancelMigrationOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringArrayOutput { return v.DbsToCancelMigrationOn }).(pulumi.StringArrayOutput)
}

// Number of databases to migrate
func (o MigrationOutput) DbsToMigrate() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringArrayOutput { return v.DbsToMigrate }).(pulumi.StringArrayOutput)
}

// When you want to trigger cutover for specific databases send triggerCutover flag as True and database names in this array
func (o MigrationOutput) DbsToTriggerCutoverOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringArrayOutput { return v.DbsToTriggerCutoverOn }).(pulumi.StringArrayOutput)
}

// The geo-location where the resource lives
func (o MigrationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// To migrate roles and permissions we need to send this flag as True
func (o MigrationOutput) MigrateRoles() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.MigrateRoles }).(pulumi.StringPtrOutput)
}

// ID for migration, a GUID.
func (o MigrationOutput) MigrationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringOutput { return v.MigrationId }).(pulumi.StringOutput)
}

// ResourceId of the private endpoint migration instance
func (o MigrationOutput) MigrationInstanceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.MigrationInstanceResourceId }).(pulumi.StringPtrOutput)
}

// There are two types of migration modes Online and Offline
func (o MigrationOutput) MigrationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.MigrationMode }).(pulumi.StringPtrOutput)
}

// This indicates the supported Migration option for the migration
func (o MigrationOutput) MigrationOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.MigrationOption }).(pulumi.StringPtrOutput)
}

// End time in UTC for migration window
func (o MigrationOutput) MigrationWindowEndTimeInUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.MigrationWindowEndTimeInUtc }).(pulumi.StringPtrOutput)
}

// Start time in UTC for migration window
func (o MigrationOutput) MigrationWindowStartTimeInUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.MigrationWindowStartTimeInUtc }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o MigrationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Indicates whether the databases on the target server can be overwritten, if already present. If set to False, the migration workflow will wait for a confirmation, if it detects that the database already exists.
func (o MigrationOutput) OverwriteDbsInTarget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.OverwriteDbsInTarget }).(pulumi.StringPtrOutput)
}

// Indicates whether to setup LogicalReplicationOnSourceDb, if needed
func (o MigrationOutput) SetupLogicalReplicationOnSourceDbIfNeeded() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.SetupLogicalReplicationOnSourceDbIfNeeded }).(pulumi.StringPtrOutput)
}

// Source server fully qualified domain name (FQDN) or IP address. It is a optional value, if customer provide it, migration service will always use it for connection
func (o MigrationOutput) SourceDbServerFullyQualifiedDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.SourceDbServerFullyQualifiedDomainName }).(pulumi.StringPtrOutput)
}

// Metadata of the source database server
func (o MigrationOutput) SourceDbServerMetadata() DbServerMetadataResponseOutput {
	return o.ApplyT(func(v *Migration) DbServerMetadataResponseOutput { return v.SourceDbServerMetadata }).(DbServerMetadataResponseOutput)
}

// ResourceId of the source database server in case the sourceType is PostgreSQLSingleServer. For other source types this should be ipaddress:port@username or hostname:port@username
func (o MigrationOutput) SourceDbServerResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.SourceDbServerResourceId }).(pulumi.StringPtrOutput)
}

// migration source server type : OnPremises, AWS, GCP, AzureVM, PostgreSQLSingleServer, AWS_RDS, AWS_AURORA, AWS_EC2, GCP_CloudSQL, GCP_AlloyDB, GCP_Compute, or EDB
func (o MigrationOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.SourceType }).(pulumi.StringPtrOutput)
}

// SSL modes for migration. Default SSL mode for PostgreSQLSingleServer is VerifyFull and Prefer for other source types
func (o MigrationOutput) SslMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.SslMode }).(pulumi.StringPtrOutput)
}

// Indicates whether the data migration should start right away
func (o MigrationOutput) StartDataMigration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.StartDataMigration }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MigrationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Migration) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o MigrationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Target server fully qualified domain name (FQDN) or IP address. It is a optional value, if customer provide it, migration service will always use it for connection
func (o MigrationOutput) TargetDbServerFullyQualifiedDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.TargetDbServerFullyQualifiedDomainName }).(pulumi.StringPtrOutput)
}

// Metadata of the target database server
func (o MigrationOutput) TargetDbServerMetadata() DbServerMetadataResponseOutput {
	return o.ApplyT(func(v *Migration) DbServerMetadataResponseOutput { return v.TargetDbServerMetadata }).(DbServerMetadataResponseOutput)
}

// ResourceId of the source database server
func (o MigrationOutput) TargetDbServerResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringOutput { return v.TargetDbServerResourceId }).(pulumi.StringOutput)
}

// To trigger cutover for entire migration we need to send this flag as True
func (o MigrationOutput) TriggerCutover() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringPtrOutput { return v.TriggerCutover }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MigrationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Migration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MigrationOutput{})
}
