// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the on-premises integration runtime connection information for encrypting the on-premises data source credentials.
//
// Uses Azure REST API version 2018-06-01.
func GetIntegrationRuntimeConnectionInfo(ctx *pulumi.Context, args *GetIntegrationRuntimeConnectionInfoArgs, opts ...pulumi.InvokeOption) (*GetIntegrationRuntimeConnectionInfoResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetIntegrationRuntimeConnectionInfoResult
	err := ctx.Invoke("azure-native:datafactory:getIntegrationRuntimeConnectionInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetIntegrationRuntimeConnectionInfoArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The integration runtime name.
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Connection information for encrypting the on-premises data source credentials.
type GetIntegrationRuntimeConnectionInfoResult struct {
	// The on-premises integration runtime host URL.
	HostServiceUri string `pulumi:"hostServiceUri"`
	// The integration runtime SSL certificate thumbprint. Click-Once application uses it to do server validation.
	IdentityCertThumbprint string `pulumi:"identityCertThumbprint"`
	// Whether the identity certificate is expired.
	IsIdentityCertExprired bool `pulumi:"isIdentityCertExprired"`
	// The public key for encrypting a credential when transferring the credential to the integration runtime.
	PublicKey string `pulumi:"publicKey"`
	// The token generated in service. Callers use this token to authenticate to integration runtime.
	ServiceToken string `pulumi:"serviceToken"`
	// The integration runtime version.
	Version string `pulumi:"version"`
}

func GetIntegrationRuntimeConnectionInfoOutput(ctx *pulumi.Context, args GetIntegrationRuntimeConnectionInfoOutputArgs, opts ...pulumi.InvokeOption) GetIntegrationRuntimeConnectionInfoResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetIntegrationRuntimeConnectionInfoResultOutput, error) {
			args := v.(GetIntegrationRuntimeConnectionInfoArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:datafactory:getIntegrationRuntimeConnectionInfo", args, GetIntegrationRuntimeConnectionInfoResultOutput{}, options).(GetIntegrationRuntimeConnectionInfoResultOutput), nil
		}).(GetIntegrationRuntimeConnectionInfoResultOutput)
}

type GetIntegrationRuntimeConnectionInfoOutputArgs struct {
	// The factory name.
	FactoryName pulumi.StringInput `pulumi:"factoryName"`
	// The integration runtime name.
	IntegrationRuntimeName pulumi.StringInput `pulumi:"integrationRuntimeName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetIntegrationRuntimeConnectionInfoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIntegrationRuntimeConnectionInfoArgs)(nil)).Elem()
}

// Connection information for encrypting the on-premises data source credentials.
type GetIntegrationRuntimeConnectionInfoResultOutput struct{ *pulumi.OutputState }

func (GetIntegrationRuntimeConnectionInfoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIntegrationRuntimeConnectionInfoResult)(nil)).Elem()
}

func (o GetIntegrationRuntimeConnectionInfoResultOutput) ToGetIntegrationRuntimeConnectionInfoResultOutput() GetIntegrationRuntimeConnectionInfoResultOutput {
	return o
}

func (o GetIntegrationRuntimeConnectionInfoResultOutput) ToGetIntegrationRuntimeConnectionInfoResultOutputWithContext(ctx context.Context) GetIntegrationRuntimeConnectionInfoResultOutput {
	return o
}

// The on-premises integration runtime host URL.
func (o GetIntegrationRuntimeConnectionInfoResultOutput) HostServiceUri() pulumi.StringOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeConnectionInfoResult) string { return v.HostServiceUri }).(pulumi.StringOutput)
}

// The integration runtime SSL certificate thumbprint. Click-Once application uses it to do server validation.
func (o GetIntegrationRuntimeConnectionInfoResultOutput) IdentityCertThumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeConnectionInfoResult) string { return v.IdentityCertThumbprint }).(pulumi.StringOutput)
}

// Whether the identity certificate is expired.
func (o GetIntegrationRuntimeConnectionInfoResultOutput) IsIdentityCertExprired() pulumi.BoolOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeConnectionInfoResult) bool { return v.IsIdentityCertExprired }).(pulumi.BoolOutput)
}

// The public key for encrypting a credential when transferring the credential to the integration runtime.
func (o GetIntegrationRuntimeConnectionInfoResultOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeConnectionInfoResult) string { return v.PublicKey }).(pulumi.StringOutput)
}

// The token generated in service. Callers use this token to authenticate to integration runtime.
func (o GetIntegrationRuntimeConnectionInfoResultOutput) ServiceToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeConnectionInfoResult) string { return v.ServiceToken }).(pulumi.StringOutput)
}

// The integration runtime version.
func (o GetIntegrationRuntimeConnectionInfoResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetIntegrationRuntimeConnectionInfoResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIntegrationRuntimeConnectionInfoResultOutput{})
}
