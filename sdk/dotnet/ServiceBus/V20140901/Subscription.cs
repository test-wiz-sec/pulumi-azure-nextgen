// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ServiceBus.V20140901
{
    /// <summary>
    /// Description of subscription resource.
    /// </summary>
    public partial class Subscription : Pulumi.CustomResource
    {
        /// <summary>
        /// Last time there was a receive request to this subscription.
        /// </summary>
        [Output("accessedAt")]
        public Output<string> AccessedAt { get; private set; } = null!;

        /// <summary>
        /// TimeSpan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
        /// </summary>
        [Output("autoDeleteOnIdle")]
        public Output<string?> AutoDeleteOnIdle { get; private set; } = null!;

        /// <summary>
        /// Message Count Details.
        /// </summary>
        [Output("countDetails")]
        public Output<Outputs.MessageCountDetailsResponse> CountDetails { get; private set; } = null!;

        /// <summary>
        /// Exact time the message was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Value that indicates whether a subscription has dead letter support on filter evaluation exceptions.
        /// </summary>
        [Output("deadLetteringOnFilterEvaluationExceptions")]
        public Output<bool?> DeadLetteringOnFilterEvaluationExceptions { get; private set; } = null!;

        /// <summary>
        /// Value that indicates whether a subscription has dead letter support when a message expires.
        /// </summary>
        [Output("deadLetteringOnMessageExpiration")]
        public Output<bool?> DeadLetteringOnMessageExpiration { get; private set; } = null!;

        /// <summary>
        /// Default message time to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
        /// </summary>
        [Output("defaultMessageTimeToLive")]
        public Output<string?> DefaultMessageTimeToLive { get; private set; } = null!;

        /// <summary>
        /// Value that indicates whether server-side batched operations are enabled.
        /// </summary>
        [Output("enableBatchedOperations")]
        public Output<bool?> EnableBatchedOperations { get; private set; } = null!;

        /// <summary>
        /// Entity availability status for the topic.
        /// </summary>
        [Output("entityAvailabilityStatus")]
        public Output<string?> EntityAvailabilityStatus { get; private set; } = null!;

        /// <summary>
        /// Value that indicates whether the entity description is read-only.
        /// </summary>
        [Output("isReadOnly")]
        public Output<bool?> IsReadOnly { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The lock duration time span for the subscription.
        /// </summary>
        [Output("lockDuration")]
        public Output<string?> LockDuration { get; private set; } = null!;

        /// <summary>
        /// Number of maximum deliveries.
        /// </summary>
        [Output("maxDeliveryCount")]
        public Output<int?> MaxDeliveryCount { get; private set; } = null!;

        /// <summary>
        /// Number of messages.
        /// </summary>
        [Output("messageCount")]
        public Output<int> MessageCount { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Value indicating if a subscription supports the concept of sessions.
        /// </summary>
        [Output("requiresSession")]
        public Output<bool?> RequiresSession { get; private set; } = null!;

        /// <summary>
        /// Enumerates the possible values for the status of a messaging entity.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The exact time the message was updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a Subscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subscription(string name, SubscriptionArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:servicebus/v20140901:Subscription", name, args ?? new SubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Subscription(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:servicebus/v20140901:Subscription", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:servicebus/latest:Subscription"},
                    new Pulumi.Alias { Type = "azure-nextgen:servicebus/v20150801:Subscription"},
                    new Pulumi.Alias { Type = "azure-nextgen:servicebus/v20170401:Subscription"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Subscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subscription Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Subscription(name, id, options);
        }
    }

    public sealed class SubscriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// TimeSpan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
        /// </summary>
        [Input("autoDeleteOnIdle")]
        public Input<string>? AutoDeleteOnIdle { get; set; }

        /// <summary>
        /// Value that indicates whether a subscription has dead letter support on filter evaluation exceptions.
        /// </summary>
        [Input("deadLetteringOnFilterEvaluationExceptions")]
        public Input<bool>? DeadLetteringOnFilterEvaluationExceptions { get; set; }

        /// <summary>
        /// Value that indicates whether a subscription has dead letter support when a message expires.
        /// </summary>
        [Input("deadLetteringOnMessageExpiration")]
        public Input<bool>? DeadLetteringOnMessageExpiration { get; set; }

        /// <summary>
        /// Default message time to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
        /// </summary>
        [Input("defaultMessageTimeToLive")]
        public Input<string>? DefaultMessageTimeToLive { get; set; }

        /// <summary>
        /// Value that indicates whether server-side batched operations are enabled.
        /// </summary>
        [Input("enableBatchedOperations")]
        public Input<bool>? EnableBatchedOperations { get; set; }

        /// <summary>
        /// Entity availability status for the topic.
        /// </summary>
        [Input("entityAvailabilityStatus")]
        public Input<string>? EntityAvailabilityStatus { get; set; }

        /// <summary>
        /// Value that indicates whether the entity description is read-only.
        /// </summary>
        [Input("isReadOnly")]
        public Input<bool>? IsReadOnly { get; set; }

        /// <summary>
        /// Subscription data center location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The lock duration time span for the subscription.
        /// </summary>
        [Input("lockDuration")]
        public Input<string>? LockDuration { get; set; }

        /// <summary>
        /// Number of maximum deliveries.
        /// </summary>
        [Input("maxDeliveryCount")]
        public Input<int>? MaxDeliveryCount { get; set; }

        /// <summary>
        /// The namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Value indicating if a subscription supports the concept of sessions.
        /// </summary>
        [Input("requiresSession")]
        public Input<bool>? RequiresSession { get; set; }

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Enumerates the possible values for the status of a messaging entity.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The subscription name.
        /// </summary>
        [Input("subscriptionName", required: true)]
        public Input<string> SubscriptionName { get; set; } = null!;

        /// <summary>
        /// The topic name.
        /// </summary>
        [Input("topicName", required: true)]
        public Input<string> TopicName { get; set; } = null!;

        /// <summary>
        /// Resource manager type of the resource.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SubscriptionArgs()
        {
        }
    }
}