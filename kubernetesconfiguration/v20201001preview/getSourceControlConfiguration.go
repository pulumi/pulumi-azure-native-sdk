// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The SourceControl Configuration object returned in Get & Put response.
func LookupSourceControlConfiguration(ctx *pulumi.Context, args *LookupSourceControlConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupSourceControlConfigurationResult, error) {
	var rv LookupSourceControlConfigurationResult
	err := ctx.Invoke("azure-native:kubernetesconfiguration/v20201001preview:getSourceControlConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSourceControlConfigurationArgs struct {
	ClusterName                    string `pulumi:"clusterName"`
	ClusterResourceName            string `pulumi:"clusterResourceName"`
	ClusterRp                      string `pulumi:"clusterRp"`
	ResourceGroupName              string `pulumi:"resourceGroupName"`
	SourceControlConfigurationName string `pulumi:"sourceControlConfigurationName"`
}

// The SourceControl Configuration object returned in Get & Put response.
type LookupSourceControlConfigurationResult struct {
	ComplianceStatus               ComplianceStatusResponse        `pulumi:"complianceStatus"`
	ConfigurationProtectedSettings map[string]string               `pulumi:"configurationProtectedSettings"`
	EnableHelmOperator             *bool                           `pulumi:"enableHelmOperator"`
	HelmOperatorProperties         *HelmOperatorPropertiesResponse `pulumi:"helmOperatorProperties"`
	Id                             string                          `pulumi:"id"`
	Name                           string                          `pulumi:"name"`
	OperatorInstanceName           *string                         `pulumi:"operatorInstanceName"`
	OperatorNamespace              *string                         `pulumi:"operatorNamespace"`
	OperatorParams                 *string                         `pulumi:"operatorParams"`
	OperatorScope                  *string                         `pulumi:"operatorScope"`
	OperatorType                   *string                         `pulumi:"operatorType"`
	ProvisioningState              string                          `pulumi:"provisioningState"`
	RepositoryPublicKey            string                          `pulumi:"repositoryPublicKey"`
	RepositoryUrl                  *string                         `pulumi:"repositoryUrl"`
	SshKnownHostsContents          *string                         `pulumi:"sshKnownHostsContents"`
	SystemData                     *SystemDataResponse             `pulumi:"systemData"`
	Type                           string                          `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupSourceControlConfigurationResult
func (val *LookupSourceControlConfigurationResult) Defaults() *LookupSourceControlConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.OperatorNamespace) {
		operatorNamespace_ := "default"
		tmp.OperatorNamespace = &operatorNamespace_
	}
	return &tmp
}

func LookupSourceControlConfigurationOutput(ctx *pulumi.Context, args LookupSourceControlConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupSourceControlConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceControlConfigurationResult, error) {
			args := v.(LookupSourceControlConfigurationArgs)
			r, err := LookupSourceControlConfiguration(ctx, &args, opts...)
			var s LookupSourceControlConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceControlConfigurationResultOutput)
}

type LookupSourceControlConfigurationOutputArgs struct {
	ClusterName                    pulumi.StringInput `pulumi:"clusterName"`
	ClusterResourceName            pulumi.StringInput `pulumi:"clusterResourceName"`
	ClusterRp                      pulumi.StringInput `pulumi:"clusterRp"`
	ResourceGroupName              pulumi.StringInput `pulumi:"resourceGroupName"`
	SourceControlConfigurationName pulumi.StringInput `pulumi:"sourceControlConfigurationName"`
}

func (LookupSourceControlConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceControlConfigurationArgs)(nil)).Elem()
}

// The SourceControl Configuration object returned in Get & Put response.
type LookupSourceControlConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupSourceControlConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceControlConfigurationResult)(nil)).Elem()
}

func (o LookupSourceControlConfigurationResultOutput) ToLookupSourceControlConfigurationResultOutput() LookupSourceControlConfigurationResultOutput {
	return o
}

func (o LookupSourceControlConfigurationResultOutput) ToLookupSourceControlConfigurationResultOutputWithContext(ctx context.Context) LookupSourceControlConfigurationResultOutput {
	return o
}

func (o LookupSourceControlConfigurationResultOutput) ComplianceStatus() ComplianceStatusResponseOutput {
	return o.ApplyT(func(v LookupSourceControlConfigurationResult) ComplianceStatusResponse { return v.ComplianceStatus }).(ComplianceStatusResponseOutput)
}

func (o LookupSourceControlConfigurationResultOutput) ConfigurationProtectedSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSourceControlConfigurationResult) map[string]string {
		return v.ConfigurationProtectedSettings
	}).(pulumi.StringMapOutput)
}

func (o LookupSourceControlConfigurationResultOutput) EnableHelmOperator() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSourceControlConfigurationResult) *bool { return v.EnableHelmOperator }).(pulumi.BoolPtrOutput)
}

func (o LookupSourceControlConfigurationResultOutput) HelmOperatorProperties() HelmOperatorPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupSourceControlConfigurationResult) *HelmOperatorPropertiesResponse {
		return v.HelmOperatorProperties
	}).(HelmOperatorPropertiesResponsePtrOutput)
}

func (o LookupSourceControlConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSourceControlConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSourceControlConfigurationResultOutput) OperatorInstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlConfigurationResult) *string { return v.OperatorInstanceName }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlConfigurationResultOutput) OperatorNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlConfigurationResult) *string { return v.OperatorNamespace }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlConfigurationResultOutput) OperatorParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlConfigurationResult) *string { return v.OperatorParams }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlConfigurationResultOutput) OperatorScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlConfigurationResult) *string { return v.OperatorScope }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlConfigurationResultOutput) OperatorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlConfigurationResult) *string { return v.OperatorType }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlConfigurationResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlConfigurationResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSourceControlConfigurationResultOutput) RepositoryPublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlConfigurationResult) string { return v.RepositoryPublicKey }).(pulumi.StringOutput)
}

func (o LookupSourceControlConfigurationResultOutput) RepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlConfigurationResult) *string { return v.RepositoryUrl }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlConfigurationResultOutput) SshKnownHostsContents() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceControlConfigurationResult) *string { return v.SshKnownHostsContents }).(pulumi.StringPtrOutput)
}

func (o LookupSourceControlConfigurationResultOutput) SystemData() SystemDataResponsePtrOutput {
	return o.ApplyT(func(v LookupSourceControlConfigurationResult) *SystemDataResponse { return v.SystemData }).(SystemDataResponsePtrOutput)
}

func (o LookupSourceControlConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSourceControlConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceControlConfigurationResultOutput{})
}