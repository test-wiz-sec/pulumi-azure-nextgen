// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataBoxEdge.V20190801
{
    public static class GetUser
    {
        public static Task<GetUserResult> InvokeAsync(GetUserArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("azure-nextgen:databoxedge/v20190801:getUser", args ?? new GetUserArgs(), options.WithVersion());
    }


    public sealed class GetUserArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The device name.
        /// </summary>
        [Input("deviceName", required: true)]
        public string DeviceName { get; set; } = null!;

        /// <summary>
        /// The user name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetUserArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// The password details.
        /// </summary>
        public readonly Outputs.AsymmetricEncryptedSecretResponse? EncryptedPassword;
        /// <summary>
        /// The object name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of shares that the user has rights on. This field should not be specified during user creation.
        /// </summary>
        public readonly ImmutableArray<Outputs.ShareAccessRightResponse> ShareAccessRights;
        /// <summary>
        /// The hierarchical type of the object.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Type of the user.
        /// </summary>
        public readonly string UserType;

        [OutputConstructor]
        private GetUserResult(
            Outputs.AsymmetricEncryptedSecretResponse? encryptedPassword,

            string name,

            ImmutableArray<Outputs.ShareAccessRightResponse> shareAccessRights,

            string type,

            string userType)
        {
            EncryptedPassword = encryptedPassword;
            Name = name;
            ShareAccessRights = shareAccessRights;
            Type = type;
            UserType = userType;
        }
    }
}
