// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.CognitiveServices.V20170418.Inputs
{

    /// <summary>
    /// The api properties for special APIs.
    /// </summary>
    public sealed class CognitiveServicesAccountApiPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Metrics Advisor Only) The Azure AD Client Id (Application Id).
        /// </summary>
        [Input("aadClientId")]
        public Input<string>? AadClientId { get; set; }

        /// <summary>
        /// (Metrics Advisor Only) The Azure AD Tenant Id.
        /// </summary>
        [Input("aadTenantId")]
        public Input<string>? AadTenantId { get; set; }

        /// <summary>
        /// (Personalization Only) The flag to enable statistics of Bing Search.
        /// </summary>
        [Input("eventHubConnectionString")]
        public Input<string>? EventHubConnectionString { get; set; }

        /// <summary>
        /// (QnAMaker Only) The runtime endpoint of QnAMaker.
        /// </summary>
        [Input("qnaRuntimeEndpoint")]
        public Input<string>? QnaRuntimeEndpoint { get; set; }

        /// <summary>
        /// (Bing Search Only) The flag to enable statistics of Bing Search.
        /// </summary>
        [Input("statisticsEnabled")]
        public Input<bool>? StatisticsEnabled { get; set; }

        /// <summary>
        /// (Personalization Only) The storage account connection string.
        /// </summary>
        [Input("storageAccountConnectionString")]
        public Input<string>? StorageAccountConnectionString { get; set; }

        /// <summary>
        /// (Metrics Advisor Only) The super user of Metrics Advisor.
        /// </summary>
        [Input("superUser")]
        public Input<string>? SuperUser { get; set; }

        /// <summary>
        /// (Metrics Advisor Only) The website name of Metrics Advisor.
        /// </summary>
        [Input("websiteName")]
        public Input<string>? WebsiteName { get; set; }

        public CognitiveServicesAccountApiPropertiesArgs()
        {
        }
    }
}
