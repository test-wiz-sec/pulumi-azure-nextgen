// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerService.V20190430.Outputs
{

    [OutputType]
    public sealed class OpenShiftManagedClusterAuthProfileResponse
    {
        /// <summary>
        /// Type of authentication profile to use.
        /// </summary>
        public readonly ImmutableArray<Outputs.OpenShiftManagedClusterIdentityProviderResponse> IdentityProviders;

        [OutputConstructor]
        private OpenShiftManagedClusterAuthProfileResponse(ImmutableArray<Outputs.OpenShiftManagedClusterIdentityProviderResponse> identityProviders)
        {
            IdentityProviders = identityProviders;
        }
    }
}