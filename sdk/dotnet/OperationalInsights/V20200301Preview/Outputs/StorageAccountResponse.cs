// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.OperationalInsights.V20200301Preview.Outputs
{

    [OutputType]
    public sealed class StorageAccountResponse
    {
        /// <summary>
        /// The Azure Resource Manager ID of the storage account resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The storage account key.
        /// </summary>
        public readonly string Key;

        [OutputConstructor]
        private StorageAccountResponse(
            string id,

            string key)
        {
            Id = id;
            Key = key;
        }
    }
}