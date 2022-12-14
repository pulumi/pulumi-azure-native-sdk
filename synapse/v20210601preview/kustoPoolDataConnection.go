// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Class representing a data connection.
//
// Deprecated: Please use one of the variants: EventGridDataConnection, EventHubDataConnection, IotHubDataConnection.
type KustoPoolDataConnection struct {
	pulumi.CustomResourceState

	// Kind of the endpoint for the data connection
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewKustoPoolDataConnection registers a new resource with the given unique name, arguments, and options.
func NewKustoPoolDataConnection(ctx *pulumi.Context,
	name string, args *KustoPoolDataConnectionArgs, opts ...pulumi.ResourceOption) (*KustoPoolDataConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.KustoPoolName == nil {
		return nil, errors.New("invalid value for required argument 'KustoPoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:KustoPoolDataConnection"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:KustoPoolDataConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource KustoPoolDataConnection
	err := ctx.RegisterResource("azure-native:synapse/v20210601preview:KustoPoolDataConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKustoPoolDataConnection gets an existing KustoPoolDataConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKustoPoolDataConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KustoPoolDataConnectionState, opts ...pulumi.ResourceOption) (*KustoPoolDataConnection, error) {
	var resource KustoPoolDataConnection
	err := ctx.ReadResource("azure-native:synapse/v20210601preview:KustoPoolDataConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KustoPoolDataConnection resources.
type kustoPoolDataConnectionState struct {
}

type KustoPoolDataConnectionState struct {
}

func (KustoPoolDataConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolDataConnectionState)(nil)).Elem()
}

type kustoPoolDataConnectionArgs struct {
	// The name of the data connection.
	DataConnectionName *string `pulumi:"dataConnectionName"`
	// The name of the database in the Kusto pool.
	DatabaseName string `pulumi:"databaseName"`
	// Kind of the endpoint for the data connection
	Kind string `pulumi:"kind"`
	// The name of the Kusto pool.
	KustoPoolName string `pulumi:"kustoPoolName"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a KustoPoolDataConnection resource.
type KustoPoolDataConnectionArgs struct {
	// The name of the data connection.
	DataConnectionName pulumi.StringPtrInput
	// The name of the database in the Kusto pool.
	DatabaseName pulumi.StringInput
	// Kind of the endpoint for the data connection
	Kind pulumi.StringInput
	// The name of the Kusto pool.
	KustoPoolName pulumi.StringInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (KustoPoolDataConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kustoPoolDataConnectionArgs)(nil)).Elem()
}

type KustoPoolDataConnectionInput interface {
	pulumi.Input

	ToKustoPoolDataConnectionOutput() KustoPoolDataConnectionOutput
	ToKustoPoolDataConnectionOutputWithContext(ctx context.Context) KustoPoolDataConnectionOutput
}

func (*KustoPoolDataConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoPoolDataConnection)(nil)).Elem()
}

func (i *KustoPoolDataConnection) ToKustoPoolDataConnectionOutput() KustoPoolDataConnectionOutput {
	return i.ToKustoPoolDataConnectionOutputWithContext(context.Background())
}

func (i *KustoPoolDataConnection) ToKustoPoolDataConnectionOutputWithContext(ctx context.Context) KustoPoolDataConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KustoPoolDataConnectionOutput)
}

type KustoPoolDataConnectionOutput struct{ *pulumi.OutputState }

func (KustoPoolDataConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KustoPoolDataConnection)(nil)).Elem()
}

func (o KustoPoolDataConnectionOutput) ToKustoPoolDataConnectionOutput() KustoPoolDataConnectionOutput {
	return o
}

func (o KustoPoolDataConnectionOutput) ToKustoPoolDataConnectionOutputWithContext(ctx context.Context) KustoPoolDataConnectionOutput {
	return o
}

// Kind of the endpoint for the data connection
func (o KustoPoolDataConnectionOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDataConnection) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Resource location.
func (o KustoPoolDataConnectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KustoPoolDataConnection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o KustoPoolDataConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDataConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o KustoPoolDataConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *KustoPoolDataConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o KustoPoolDataConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *KustoPoolDataConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KustoPoolDataConnectionOutput{})
}
