// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Web.V20181101
{
    public static class ListWebAppHybridConnectionKeysSlot
    {
        public static Task<ListWebAppHybridConnectionKeysSlotResult> InvokeAsync(ListWebAppHybridConnectionKeysSlotArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListWebAppHybridConnectionKeysSlotResult>("azure-nextgen:web/v20181101:listWebAppHybridConnectionKeysSlot", args ?? new ListWebAppHybridConnectionKeysSlotArgs(), options.WithVersion());
    }


    public sealed class ListWebAppHybridConnectionKeysSlotArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the web app.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The namespace for this hybrid connection.
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// The relay name for this hybrid connection.
        /// </summary>
        [Input("relayName", required: true)]
        public string RelayName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the slot for the web app.
        /// </summary>
        [Input("slot", required: true)]
        public string Slot { get; set; } = null!;

        public ListWebAppHybridConnectionKeysSlotArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListWebAppHybridConnectionKeysSlotResult
    {
        /// <summary>
        /// Kind of resource.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// Resource Name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of the send key.
        /// </summary>
        public readonly string SendKeyName;
        /// <summary>
        /// The value of the send key.
        /// </summary>
        public readonly string SendKeyValue;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ListWebAppHybridConnectionKeysSlotResult(
            string? kind,

            string name,

            string sendKeyName,

            string sendKeyValue,

            string type)
        {
            Kind = kind;
            Name = name;
            SendKeyName = sendKeyName;
            SendKeyValue = sendKeyValue;
            Type = type;
        }
    }
}
