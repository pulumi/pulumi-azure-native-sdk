// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package azuredata

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A SQL server.
// API Version: 2019-07-24-preview.
type SqlServer struct {
	pulumi.CustomResourceState

	// Cores of the Sql Server.
	Cores pulumi.IntPtrOutput `pulumi:"cores"`
	// Sql Server Edition.
	Edition pulumi.StringPtrOutput `pulumi:"edition"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Sql Server Json Property Bag.
	PropertyBag pulumi.StringPtrOutput `pulumi:"propertyBag"`
	// ID for Parent Sql Server Registration.
	RegistrationID pulumi.StringPtrOutput `pulumi:"registrationID"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
	// Version of the Sql Server.
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewSqlServer registers a new resource with the given unique name, arguments, and options.
func NewSqlServer(ctx *pulumi.Context,
	name string, args *SqlServerArgs, opts ...pulumi.ResourceOption) (*SqlServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SqlServerRegistrationName == nil {
		return nil, errors.New("invalid value for required argument 'SqlServerRegistrationName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:azuredata/v20170301preview:SqlServer"),
		},
		{
			Type: pulumi.String("azure-native:azuredata/v20190724preview:SqlServer"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlServer
	err := ctx.RegisterResource("azure-native:azuredata:SqlServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlServer gets an existing SqlServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlServerState, opts ...pulumi.ResourceOption) (*SqlServer, error) {
	var resource SqlServer
	err := ctx.ReadResource("azure-native:azuredata:SqlServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlServer resources.
type sqlServerState struct {
}

type SqlServerState struct {
}

func (SqlServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerState)(nil)).Elem()
}

type sqlServerArgs struct {
	// Cores of the Sql Server.
	Cores *int `pulumi:"cores"`
	// Sql Server Edition.
	Edition *string `pulumi:"edition"`
	// Sql Server Json Property Bag.
	PropertyBag *string `pulumi:"propertyBag"`
	// ID for Parent Sql Server Registration.
	RegistrationID *string `pulumi:"registrationID"`
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the SQL Server.
	SqlServerName *string `pulumi:"sqlServerName"`
	// Name of the SQL Server registration.
	SqlServerRegistrationName string `pulumi:"sqlServerRegistrationName"`
	// Version of the Sql Server.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a SqlServer resource.
type SqlServerArgs struct {
	// Cores of the Sql Server.
	Cores pulumi.IntPtrInput
	// Sql Server Edition.
	Edition pulumi.StringPtrInput
	// Sql Server Json Property Bag.
	PropertyBag pulumi.StringPtrInput
	// ID for Parent Sql Server Registration.
	RegistrationID pulumi.StringPtrInput
	// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// Name of the SQL Server.
	SqlServerName pulumi.StringPtrInput
	// Name of the SQL Server registration.
	SqlServerRegistrationName pulumi.StringInput
	// Version of the Sql Server.
	Version pulumi.StringPtrInput
}

func (SqlServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlServerArgs)(nil)).Elem()
}

type SqlServerInput interface {
	pulumi.Input

	ToSqlServerOutput() SqlServerOutput
	ToSqlServerOutputWithContext(ctx context.Context) SqlServerOutput
}

func (*SqlServer) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServer)(nil)).Elem()
}

func (i *SqlServer) ToSqlServerOutput() SqlServerOutput {
	return i.ToSqlServerOutputWithContext(context.Background())
}

func (i *SqlServer) ToSqlServerOutputWithContext(ctx context.Context) SqlServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerOutput)
}

type SqlServerOutput struct{ *pulumi.OutputState }

func (SqlServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlServer)(nil)).Elem()
}

func (o SqlServerOutput) ToSqlServerOutput() SqlServerOutput {
	return o
}

func (o SqlServerOutput) ToSqlServerOutputWithContext(ctx context.Context) SqlServerOutput {
	return o
}

// Cores of the Sql Server.
func (o SqlServerOutput) Cores() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlServer) pulumi.IntPtrOutput { return v.Cores }).(pulumi.IntPtrOutput)
}

// Sql Server Edition.
func (o SqlServerOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServer) pulumi.StringPtrOutput { return v.Edition }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o SqlServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Sql Server Json Property Bag.
func (o SqlServerOutput) PropertyBag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServer) pulumi.StringPtrOutput { return v.PropertyBag }).(pulumi.StringPtrOutput)
}

// ID for Parent Sql Server Registration.
func (o SqlServerOutput) RegistrationID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServer) pulumi.StringPtrOutput { return v.RegistrationID }).(pulumi.StringPtrOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o SqlServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Version of the Sql Server.
func (o SqlServerOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlServer) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlServerOutput{})
}