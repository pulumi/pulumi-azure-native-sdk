// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The top level data export resource container.
type DataExport struct {
	pulumi.CustomResourceState

	// The latest data export rule modification time.
	CreatedDate pulumi.StringPtrOutput `pulumi:"createdDate"`
	// The data export rule ID.
	DataExportId pulumi.StringPtrOutput `pulumi:"dataExportId"`
	// Active when enabled.
	Enable pulumi.BoolPtrOutput `pulumi:"enable"`
	// Optional. Allows to define an Event Hub name. Not applicable when destination is Storage Account.
	EventHubName pulumi.StringPtrOutput `pulumi:"eventHubName"`
	// Date and time when the export was last modified.
	LastModifiedDate pulumi.StringPtrOutput `pulumi:"lastModifiedDate"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The destination resource ID. This can be copied from the Properties entry of the destination resource in Azure.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// An array of tables to export, for example: [“Heartbeat, SecurityEvent”].
	TableNames pulumi.StringArrayOutput `pulumi:"tableNames"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataExport registers a new resource with the given unique name, arguments, and options.
func NewDataExport(ctx *pulumi.Context,
	name string, args *DataExportArgs, opts ...pulumi.ResourceOption) (*DataExport, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.TableNames == nil {
		return nil, errors.New("invalid value for required argument 'TableNames'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:operationalinsights:DataExport"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20190801preview:DataExport"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:DataExport"),
		},
	})
	opts = append(opts, aliases)
	var resource DataExport
	err := ctx.RegisterResource("azure-native:operationalinsights/v20200801:DataExport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataExport gets an existing DataExport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataExport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataExportState, opts ...pulumi.ResourceOption) (*DataExport, error) {
	var resource DataExport
	err := ctx.ReadResource("azure-native:operationalinsights/v20200801:DataExport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataExport resources.
type dataExportState struct {
}

type DataExportState struct {
}

func (DataExportState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataExportState)(nil)).Elem()
}

type dataExportArgs struct {
	// The latest data export rule modification time.
	CreatedDate *string `pulumi:"createdDate"`
	// The data export rule ID.
	DataExportId *string `pulumi:"dataExportId"`
	// The data export rule name.
	DataExportName *string `pulumi:"dataExportName"`
	// Active when enabled.
	Enable *bool `pulumi:"enable"`
	// Optional. Allows to define an Event Hub name. Not applicable when destination is Storage Account.
	EventHubName *string `pulumi:"eventHubName"`
	// Date and time when the export was last modified.
	LastModifiedDate *string `pulumi:"lastModifiedDate"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The destination resource ID. This can be copied from the Properties entry of the destination resource in Azure.
	ResourceId string `pulumi:"resourceId"`
	// An array of tables to export, for example: [“Heartbeat, SecurityEvent”].
	TableNames []string `pulumi:"tableNames"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a DataExport resource.
type DataExportArgs struct {
	// The latest data export rule modification time.
	CreatedDate pulumi.StringPtrInput
	// The data export rule ID.
	DataExportId pulumi.StringPtrInput
	// The data export rule name.
	DataExportName pulumi.StringPtrInput
	// Active when enabled.
	Enable pulumi.BoolPtrInput
	// Optional. Allows to define an Event Hub name. Not applicable when destination is Storage Account.
	EventHubName pulumi.StringPtrInput
	// Date and time when the export was last modified.
	LastModifiedDate pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The destination resource ID. This can be copied from the Properties entry of the destination resource in Azure.
	ResourceId pulumi.StringInput
	// An array of tables to export, for example: [“Heartbeat, SecurityEvent”].
	TableNames pulumi.StringArrayInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (DataExportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataExportArgs)(nil)).Elem()
}

type DataExportInput interface {
	pulumi.Input

	ToDataExportOutput() DataExportOutput
	ToDataExportOutputWithContext(ctx context.Context) DataExportOutput
}

func (*DataExport) ElementType() reflect.Type {
	return reflect.TypeOf((**DataExport)(nil)).Elem()
}

func (i *DataExport) ToDataExportOutput() DataExportOutput {
	return i.ToDataExportOutputWithContext(context.Background())
}

func (i *DataExport) ToDataExportOutputWithContext(ctx context.Context) DataExportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataExportOutput)
}

type DataExportOutput struct{ *pulumi.OutputState }

func (DataExportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataExport)(nil)).Elem()
}

func (o DataExportOutput) ToDataExportOutput() DataExportOutput {
	return o
}

func (o DataExportOutput) ToDataExportOutputWithContext(ctx context.Context) DataExportOutput {
	return o
}

// The latest data export rule modification time.
func (o DataExportOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataExport) pulumi.StringPtrOutput { return v.CreatedDate }).(pulumi.StringPtrOutput)
}

// The data export rule ID.
func (o DataExportOutput) DataExportId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataExport) pulumi.StringPtrOutput { return v.DataExportId }).(pulumi.StringPtrOutput)
}

// Active when enabled.
func (o DataExportOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DataExport) pulumi.BoolPtrOutput { return v.Enable }).(pulumi.BoolPtrOutput)
}

// Optional. Allows to define an Event Hub name. Not applicable when destination is Storage Account.
func (o DataExportOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataExport) pulumi.StringPtrOutput { return v.EventHubName }).(pulumi.StringPtrOutput)
}

// Date and time when the export was last modified.
func (o DataExportOutput) LastModifiedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataExport) pulumi.StringPtrOutput { return v.LastModifiedDate }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o DataExportOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExport) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The destination resource ID. This can be copied from the Properties entry of the destination resource in Azure.
func (o DataExportOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExport) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// An array of tables to export, for example: [“Heartbeat, SecurityEvent”].
func (o DataExportOutput) TableNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataExport) pulumi.StringArrayOutput { return v.TableNames }).(pulumi.StringArrayOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o DataExportOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExport) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DataExportOutput{})
}