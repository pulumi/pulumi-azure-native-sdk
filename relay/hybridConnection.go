// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package relay

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description of hybrid connection resource.
//
// Uses Azure REST API version 2024-01-01. In version 2.x of the Azure Native provider, it used API version 2021-11-01.
//
// Other available API versions: 2021-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native relay [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type HybridConnection struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The time the hybrid connection was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The number of listeners for this hybrid connection. Note that min : 1 and max:25 are supported.
	ListenerCount pulumi.IntOutput `pulumi:"listenerCount"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Returns true if client authorization is needed for this hybrid connection; otherwise, false.
	RequiresClientAuthorization pulumi.BoolPtrOutput `pulumi:"requiresClientAuthorization"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type pulumi.StringOutput `pulumi:"type"`
	// The time the namespace was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata pulumi.StringPtrOutput `pulumi:"userMetadata"`
}

// NewHybridConnection registers a new resource with the given unique name, arguments, and options.
func NewHybridConnection(ctx *pulumi.Context,
	name string, args *HybridConnectionArgs, opts ...pulumi.ResourceOption) (*HybridConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:relay/v20160701:HybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:relay/v20170401:HybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:relay/v20211101:HybridConnection"),
		},
		{
			Type: pulumi.String("azure-native:relay/v20240101:HybridConnection"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource HybridConnection
	err := ctx.RegisterResource("azure-native:relay:HybridConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHybridConnection gets an existing HybridConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHybridConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridConnectionState, opts ...pulumi.ResourceOption) (*HybridConnection, error) {
	var resource HybridConnection
	err := ctx.ReadResource("azure-native:relay:HybridConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HybridConnection resources.
type hybridConnectionState struct {
}

type HybridConnectionState struct {
}

func (HybridConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridConnectionState)(nil)).Elem()
}

type hybridConnectionArgs struct {
	// The hybrid connection name.
	HybridConnectionName *string `pulumi:"hybridConnectionName"`
	// The namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Returns true if client authorization is needed for this hybrid connection; otherwise, false.
	RequiresClientAuthorization *bool `pulumi:"requiresClientAuthorization"`
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata *string `pulumi:"userMetadata"`
}

// The set of arguments for constructing a HybridConnection resource.
type HybridConnectionArgs struct {
	// The hybrid connection name.
	HybridConnectionName pulumi.StringPtrInput
	// The namespace name
	NamespaceName pulumi.StringInput
	// Returns true if client authorization is needed for this hybrid connection; otherwise, false.
	RequiresClientAuthorization pulumi.BoolPtrInput
	// Name of the Resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata pulumi.StringPtrInput
}

func (HybridConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridConnectionArgs)(nil)).Elem()
}

type HybridConnectionInput interface {
	pulumi.Input

	ToHybridConnectionOutput() HybridConnectionOutput
	ToHybridConnectionOutputWithContext(ctx context.Context) HybridConnectionOutput
}

func (*HybridConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridConnection)(nil)).Elem()
}

func (i *HybridConnection) ToHybridConnectionOutput() HybridConnectionOutput {
	return i.ToHybridConnectionOutputWithContext(context.Background())
}

func (i *HybridConnection) ToHybridConnectionOutputWithContext(ctx context.Context) HybridConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridConnectionOutput)
}

type HybridConnectionOutput struct{ *pulumi.OutputState }

func (HybridConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridConnection)(nil)).Elem()
}

func (o HybridConnectionOutput) ToHybridConnectionOutput() HybridConnectionOutput {
	return o
}

func (o HybridConnectionOutput) ToHybridConnectionOutputWithContext(ctx context.Context) HybridConnectionOutput {
	return o
}

// The Azure API version of the resource.
func (o HybridConnectionOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridConnection) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The time the hybrid connection was created.
func (o HybridConnectionOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridConnection) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The number of listeners for this hybrid connection. Note that min : 1 and max:25 are supported.
func (o HybridConnectionOutput) ListenerCount() pulumi.IntOutput {
	return o.ApplyT(func(v *HybridConnection) pulumi.IntOutput { return v.ListenerCount }).(pulumi.IntOutput)
}

// The geo-location where the resource lives
func (o HybridConnectionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridConnection) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o HybridConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Returns true if client authorization is needed for this hybrid connection; otherwise, false.
func (o HybridConnectionOutput) RequiresClientAuthorization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HybridConnection) pulumi.BoolPtrOutput { return v.RequiresClientAuthorization }).(pulumi.BoolPtrOutput)
}

// The system meta data relating to this resource.
func (o HybridConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *HybridConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
func (o HybridConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The time the namespace was updated.
func (o HybridConnectionOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridConnection) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
func (o HybridConnectionOutput) UserMetadata() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HybridConnection) pulumi.StringPtrOutput { return v.UserMetadata }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(HybridConnectionOutput{})
}
