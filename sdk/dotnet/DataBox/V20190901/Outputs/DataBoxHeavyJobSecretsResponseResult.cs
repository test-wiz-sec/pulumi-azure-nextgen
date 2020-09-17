// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataBox.V20190901.Outputs
{

    [OutputType]
    public sealed class DataBoxHeavyJobSecretsResponseResult
    {
        /// <summary>
        /// Contains the list of secret objects for a databox heavy job.
        /// </summary>
        public readonly ImmutableArray<Outputs.DataBoxHeavySecretResponseResult> CabinetPodSecrets;
        /// <summary>
        /// Dc Access Security Code for Customer Managed Shipping
        /// </summary>
        public readonly Outputs.DcAccessSecurityCodeResponseResult? DcAccessSecurityCode;
        /// <summary>
        /// Used to indicate what type of job secrets object.
        /// </summary>
        public readonly string JobSecretsType;

        [OutputConstructor]
        private DataBoxHeavyJobSecretsResponseResult(
            ImmutableArray<Outputs.DataBoxHeavySecretResponseResult> cabinetPodSecrets,

            Outputs.DcAccessSecurityCodeResponseResult? dcAccessSecurityCode,

            string jobSecretsType)
        {
            CabinetPodSecrets = cabinetPodSecrets;
            DcAccessSecurityCode = dcAccessSecurityCode;
            JobSecretsType = jobSecretsType;
        }
    }
}