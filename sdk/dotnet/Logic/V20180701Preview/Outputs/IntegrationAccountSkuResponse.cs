// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Logic.V20180701Preview.Outputs
{

    [OutputType]
    public sealed class IntegrationAccountSkuResponse
    {
        /// <summary>
        /// The sku name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private IntegrationAccountSkuResponse(string name)
        {
            Name = name;
        }
    }
}