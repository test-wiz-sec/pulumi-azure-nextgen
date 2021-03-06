// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20140401

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	var rv LookupDatabaseResult
	err := ctx.Invoke("azure-nextgen:sql/v20140401:getDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseArgs struct {
	// The name of the database to be retrieved.
	DatabaseName string `pulumi:"databaseName"`
	// A comma separated list of child objects to expand in the response. Possible properties: serviceTierAdvisors, transparentDataEncryption.
	Expand *string `pulumi:"expand"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
}

// Represents a database.
type LookupDatabaseResult struct {
	// The collation of the database. If createMode is not Default, this value is ignored.
	Collation *string `pulumi:"collation"`
	// The containment state of the database.
	ContainmentState int `pulumi:"containmentState"`
	// Specifies the mode of database creation.
	//
	// Default: regular database creation.
	//
	// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
	//
	// OnlineSecondary/NonReadableSecondary: creates a database as a (readable or nonreadable) secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
	//
	// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
	//
	// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
	//
	// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
	//
	// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
	//
	// Copy, NonReadableSecondary, OnlineSecondary and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
	CreateMode *string `pulumi:"createMode"`
	// The creation date of the database (ISO8601 format).
	CreationDate string `pulumi:"creationDate"`
	// The current service level objective ID of the database. This is the ID of the service level objective that is currently active.
	CurrentServiceObjectiveId string `pulumi:"currentServiceObjectiveId"`
	// The ID of the database.
	DatabaseId string `pulumi:"databaseId"`
	// The default secondary region for this database.
	DefaultSecondaryLocation string `pulumi:"defaultSecondaryLocation"`
	// This records the earliest start date and time that restore is available for this database (ISO8601 format).
	EarliestRestoreDate string `pulumi:"earliestRestoreDate"`
	// The edition of the database. The DatabaseEditions enumeration contains all the valid editions. If createMode is NonReadableSecondary or OnlineSecondary, this value is ignored.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	Edition *string `pulumi:"edition"`
	// The name of the elastic pool the database is in. If elasticPoolName and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveName is ignored. Not supported for DataWarehouse edition.
	ElasticPoolName *string `pulumi:"elasticPoolName"`
	// The resource identifier of the failover group containing this database.
	FailoverGroupId string `pulumi:"failoverGroupId"`
	// Kind of database.  This is metadata used for the Azure portal experience.
	Kind string `pulumi:"kind"`
	// Resource location.
	Location string `pulumi:"location"`
	// The max size of the database expressed in bytes. If createMode is not Default, this value is ignored. To see possible values, query the capabilities API (/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationID}/capabilities) referred to by operationId: "Capabilities_ListByLocation."
	MaxSizeBytes *string `pulumi:"maxSizeBytes"`
	// Resource name.
	Name string `pulumi:"name"`
	// Conditional. If the database is a geo-secondary, readScale indicates whether read-only connections are allowed to this database or not. Not supported for DataWarehouse edition.
	ReadScale *string `pulumi:"readScale"`
	// The recommended indices for this database.
	RecommendedIndex []RecommendedIndexResponse `pulumi:"recommendedIndex"`
	// Conditional. If createMode is RestoreLongTermRetentionBackup, then this value is required. Specifies the resource ID of the recovery point to restore from.
	RecoveryServicesRecoveryPointResourceId *string `pulumi:"recoveryServicesRecoveryPointResourceId"`
	// The configured service level objective ID of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of currentServiceObjectiveId property. If requestedServiceObjectiveId and requestedServiceObjectiveName are both updated, the value of requestedServiceObjectiveId overrides the value of requestedServiceObjectiveName.
	//
	// The list of SKUs may vary by region and support offer. To determine the service objective ids that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API.
	RequestedServiceObjectiveId *string `pulumi:"requestedServiceObjectiveId"`
	// The name of the configured service level objective of the database. This is the service level objective that is in the process of being applied to the database. Once successfully updated, it will match the value of serviceLevelObjective property.
	//
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
	RequestedServiceObjectiveName *string `pulumi:"requestedServiceObjectiveName"`
	// Conditional. If createMode is PointInTimeRestore, this value is required. If createMode is Restore, this value is optional. Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database. Must be greater than or equal to the source database's earliestRestoreDate value.
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// Indicates the name of the sample schema to apply when creating this database. If createMode is not Default, this value is ignored. Not supported for DataWarehouse edition.
	SampleName *string `pulumi:"sampleName"`
	// The current service level objective of the database.
	ServiceLevelObjective string `pulumi:"serviceLevelObjective"`
	// The list of service tier advisors for this database. Expanded property
	ServiceTierAdvisors []ServiceTierAdvisorResponse `pulumi:"serviceTierAdvisors"`
	// Conditional. If createMode is Restore and sourceDatabaseId is the deleted database's original resource id when it existed (as opposed to its current restorable dropped database id), then this value is required. Specifies the time that the database was deleted.
	SourceDatabaseDeletionDate *string `pulumi:"sourceDatabaseDeletionDate"`
	// Conditional. If createMode is Copy, NonReadableSecondary, OnlineSecondary, PointInTimeRestore, Recovery, or Restore, then this value is required. Specifies the resource ID of the source database. If createMode is NonReadableSecondary or OnlineSecondary, the name of the source database must be the same as the new database being created.
	SourceDatabaseId *string `pulumi:"sourceDatabaseId"`
	// The status of the database.
	Status string `pulumi:"status"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The transparent data encryption info for this database.
	TransparentDataEncryption []TransparentDataEncryptionResponse `pulumi:"transparentDataEncryption"`
	// Resource type.
	Type string `pulumi:"type"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}
