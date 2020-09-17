// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DBforPostgreSQL.V20171201Preview
{
    public static class GetServerAdministrator
    {
        public static Task<GetServerAdministratorResult> InvokeAsync(GetServerAdministratorArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerAdministratorResult>("azure-nextgen:dbforpostgresql/v20171201preview:getServerAdministrator", args ?? new GetServerAdministratorArgs(), options.WithVersion());
    }


    public sealed class GetServerAdministratorArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        public GetServerAdministratorArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServerAdministratorResult
    {
        /// <summary>
        /// The type of administrator.
        /// </summary>
        public readonly string AdministratorType;
        /// <summary>
        /// The server administrator login account name.
        /// </summary>
        public readonly string Login;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The server administrator Sid (Secure ID).
        /// </summary>
        public readonly string Sid;
        /// <summary>
        /// The server Active Directory Administrator tenant id.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetServerAdministratorResult(
            string administratorType,

            string login,

            string name,

            string sid,

            string tenantId,

            string type)
        {
            AdministratorType = administratorType;
            Login = login;
            Name = name;
            Sid = sid;
            TenantId = tenantId;
            Type = type;
        }
    }
}
