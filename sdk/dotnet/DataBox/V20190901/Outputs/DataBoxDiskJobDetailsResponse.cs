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
    public sealed class DataBoxDiskJobDetailsResponse
    {
        /// <summary>
        /// Shared access key to download the chain of custody logs
        /// </summary>
        public readonly string ChainOfCustodySasKey;
        /// <summary>
        /// Contact details for notification and shipping.
        /// </summary>
        public readonly Outputs.ContactDetailsResponse ContactDetails;
        /// <summary>
        /// List of copy log details.
        /// </summary>
        public readonly ImmutableArray<Union<Outputs.DataBoxAccountCopyLogDetailsResponse, Union<Outputs.DataBoxDiskCopyLogDetailsResponse, Outputs.DataBoxHeavyAccountCopyLogDetailsResponse>>> CopyLogDetails;
        /// <summary>
        /// Copy progress per disk.
        /// </summary>
        public readonly ImmutableArray<Outputs.DataBoxDiskCopyProgressResponse> CopyProgress;
        /// <summary>
        /// Delivery package shipping details.
        /// </summary>
        public readonly Outputs.PackageShippingDetailsResponse DeliveryPackage;
        /// <summary>
        /// Destination account details.
        /// </summary>
        public readonly ImmutableArray<Union<Outputs.DestinationManagedDiskDetailsResponse, Outputs.DestinationStorageAccountDetailsResponse>> DestinationAccountDetails;
        /// <summary>
        /// Contains the map of disk serial number to the disk size being used for the job. Is returned only after the disks are shipped to the customer.
        /// </summary>
        public readonly ImmutableDictionary<string, int> DisksAndSizeDetails;
        /// <summary>
        /// Error details for failure. This is optional.
        /// </summary>
        public readonly ImmutableArray<Outputs.JobErrorDetailsResponse> ErrorDetails;
        /// <summary>
        /// The expected size of the data, which needs to be transferred in this job, in terabytes.
        /// </summary>
        public readonly int? ExpectedDataSizeInTerabytes;
        /// <summary>
        /// Indicates the type of job details.
        /// </summary>
        public readonly string JobDetailsType;
        /// <summary>
        /// List of stages that run in the job.
        /// </summary>
        public readonly ImmutableArray<Outputs.JobStagesResponse> JobStages;
        /// <summary>
        /// User entered passkey for DataBox Disk job.
        /// </summary>
        public readonly string? Passkey;
        /// <summary>
        /// Preferences for the order.
        /// </summary>
        public readonly Outputs.PreferencesResponse? Preferences;
        /// <summary>
        /// User preference on what size disks are needed for the job. The map is from the disk size in TB to the count. Eg. {2,5} means 5 disks of 2 TB size. Key is string but will be checked against an int.
        /// </summary>
        public readonly ImmutableDictionary<string, int>? PreferredDisks;
        /// <summary>
        /// Return package shipping details.
        /// </summary>
        public readonly Outputs.PackageShippingDetailsResponse ReturnPackage;
        /// <summary>
        /// Shared access key to download the return shipment label
        /// </summary>
        public readonly string ReverseShipmentLabelSasKey;
        /// <summary>
        /// Shipping address of the customer.
        /// </summary>
        public readonly Outputs.ShippingAddressResponse ShippingAddress;

        [OutputConstructor]
        private DataBoxDiskJobDetailsResponse(
            string chainOfCustodySasKey,

            Outputs.ContactDetailsResponse contactDetails,

            ImmutableArray<Union<Outputs.DataBoxAccountCopyLogDetailsResponse, Union<Outputs.DataBoxDiskCopyLogDetailsResponse, Outputs.DataBoxHeavyAccountCopyLogDetailsResponse>>> copyLogDetails,

            ImmutableArray<Outputs.DataBoxDiskCopyProgressResponse> copyProgress,

            Outputs.PackageShippingDetailsResponse deliveryPackage,

            ImmutableArray<Union<Outputs.DestinationManagedDiskDetailsResponse, Outputs.DestinationStorageAccountDetailsResponse>> destinationAccountDetails,

            ImmutableDictionary<string, int> disksAndSizeDetails,

            ImmutableArray<Outputs.JobErrorDetailsResponse> errorDetails,

            int? expectedDataSizeInTerabytes,

            string jobDetailsType,

            ImmutableArray<Outputs.JobStagesResponse> jobStages,

            string? passkey,

            Outputs.PreferencesResponse? preferences,

            ImmutableDictionary<string, int>? preferredDisks,

            Outputs.PackageShippingDetailsResponse returnPackage,

            string reverseShipmentLabelSasKey,

            Outputs.ShippingAddressResponse shippingAddress)
        {
            ChainOfCustodySasKey = chainOfCustodySasKey;
            ContactDetails = contactDetails;
            CopyLogDetails = copyLogDetails;
            CopyProgress = copyProgress;
            DeliveryPackage = deliveryPackage;
            DestinationAccountDetails = destinationAccountDetails;
            DisksAndSizeDetails = disksAndSizeDetails;
            ErrorDetails = errorDetails;
            ExpectedDataSizeInTerabytes = expectedDataSizeInTerabytes;
            JobDetailsType = jobDetailsType;
            JobStages = jobStages;
            Passkey = passkey;
            Preferences = preferences;
            PreferredDisks = preferredDisks;
            ReturnPackage = returnPackage;
            ReverseShipmentLabelSasKey = reverseShipmentLabelSasKey;
            ShippingAddress = shippingAddress;
        }
    }
}