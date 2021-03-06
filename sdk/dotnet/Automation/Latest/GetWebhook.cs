// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Automation.Latest
{
    public static class GetWebhook
    {
        public static Task<GetWebhookResult> InvokeAsync(GetWebhookArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWebhookResult>("azure-nextgen:automation/latest:getWebhook", args ?? new GetWebhookArgs(), options.WithVersion());
    }


    public sealed class GetWebhookArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public string AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The webhook name.
        /// </summary>
        [Input("webhookName", required: true)]
        public string WebhookName { get; set; } = null!;

        public GetWebhookArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWebhookResult
    {
        /// <summary>
        /// Gets or sets the creation time.
        /// </summary>
        public readonly string? CreationTime;
        /// <summary>
        /// Gets or sets the description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Gets or sets the expiry time.
        /// </summary>
        public readonly string? ExpiryTime;
        /// <summary>
        /// Gets or sets the value of the enabled flag of the webhook.
        /// </summary>
        public readonly bool? IsEnabled;
        /// <summary>
        /// Gets or sets the last invoked time.
        /// </summary>
        public readonly string? LastInvokedTime;
        /// <summary>
        /// Details of the user who last modified the Webhook
        /// </summary>
        public readonly string? LastModifiedBy;
        /// <summary>
        /// Gets or sets the last modified time.
        /// </summary>
        public readonly string? LastModifiedTime;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets or sets the parameters of the job that is created when the webhook calls the runbook it is associated with.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Parameters;
        /// <summary>
        /// Gets or sets the name of the hybrid worker group the webhook job will run on.
        /// </summary>
        public readonly string? RunOn;
        /// <summary>
        /// Gets or sets the runbook the webhook is associated with.
        /// </summary>
        public readonly Outputs.RunbookAssociationPropertyResponse? Runbook;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Gets or sets the webhook uri.
        /// </summary>
        public readonly string? Uri;

        [OutputConstructor]
        private GetWebhookResult(
            string? creationTime,

            string? description,

            string? expiryTime,

            bool? isEnabled,

            string? lastInvokedTime,

            string? lastModifiedBy,

            string? lastModifiedTime,

            string name,

            ImmutableDictionary<string, string>? parameters,

            string? runOn,

            Outputs.RunbookAssociationPropertyResponse? runbook,

            string type,

            string? uri)
        {
            CreationTime = creationTime;
            Description = description;
            ExpiryTime = expiryTime;
            IsEnabled = isEnabled;
            LastInvokedTime = lastInvokedTime;
            LastModifiedBy = lastModifiedBy;
            LastModifiedTime = lastModifiedTime;
            Name = name;
            Parameters = parameters;
            RunOn = runOn;
            Runbook = runbook;
            Type = type;
            Uri = uri;
        }
    }
}
