// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Databricks.V20180401.Outputs
{

    [OutputType]
    public sealed class WorkspaceEncryptionParameterResponse
    {
        /// <summary>
        /// The type of variable that this is
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// The value which should be used for this field.
        /// </summary>
        public readonly Outputs.EncryptionResponse? Value;

        [OutputConstructor]
        private WorkspaceEncryptionParameterResponse(
            string? type,

            Outputs.EncryptionResponse? value)
        {
            Type = type;
            Value = value;
        }
    }
}