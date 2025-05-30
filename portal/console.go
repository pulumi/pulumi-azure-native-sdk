// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package portal

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Cloud shell console
//
// Uses Azure REST API version 2018-10-01. In version 2.x of the Azure Native provider, it used API version 2018-10-01.
type Console struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Cloud shell console properties.
	Properties ConsolePropertiesResponseOutput `pulumi:"properties"`
}

// NewConsole registers a new resource with the given unique name, arguments, and options.
func NewConsole(ctx *pulumi.Context,
	name string, args *ConsoleArgs, opts ...pulumi.ResourceOption) (*Console, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:portal/v20181001:Console"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Console
	err := ctx.RegisterResource("azure-native:portal:Console", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsole gets an existing Console resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsoleState, opts ...pulumi.ResourceOption) (*Console, error) {
	var resource Console
	err := ctx.ReadResource("azure-native:portal:Console", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Console resources.
type consoleState struct {
}

type ConsoleState struct {
}

func (ConsoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*consoleState)(nil)).Elem()
}

type consoleArgs struct {
	// The name of the console
	ConsoleName *string `pulumi:"consoleName"`
	// Cloud shell properties for creating a console.
	Properties ConsoleCreateProperties `pulumi:"properties"`
}

// The set of arguments for constructing a Console resource.
type ConsoleArgs struct {
	// The name of the console
	ConsoleName pulumi.StringPtrInput
	// Cloud shell properties for creating a console.
	Properties ConsoleCreatePropertiesInput
}

func (ConsoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consoleArgs)(nil)).Elem()
}

type ConsoleInput interface {
	pulumi.Input

	ToConsoleOutput() ConsoleOutput
	ToConsoleOutputWithContext(ctx context.Context) ConsoleOutput
}

func (*Console) ElementType() reflect.Type {
	return reflect.TypeOf((**Console)(nil)).Elem()
}

func (i *Console) ToConsoleOutput() ConsoleOutput {
	return i.ToConsoleOutputWithContext(context.Background())
}

func (i *Console) ToConsoleOutputWithContext(ctx context.Context) ConsoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsoleOutput)
}

type ConsoleOutput struct{ *pulumi.OutputState }

func (ConsoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Console)(nil)).Elem()
}

func (o ConsoleOutput) ToConsoleOutput() ConsoleOutput {
	return o
}

func (o ConsoleOutput) ToConsoleOutputWithContext(ctx context.Context) ConsoleOutput {
	return o
}

// The Azure API version of the resource.
func (o ConsoleOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Console) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Cloud shell console properties.
func (o ConsoleOutput) Properties() ConsolePropertiesResponseOutput {
	return o.ApplyT(func(v *Console) ConsolePropertiesResponseOutput { return v.Properties }).(ConsolePropertiesResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ConsoleOutput{})
}
