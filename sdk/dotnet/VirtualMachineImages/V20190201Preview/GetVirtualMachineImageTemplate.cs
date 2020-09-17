// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.VirtualMachineImages.V20190201Preview
{
    public static class GetVirtualMachineImageTemplate
    {
        public static Task<GetVirtualMachineImageTemplateResult> InvokeAsync(GetVirtualMachineImageTemplateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualMachineImageTemplateResult>("azure-nextgen:virtualmachineimages/v20190201preview:getVirtualMachineImageTemplate", args ?? new GetVirtualMachineImageTemplateArgs(), options.WithVersion());
    }


    public sealed class GetVirtualMachineImageTemplateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the image Template
        /// </summary>
        [Input("imageTemplateName", required: true)]
        public string ImageTemplateName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetVirtualMachineImageTemplateArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVirtualMachineImageTemplateResult
    {
        /// <summary>
        /// Specifies the properties used to describe the customization steps of the image, like Image source etc
        /// </summary>
        public readonly ImmutableArray<Union<Outputs.ImageTemplatePowerShellCustomizerResponse, Union<Outputs.ImageTemplateRestartCustomizerResponse, Outputs.ImageTemplateShellCustomizerResponse>>> Customize;
        /// <summary>
        /// The distribution targets where the image output needs to go to.
        /// </summary>
        public readonly ImmutableArray<Union<Outputs.ImageTemplateManagedImageDistributorResponse, Union<Outputs.ImageTemplateSharedImageDistributorResponse, Outputs.ImageTemplateVhdDistributorResponse>>> Distribute;
        /// <summary>
        /// State of 'run' that is currently executing or was last executed.
        /// </summary>
        public readonly Outputs.ImageTemplateLastRunStatusResponse LastRunStatus;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provisioning error, if any
        /// </summary>
        public readonly Outputs.ProvisioningErrorResponse ProvisioningError;
        /// <summary>
        /// Provisioning state of the resource
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Specifies the properties used to describe the source image.
        /// </summary>
        public readonly Union<Outputs.ImageTemplateIsoSourceResponse, Union<Outputs.ImageTemplateManagedImageSourceResponse, Outputs.ImageTemplatePlatformImageSourceResponse>> Source;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetVirtualMachineImageTemplateResult(
            ImmutableArray<Union<Outputs.ImageTemplatePowerShellCustomizerResponse, Union<Outputs.ImageTemplateRestartCustomizerResponse, Outputs.ImageTemplateShellCustomizerResponse>>> customize,

            ImmutableArray<Union<Outputs.ImageTemplateManagedImageDistributorResponse, Union<Outputs.ImageTemplateSharedImageDistributorResponse, Outputs.ImageTemplateVhdDistributorResponse>>> distribute,

            Outputs.ImageTemplateLastRunStatusResponse lastRunStatus,

            string location,

            string name,

            Outputs.ProvisioningErrorResponse provisioningError,

            string provisioningState,

            Union<Outputs.ImageTemplateIsoSourceResponse, Union<Outputs.ImageTemplateManagedImageSourceResponse, Outputs.ImageTemplatePlatformImageSourceResponse>> source,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Customize = customize;
            Distribute = distribute;
            LastRunStatus = lastRunStatus;
            Location = location;
            Name = name;
            ProvisioningError = provisioningError;
            ProvisioningState = provisioningState;
            Source = source;
            Tags = tags;
            Type = type;
        }
    }
}