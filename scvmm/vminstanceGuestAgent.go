// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scvmm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Defines the GuestAgent.
//
// Uses Azure REST API version 2023-04-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-04-01-preview.
type VMInstanceGuestAgent struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Username / Password Credentials to provision guest agent.
	Credentials GuestCredentialResponsePtrOutput `pulumi:"credentials"`
	// Gets the name of the corresponding resource in Kubernetes.
	CustomResourceName pulumi.StringOutput `pulumi:"customResourceName"`
	// HTTP Proxy configuration for the VM.
	HttpProxyConfig HttpProxyConfigurationResponsePtrOutput `pulumi:"httpProxyConfig"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the guest agent provisioning action.
	ProvisioningAction pulumi.StringPtrOutput `pulumi:"provisioningAction"`
	// Gets or sets the provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Gets or sets the guest agent status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Gets or sets a unique identifier for this resource.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewVMInstanceGuestAgent registers a new resource with the given unique name, arguments, and options.
func NewVMInstanceGuestAgent(ctx *pulumi.Context,
	name string, args *VMInstanceGuestAgentArgs, opts ...pulumi.ResourceOption) (*VMInstanceGuestAgent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceUri == nil {
		return nil, errors.New("invalid value for required argument 'ResourceUri'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:scvmm/v20230401preview:VMInstanceGuestAgent"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20231007:GuestAgent"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20231007:VMInstanceGuestAgent"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20240601:GuestAgent"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20240601:VMInstanceGuestAgent"),
		},
		{
			Type: pulumi.String("azure-native:scvmm/v20250313:VMInstanceGuestAgent"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource VMInstanceGuestAgent
	err := ctx.RegisterResource("azure-native:scvmm:VMInstanceGuestAgent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVMInstanceGuestAgent gets an existing VMInstanceGuestAgent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVMInstanceGuestAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VMInstanceGuestAgentState, opts ...pulumi.ResourceOption) (*VMInstanceGuestAgent, error) {
	var resource VMInstanceGuestAgent
	err := ctx.ReadResource("azure-native:scvmm:VMInstanceGuestAgent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VMInstanceGuestAgent resources.
type vminstanceGuestAgentState struct {
}

type VMInstanceGuestAgentState struct {
}

func (VMInstanceGuestAgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*vminstanceGuestAgentState)(nil)).Elem()
}

type vminstanceGuestAgentArgs struct {
	// Username / Password Credentials to provision guest agent.
	Credentials *GuestCredential `pulumi:"credentials"`
	// HTTP Proxy configuration for the VM.
	HttpProxyConfig *HttpProxyConfiguration `pulumi:"httpProxyConfig"`
	// Gets or sets the guest agent provisioning action.
	ProvisioningAction *string `pulumi:"provisioningAction"`
	// The fully qualified Azure Resource manager identifier of the Hybrid Compute machine resource to be extended.
	ResourceUri string `pulumi:"resourceUri"`
}

// The set of arguments for constructing a VMInstanceGuestAgent resource.
type VMInstanceGuestAgentArgs struct {
	// Username / Password Credentials to provision guest agent.
	Credentials GuestCredentialPtrInput
	// HTTP Proxy configuration for the VM.
	HttpProxyConfig HttpProxyConfigurationPtrInput
	// Gets or sets the guest agent provisioning action.
	ProvisioningAction pulumi.StringPtrInput
	// The fully qualified Azure Resource manager identifier of the Hybrid Compute machine resource to be extended.
	ResourceUri pulumi.StringInput
}

func (VMInstanceGuestAgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vminstanceGuestAgentArgs)(nil)).Elem()
}

type VMInstanceGuestAgentInput interface {
	pulumi.Input

	ToVMInstanceGuestAgentOutput() VMInstanceGuestAgentOutput
	ToVMInstanceGuestAgentOutputWithContext(ctx context.Context) VMInstanceGuestAgentOutput
}

func (*VMInstanceGuestAgent) ElementType() reflect.Type {
	return reflect.TypeOf((**VMInstanceGuestAgent)(nil)).Elem()
}

func (i *VMInstanceGuestAgent) ToVMInstanceGuestAgentOutput() VMInstanceGuestAgentOutput {
	return i.ToVMInstanceGuestAgentOutputWithContext(context.Background())
}

func (i *VMInstanceGuestAgent) ToVMInstanceGuestAgentOutputWithContext(ctx context.Context) VMInstanceGuestAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMInstanceGuestAgentOutput)
}

type VMInstanceGuestAgentOutput struct{ *pulumi.OutputState }

func (VMInstanceGuestAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMInstanceGuestAgent)(nil)).Elem()
}

func (o VMInstanceGuestAgentOutput) ToVMInstanceGuestAgentOutput() VMInstanceGuestAgentOutput {
	return o
}

func (o VMInstanceGuestAgentOutput) ToVMInstanceGuestAgentOutputWithContext(ctx context.Context) VMInstanceGuestAgentOutput {
	return o
}

// The Azure API version of the resource.
func (o VMInstanceGuestAgentOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *VMInstanceGuestAgent) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Username / Password Credentials to provision guest agent.
func (o VMInstanceGuestAgentOutput) Credentials() GuestCredentialResponsePtrOutput {
	return o.ApplyT(func(v *VMInstanceGuestAgent) GuestCredentialResponsePtrOutput { return v.Credentials }).(GuestCredentialResponsePtrOutput)
}

// Gets the name of the corresponding resource in Kubernetes.
func (o VMInstanceGuestAgentOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VMInstanceGuestAgent) pulumi.StringOutput { return v.CustomResourceName }).(pulumi.StringOutput)
}

// HTTP Proxy configuration for the VM.
func (o VMInstanceGuestAgentOutput) HttpProxyConfig() HttpProxyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VMInstanceGuestAgent) HttpProxyConfigurationResponsePtrOutput { return v.HttpProxyConfig }).(HttpProxyConfigurationResponsePtrOutput)
}

// The name of the resource
func (o VMInstanceGuestAgentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VMInstanceGuestAgent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets or sets the guest agent provisioning action.
func (o VMInstanceGuestAgentOutput) ProvisioningAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMInstanceGuestAgent) pulumi.StringPtrOutput { return v.ProvisioningAction }).(pulumi.StringPtrOutput)
}

// Gets or sets the provisioning state.
func (o VMInstanceGuestAgentOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VMInstanceGuestAgent) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Gets or sets the guest agent status.
func (o VMInstanceGuestAgentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VMInstanceGuestAgent) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o VMInstanceGuestAgentOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VMInstanceGuestAgent) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o VMInstanceGuestAgentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VMInstanceGuestAgent) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Gets or sets a unique identifier for this resource.
func (o VMInstanceGuestAgentOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *VMInstanceGuestAgent) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VMInstanceGuestAgentOutput{})
}
