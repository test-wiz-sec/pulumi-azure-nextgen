// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Devices.V20200901Preview.Outputs
{

    [OutputType]
    public sealed class KeyVaultKeyPropertiesResponse
    {
        /// <summary>
        /// The identifier of the key.
        /// </summary>
        public readonly string? KeyIdentifier;

        [OutputConstructor]
        private KeyVaultKeyPropertiesResponse(string? keyIdentifier)
        {
            KeyIdentifier = keyIdentifier;
        }
    }
}
