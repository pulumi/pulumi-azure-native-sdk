// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package migrate

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Migrate project.
//
// Uses Azure REST API version 2020-05-01. In version 2.x of the Azure Native provider, it used API version 2020-05-01.
//
// Other available API versions: 2023-01-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type MigrateProjectsControllerMigrateProject struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// For optimistic concurrency control.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Azure location in which project is created.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Name of the project.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of a migrate project.
	Properties MigrateProjectPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Type of the object = [Microsoft.Migrate/migrateProjects].
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMigrateProjectsControllerMigrateProject registers a new resource with the given unique name, arguments, and options.
func NewMigrateProjectsControllerMigrateProject(ctx *pulumi.Context,
	name string, args *MigrateProjectsControllerMigrateProjectArgs, opts ...pulumi.ResourceOption) (*MigrateProjectsControllerMigrateProject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:migrate/v20180901preview:MigrateProject"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20180901preview:MigrateProjectsControllerMigrateProject"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20200501:MigrateProjectsControllerMigrateProject"),
		},
		{
			Type: pulumi.String("azure-native:migrate/v20230101:MigrateProjectsControllerMigrateProject"),
		},
		{
			Type: pulumi.String("azure-native:migrate:MigrateProject"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MigrateProjectsControllerMigrateProject
	err := ctx.RegisterResource("azure-native:migrate:MigrateProjectsControllerMigrateProject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMigrateProjectsControllerMigrateProject gets an existing MigrateProjectsControllerMigrateProject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMigrateProjectsControllerMigrateProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MigrateProjectsControllerMigrateProjectState, opts ...pulumi.ResourceOption) (*MigrateProjectsControllerMigrateProject, error) {
	var resource MigrateProjectsControllerMigrateProject
	err := ctx.ReadResource("azure-native:migrate:MigrateProjectsControllerMigrateProject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MigrateProjectsControllerMigrateProject resources.
type migrateProjectsControllerMigrateProjectState struct {
}

type MigrateProjectsControllerMigrateProjectState struct {
}

func (MigrateProjectsControllerMigrateProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*migrateProjectsControllerMigrateProjectState)(nil)).Elem()
}

type migrateProjectsControllerMigrateProjectArgs struct {
	// For optimistic concurrency control.
	ETag *string `pulumi:"eTag"`
	// Azure location in which project is created.
	Location *string `pulumi:"location"`
	// Migrate project name.
	MigrateProjectName *string `pulumi:"migrateProjectName"`
	// Properties of a migrate project.
	Properties *MigrateProjectProperties `pulumi:"properties"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a MigrateProjectsControllerMigrateProject resource.
type MigrateProjectsControllerMigrateProjectArgs struct {
	// For optimistic concurrency control.
	ETag pulumi.StringPtrInput
	// Azure location in which project is created.
	Location pulumi.StringPtrInput
	// Migrate project name.
	MigrateProjectName pulumi.StringPtrInput
	// Properties of a migrate project.
	Properties MigrateProjectPropertiesPtrInput
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName pulumi.StringInput
}

func (MigrateProjectsControllerMigrateProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*migrateProjectsControllerMigrateProjectArgs)(nil)).Elem()
}

type MigrateProjectsControllerMigrateProjectInput interface {
	pulumi.Input

	ToMigrateProjectsControllerMigrateProjectOutput() MigrateProjectsControllerMigrateProjectOutput
	ToMigrateProjectsControllerMigrateProjectOutputWithContext(ctx context.Context) MigrateProjectsControllerMigrateProjectOutput
}

func (*MigrateProjectsControllerMigrateProject) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProjectsControllerMigrateProject)(nil)).Elem()
}

func (i *MigrateProjectsControllerMigrateProject) ToMigrateProjectsControllerMigrateProjectOutput() MigrateProjectsControllerMigrateProjectOutput {
	return i.ToMigrateProjectsControllerMigrateProjectOutputWithContext(context.Background())
}

func (i *MigrateProjectsControllerMigrateProject) ToMigrateProjectsControllerMigrateProjectOutputWithContext(ctx context.Context) MigrateProjectsControllerMigrateProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectsControllerMigrateProjectOutput)
}

type MigrateProjectsControllerMigrateProjectOutput struct{ *pulumi.OutputState }

func (MigrateProjectsControllerMigrateProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProjectsControllerMigrateProject)(nil)).Elem()
}

func (o MigrateProjectsControllerMigrateProjectOutput) ToMigrateProjectsControllerMigrateProjectOutput() MigrateProjectsControllerMigrateProjectOutput {
	return o
}

func (o MigrateProjectsControllerMigrateProjectOutput) ToMigrateProjectsControllerMigrateProjectOutputWithContext(ctx context.Context) MigrateProjectsControllerMigrateProjectOutput {
	return o
}

// The Azure API version of the resource.
func (o MigrateProjectsControllerMigrateProjectOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *MigrateProjectsControllerMigrateProject) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// For optimistic concurrency control.
func (o MigrateProjectsControllerMigrateProjectOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrateProjectsControllerMigrateProject) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

// Azure location in which project is created.
func (o MigrateProjectsControllerMigrateProjectOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrateProjectsControllerMigrateProject) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Name of the project.
func (o MigrateProjectsControllerMigrateProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MigrateProjectsControllerMigrateProject) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of a migrate project.
func (o MigrateProjectsControllerMigrateProjectOutput) Properties() MigrateProjectPropertiesResponseOutput {
	return o.ApplyT(func(v *MigrateProjectsControllerMigrateProject) MigrateProjectPropertiesResponseOutput {
		return v.Properties
	}).(MigrateProjectPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o MigrateProjectsControllerMigrateProjectOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MigrateProjectsControllerMigrateProject) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Type of the object = [Microsoft.Migrate/migrateProjects].
func (o MigrateProjectsControllerMigrateProjectOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MigrateProjectsControllerMigrateProject) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MigrateProjectsControllerMigrateProjectOutput{})
}
