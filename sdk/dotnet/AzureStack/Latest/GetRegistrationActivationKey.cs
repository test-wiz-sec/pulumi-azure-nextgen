// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.AzureStack.Latest
{
    public static class GetRegistrationActivationKey
    {
        public static Task<GetRegistrationActivationKeyResult> InvokeAsync(GetRegistrationActivationKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegistrationActivationKeyResult>("azure-nextgen:azurestack/latest:getRegistrationActivationKey", args ?? new GetRegistrationActivationKeyArgs(), options.WithVersion());
    }


    public sealed class GetRegistrationActivationKeyArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Azure Stack registration.
        /// </summary>
        [Input("registrationName", required: true)]
        public string RegistrationName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group.
        /// </summary>
        [Input("resourceGroup", required: true)]
        public string ResourceGroup { get; set; } = null!;

        public GetRegistrationActivationKeyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegistrationActivationKeyResult
    {
        /// <summary>
        /// Azure Stack activation key.
        /// </summary>
        public readonly string? ActivationKey;

        [OutputConstructor]
        private GetRegistrationActivationKeyResult(string? activationKey)
        {
            ActivationKey = activationKey;
        }
    }
}
