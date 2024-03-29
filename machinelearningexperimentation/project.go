// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningexperimentation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An object that represents a machine learning project.
// Azure REST API version: 2017-05-01-preview. Prior API version in Azure Native 1.x: 2017-05-01-preview.
type Project struct {
	pulumi.CustomResourceState

	// The immutable id of the team account which contains this project.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The creation date of the project in ISO8601 format.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The description of this project.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The friendly name for this project.
	FriendlyName pulumi.StringOutput `pulumi:"friendlyName"`
	// The reference to git repo for this project.
	Gitrepo pulumi.StringPtrOutput `pulumi:"gitrepo"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The immutable id of this project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The current deployment state of project resource. The provisioningState is to indicate states for resource provisioning.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The immutable id of the workspace which contains this project.
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.FriendlyName == nil {
		return nil, errors.New("invalid value for required argument 'FriendlyName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningexperimentation/v20170501preview:Project"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Project
	err := ctx.RegisterResource("azure-native:machinelearningexperimentation:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("azure-native:machinelearningexperimentation:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
}

type ProjectState struct {
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// The name of the machine learning team account.
	AccountName string `pulumi:"accountName"`
	// The description of this project.
	Description *string `pulumi:"description"`
	// The friendly name for this project.
	FriendlyName string `pulumi:"friendlyName"`
	// The reference to git repo for this project.
	Gitrepo *string `pulumi:"gitrepo"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location *string `pulumi:"location"`
	// The name of the machine learning project under a team account workspace.
	ProjectName *string `pulumi:"projectName"`
	// The name of the resource group to which the machine learning team account belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The name of the machine learning team account workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// The name of the machine learning team account.
	AccountName pulumi.StringInput
	// The description of this project.
	Description pulumi.StringPtrInput
	// The friendly name for this project.
	FriendlyName pulumi.StringInput
	// The reference to git repo for this project.
	Gitrepo pulumi.StringPtrInput
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringPtrInput
	// The name of the machine learning project under a team account workspace.
	ProjectName pulumi.StringPtrInput
	// The name of the resource group to which the machine learning team account belongs.
	ResourceGroupName pulumi.StringInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The name of the machine learning team account workspace.
	WorkspaceName pulumi.StringInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

// The immutable id of the team account which contains this project.
func (o ProjectOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The creation date of the project in ISO8601 format.
func (o ProjectOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// The description of this project.
func (o ProjectOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The friendly name for this project.
func (o ProjectOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.FriendlyName }).(pulumi.StringOutput)
}

// The reference to git repo for this project.
func (o ProjectOutput) Gitrepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.Gitrepo }).(pulumi.StringPtrOutput)
}

// The location of the resource. This cannot be changed after the resource is created.
func (o ProjectOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource.
func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The immutable id of this project.
func (o ProjectOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The current deployment state of project resource. The provisioningState is to indicate states for resource provisioning.
func (o ProjectOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The tags of the resource.
func (o ProjectOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Project) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o ProjectOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The immutable id of the workspace which contains this project.
func (o ProjectOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProjectOutput{})
}
