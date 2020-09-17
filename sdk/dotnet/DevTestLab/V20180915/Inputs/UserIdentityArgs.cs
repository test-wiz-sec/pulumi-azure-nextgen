// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DevTestLab.V20180915.Inputs
{

    /// <summary>
    /// Identity attributes of a lab user.
    /// </summary>
    public sealed class UserIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set to the app Id of the client JWT making the request.
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        /// <summary>
        /// Set to the object Id of the client JWT making the request. Not all users have object Id. For CSP (reseller) scenarios for example, object Id is not available.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        /// <summary>
        /// Set to the principal Id of the client JWT making the request. Service principal will not have the principal Id.
        /// </summary>
        [Input("principalId")]
        public Input<string>? PrincipalId { get; set; }

        /// <summary>
        /// Set to the principal name / UPN of the client JWT making the request.
        /// </summary>
        [Input("principalName")]
        public Input<string>? PrincipalName { get; set; }

        /// <summary>
        /// Set to the tenant ID of the client JWT making the request.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public UserIdentityArgs()
        {
        }
    }
}