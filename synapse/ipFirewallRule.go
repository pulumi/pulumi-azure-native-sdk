// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// IP firewall rule
//
// Uses Azure REST API version 2021-06-01. In version 2.x of the Azure Native provider, it used API version 2021-06-01.
//
// Other available API versions: 2021-04-01-preview, 2021-05-01, 2021-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native synapse [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type IpFirewallRule struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The end IP address of the firewall rule. Must be IPv4 format. Must be greater than or equal to startIpAddress
	EndIpAddress pulumi.StringPtrOutput `pulumi:"endIpAddress"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource provisioning state
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The start IP address of the firewall rule. Must be IPv4 format
	StartIpAddress pulumi.StringPtrOutput `pulumi:"startIpAddress"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIpFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewIpFirewallRule(ctx *pulumi.Context,
	name string, args *IpFirewallRuleArgs, opts ...pulumi.ResourceOption) (*IpFirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:IpFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:IpFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:IpFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:IpFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:IpFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:IpFirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:IpFirewallRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IpFirewallRule
	err := ctx.RegisterResource("azure-native:synapse:IpFirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpFirewallRule gets an existing IpFirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpFirewallRuleState, opts ...pulumi.ResourceOption) (*IpFirewallRule, error) {
	var resource IpFirewallRule
	err := ctx.ReadResource("azure-native:synapse:IpFirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpFirewallRule resources.
type ipFirewallRuleState struct {
}

type IpFirewallRuleState struct {
}

func (IpFirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipFirewallRuleState)(nil)).Elem()
}

type ipFirewallRuleArgs struct {
	// The end IP address of the firewall rule. Must be IPv4 format. Must be greater than or equal to startIpAddress
	EndIpAddress *string `pulumi:"endIpAddress"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The IP firewall rule name
	RuleName *string `pulumi:"ruleName"`
	// The start IP address of the firewall rule. Must be IPv4 format
	StartIpAddress *string `pulumi:"startIpAddress"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a IpFirewallRule resource.
type IpFirewallRuleArgs struct {
	// The end IP address of the firewall rule. Must be IPv4 format. Must be greater than or equal to startIpAddress
	EndIpAddress pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The IP firewall rule name
	RuleName pulumi.StringPtrInput
	// The start IP address of the firewall rule. Must be IPv4 format
	StartIpAddress pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (IpFirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipFirewallRuleArgs)(nil)).Elem()
}

type IpFirewallRuleInput interface {
	pulumi.Input

	ToIpFirewallRuleOutput() IpFirewallRuleOutput
	ToIpFirewallRuleOutputWithContext(ctx context.Context) IpFirewallRuleOutput
}

func (*IpFirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**IpFirewallRule)(nil)).Elem()
}

func (i *IpFirewallRule) ToIpFirewallRuleOutput() IpFirewallRuleOutput {
	return i.ToIpFirewallRuleOutputWithContext(context.Background())
}

func (i *IpFirewallRule) ToIpFirewallRuleOutputWithContext(ctx context.Context) IpFirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpFirewallRuleOutput)
}

type IpFirewallRuleOutput struct{ *pulumi.OutputState }

func (IpFirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpFirewallRule)(nil)).Elem()
}

func (o IpFirewallRuleOutput) ToIpFirewallRuleOutput() IpFirewallRuleOutput {
	return o
}

func (o IpFirewallRuleOutput) ToIpFirewallRuleOutputWithContext(ctx context.Context) IpFirewallRuleOutput {
	return o
}

// The Azure API version of the resource.
func (o IpFirewallRuleOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The end IP address of the firewall rule. Must be IPv4 format. Must be greater than or equal to startIpAddress
func (o IpFirewallRuleOutput) EndIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringPtrOutput { return v.EndIpAddress }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o IpFirewallRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource provisioning state
func (o IpFirewallRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The start IP address of the firewall rule. Must be IPv4 format
func (o IpFirewallRuleOutput) StartIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringPtrOutput { return v.StartIpAddress }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o IpFirewallRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IpFirewallRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IpFirewallRuleOutput{})
}
