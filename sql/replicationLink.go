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

// A replication link.
//
// Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2023-05-01-preview.
//
// Other available API versions: 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ReplicationLink struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Whether the user is currently allowed to terminate the link.
	IsTerminationAllowed pulumi.BoolOutput `pulumi:"isTerminationAllowed"`
	// Link type (GEO, NAMED, STANDBY). Update operation does not support NAMED.
	LinkType pulumi.StringPtrOutput `pulumi:"linkType"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource partner database.
	PartnerDatabase pulumi.StringOutput `pulumi:"partnerDatabase"`
	// Resource partner database Id.
	PartnerDatabaseId pulumi.StringOutput `pulumi:"partnerDatabaseId"`
	// Resource partner location.
	PartnerLocation pulumi.StringOutput `pulumi:"partnerLocation"`
	// Partner replication role.
	PartnerRole pulumi.StringOutput `pulumi:"partnerRole"`
	// Resource partner server.
	PartnerServer pulumi.StringOutput `pulumi:"partnerServer"`
	// Seeding completion percentage for the link.
	PercentComplete pulumi.IntOutput `pulumi:"percentComplete"`
	// Replication mode.
	ReplicationMode pulumi.StringOutput `pulumi:"replicationMode"`
	// Replication state (PENDING, SEEDING, CATCHUP, SUSPENDED).
	ReplicationState pulumi.StringOutput `pulumi:"replicationState"`
	// Local replication role.
	Role pulumi.StringOutput `pulumi:"role"`
	// Time at which the link was created.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewReplicationLink registers a new resource with the given unique name, arguments, and options.
func NewReplicationLink(ctx *pulumi.Context,
	name string, args *ReplicationLinkArgs, opts ...pulumi.ResourceOption) (*ReplicationLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
			Type: pulumi.String("azure-native:sql/v20230501preview:ReplicationLink"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230801:ReplicationLink"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20230801preview:ReplicationLink"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20240501preview:ReplicationLink"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ReplicationLink
	err := ctx.RegisterResource("azure-native:sql:ReplicationLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationLink gets an existing ReplicationLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationLinkState, opts ...pulumi.ResourceOption) (*ReplicationLink, error) {
	var resource ReplicationLink
	err := ctx.ReadResource("azure-native:sql:ReplicationLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationLink resources.
type replicationLinkState struct {
}

type ReplicationLinkState struct {
}

func (ReplicationLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationLinkState)(nil)).Elem()
}

type replicationLinkArgs struct {
	// The name of the database.
	DatabaseName string  `pulumi:"databaseName"`
	LinkId       *string `pulumi:"linkId"`
	// Link type (GEO, NAMED, STANDBY). Update operation does not support NAMED.
	LinkType *string `pulumi:"linkType"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// The set of arguments for constructing a ReplicationLink resource.
type ReplicationLinkArgs struct {
	// The name of the database.
	DatabaseName pulumi.StringInput
	LinkId       pulumi.StringPtrInput
	// Link type (GEO, NAMED, STANDBY). Update operation does not support NAMED.
	LinkType pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
}

func (ReplicationLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationLinkArgs)(nil)).Elem()
}

type ReplicationLinkInput interface {
	pulumi.Input

	ToReplicationLinkOutput() ReplicationLinkOutput
	ToReplicationLinkOutputWithContext(ctx context.Context) ReplicationLinkOutput
}

func (*ReplicationLink) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationLink)(nil)).Elem()
}

func (i *ReplicationLink) ToReplicationLinkOutput() ReplicationLinkOutput {
	return i.ToReplicationLinkOutputWithContext(context.Background())
}

func (i *ReplicationLink) ToReplicationLinkOutputWithContext(ctx context.Context) ReplicationLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationLinkOutput)
}

type ReplicationLinkOutput struct{ *pulumi.OutputState }

func (ReplicationLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationLink)(nil)).Elem()
}

func (o ReplicationLinkOutput) ToReplicationLinkOutput() ReplicationLinkOutput {
	return o
}

func (o ReplicationLinkOutput) ToReplicationLinkOutputWithContext(ctx context.Context) ReplicationLinkOutput {
	return o
}

// The Azure API version of the resource.
func (o ReplicationLinkOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationLink) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Whether the user is currently allowed to terminate the link.
func (o ReplicationLinkOutput) IsTerminationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *ReplicationLink) pulumi.BoolOutput { return v.IsTerminationAllowed }).(pulumi.BoolOutput)
}

// Link type (GEO, NAMED, STANDBY). Update operation does not support NAMED.
func (o ReplicationLinkOutput) LinkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationLink) pulumi.StringPtrOutput { return v.LinkType }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o ReplicationLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource partner database.
func (o ReplicationLinkOutput) PartnerDatabase() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationLink) pulumi.StringOutput { return v.PartnerDatabase }).(pulumi.StringOutput)
}

// Resource partner database Id.
func (o ReplicationLinkOutput) PartnerDatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationLink) pulumi.StringOutput { return v.PartnerDatabaseId }).(pulumi.StringOutput)
}

// Resource partner location.
func (o ReplicationLinkOutput) PartnerLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationLink) pulumi.StringOutput { return v.PartnerLocation }).(pulumi.StringOutput)
}

// Partner replication role.
func (o ReplicationLinkOutput) PartnerRole() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationLink) pulumi.StringOutput { return v.PartnerRole }).(pulumi.StringOutput)
}

// Resource partner server.
func (o ReplicationLinkOutput) PartnerServer() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationLink) pulumi.StringOutput { return v.PartnerServer }).(pulumi.StringOutput)
}

// Seeding completion percentage for the link.
func (o ReplicationLinkOutput) PercentComplete() pulumi.IntOutput {
	return o.ApplyT(func(v *ReplicationLink) pulumi.IntOutput { return v.PercentComplete }).(pulumi.IntOutput)
}

// Replication mode.
func (o ReplicationLinkOutput) ReplicationMode() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationLink) pulumi.StringOutput { return v.ReplicationMode }).(pulumi.StringOutput)
}

// Replication state (PENDING, SEEDING, CATCHUP, SUSPENDED).
func (o ReplicationLinkOutput) ReplicationState() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationLink) pulumi.StringOutput { return v.ReplicationState }).(pulumi.StringOutput)
}

// Local replication role.
func (o ReplicationLinkOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationLink) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// Time at which the link was created.
func (o ReplicationLinkOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationLink) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

// Resource type.
func (o ReplicationLinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ReplicationLink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplicationLinkOutput{})
}
