// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20171111preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Blueprint artifact deploys Azure resource manager template.
func LookupTemplateArtifact(ctx *pulumi.Context, args *LookupTemplateArtifactArgs, opts ...pulumi.InvokeOption) (*LookupTemplateArtifactResult, error) {
	var rv LookupTemplateArtifactResult
	err := ctx.Invoke("azure-native:blueprint/v20171111preview:getTemplateArtifact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTemplateArtifactArgs struct {
	ArtifactName        string `pulumi:"artifactName"`
	BlueprintName       string `pulumi:"blueprintName"`
	ManagementGroupName string `pulumi:"managementGroupName"`
}

// Blueprint artifact deploys Azure resource manager template.
type LookupTemplateArtifactResult struct {
	DependsOn     []string                              `pulumi:"dependsOn"`
	Description   *string                               `pulumi:"description"`
	DisplayName   *string                               `pulumi:"displayName"`
	Id            string                                `pulumi:"id"`
	Kind          string                                `pulumi:"kind"`
	Name          string                                `pulumi:"name"`
	Parameters    map[string]ParameterValueBaseResponse `pulumi:"parameters"`
	ResourceGroup *string                               `pulumi:"resourceGroup"`
	Template      interface{}                           `pulumi:"template"`
	Type          string                                `pulumi:"type"`
}

func LookupTemplateArtifactOutput(ctx *pulumi.Context, args LookupTemplateArtifactOutputArgs, opts ...pulumi.InvokeOption) LookupTemplateArtifactResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTemplateArtifactResult, error) {
			args := v.(LookupTemplateArtifactArgs)
			r, err := LookupTemplateArtifact(ctx, &args, opts...)
			var s LookupTemplateArtifactResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTemplateArtifactResultOutput)
}

type LookupTemplateArtifactOutputArgs struct {
	ArtifactName        pulumi.StringInput `pulumi:"artifactName"`
	BlueprintName       pulumi.StringInput `pulumi:"blueprintName"`
	ManagementGroupName pulumi.StringInput `pulumi:"managementGroupName"`
}

func (LookupTemplateArtifactOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateArtifactArgs)(nil)).Elem()
}

// Blueprint artifact deploys Azure resource manager template.
type LookupTemplateArtifactResultOutput struct{ *pulumi.OutputState }

func (LookupTemplateArtifactResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplateArtifactResult)(nil)).Elem()
}

func (o LookupTemplateArtifactResultOutput) ToLookupTemplateArtifactResultOutput() LookupTemplateArtifactResultOutput {
	return o
}

func (o LookupTemplateArtifactResultOutput) ToLookupTemplateArtifactResultOutputWithContext(ctx context.Context) LookupTemplateArtifactResultOutput {
	return o
}

func (o LookupTemplateArtifactResultOutput) DependsOn() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTemplateArtifactResult) []string { return v.DependsOn }).(pulumi.StringArrayOutput)
}

func (o LookupTemplateArtifactResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateArtifactResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupTemplateArtifactResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateArtifactResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupTemplateArtifactResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateArtifactResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTemplateArtifactResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateArtifactResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupTemplateArtifactResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateArtifactResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTemplateArtifactResultOutput) Parameters() ParameterValueBaseResponseMapOutput {
	return o.ApplyT(func(v LookupTemplateArtifactResult) map[string]ParameterValueBaseResponse { return v.Parameters }).(ParameterValueBaseResponseMapOutput)
}

func (o LookupTemplateArtifactResultOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTemplateArtifactResult) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o LookupTemplateArtifactResultOutput) Template() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTemplateArtifactResult) interface{} { return v.Template }).(pulumi.AnyOutput)
}

func (o LookupTemplateArtifactResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplateArtifactResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTemplateArtifactResultOutput{})
}