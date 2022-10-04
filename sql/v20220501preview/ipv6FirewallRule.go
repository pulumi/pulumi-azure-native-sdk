// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An IPv6 server firewall rule.
type IPv6FirewallRule struct {
	pulumi.CustomResourceState

	// The end IP address of the firewall rule. Must be IPv6 format. Must be greater than or equal to startIpAddress.
	EndIPv6Address pulumi.StringPtrOutput `pulumi:"endIPv6Address"`
	// Resource name.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The start IP address of the firewall rule. Must be IPv6 format.
	StartIPv6Address pulumi.StringPtrOutput `pulumi:"startIPv6Address"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIPv6FirewallRule registers a new resource with the given unique name, arguments, and options.
func NewIPv6FirewallRule(ctx *pulumi.Context,
	name string, args *IPv6FirewallRuleArgs, opts ...pulumi.ResourceOption) (*IPv6FirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:IPv6FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210801preview:IPv6FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:IPv6FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:IPv6FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:IPv6FirewallRule"),
		},
	})
	opts = append(opts, aliases)
	var resource IPv6FirewallRule
	err := ctx.RegisterResource("azure-native:sql/v20220501preview:IPv6FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIPv6FirewallRule gets an existing IPv6FirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIPv6FirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IPv6FirewallRuleState, opts ...pulumi.ResourceOption) (*IPv6FirewallRule, error) {
	var resource IPv6FirewallRule
	err := ctx.ReadResource("azure-native:sql/v20220501preview:IPv6FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IPv6FirewallRule resources.
type ipv6FirewallRuleState struct {
}

type IPv6FirewallRuleState struct {
}

func (IPv6FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv6FirewallRuleState)(nil)).Elem()
}

type ipv6FirewallRuleArgs struct {
	// The end IP address of the firewall rule. Must be IPv6 format. Must be greater than or equal to startIpAddress.
	EndIPv6Address *string `pulumi:"endIPv6Address"`
	// The name of the firewall rule.
	FirewallRuleName *string `pulumi:"firewallRuleName"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The start IP address of the firewall rule. Must be IPv6 format.
	StartIPv6Address *string `pulumi:"startIPv6Address"`
}

// The set of arguments for constructing a IPv6FirewallRule resource.
type IPv6FirewallRuleArgs struct {
	// The end IP address of the firewall rule. Must be IPv6 format. Must be greater than or equal to startIpAddress.
	EndIPv6Address pulumi.StringPtrInput
	// The name of the firewall rule.
	FirewallRuleName pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The start IP address of the firewall rule. Must be IPv6 format.
	StartIPv6Address pulumi.StringPtrInput
}

func (IPv6FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv6FirewallRuleArgs)(nil)).Elem()
}

type IPv6FirewallRuleInput interface {
	pulumi.Input

	ToIPv6FirewallRuleOutput() IPv6FirewallRuleOutput
	ToIPv6FirewallRuleOutputWithContext(ctx context.Context) IPv6FirewallRuleOutput
}

func (*IPv6FirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**IPv6FirewallRule)(nil)).Elem()
}

func (i *IPv6FirewallRule) ToIPv6FirewallRuleOutput() IPv6FirewallRuleOutput {
	return i.ToIPv6FirewallRuleOutputWithContext(context.Background())
}

func (i *IPv6FirewallRule) ToIPv6FirewallRuleOutputWithContext(ctx context.Context) IPv6FirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPv6FirewallRuleOutput)
}

type IPv6FirewallRuleOutput struct{ *pulumi.OutputState }

func (IPv6FirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPv6FirewallRule)(nil)).Elem()
}

func (o IPv6FirewallRuleOutput) ToIPv6FirewallRuleOutput() IPv6FirewallRuleOutput {
	return o
}

func (o IPv6FirewallRuleOutput) ToIPv6FirewallRuleOutputWithContext(ctx context.Context) IPv6FirewallRuleOutput {
	return o
}

// The end IP address of the firewall rule. Must be IPv6 format. Must be greater than or equal to startIpAddress.
func (o IPv6FirewallRuleOutput) EndIPv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPv6FirewallRule) pulumi.StringPtrOutput { return v.EndIPv6Address }).(pulumi.StringPtrOutput)
}

// Resource name.
func (o IPv6FirewallRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPv6FirewallRule) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The start IP address of the firewall rule. Must be IPv6 format.
func (o IPv6FirewallRuleOutput) StartIPv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IPv6FirewallRule) pulumi.StringPtrOutput { return v.StartIPv6Address }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o IPv6FirewallRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IPv6FirewallRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IPv6FirewallRuleOutput{})
}