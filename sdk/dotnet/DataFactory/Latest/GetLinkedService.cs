// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataFactory.Latest
{
    public static class GetLinkedService
    {
        public static Task<GetLinkedServiceResult> InvokeAsync(GetLinkedServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLinkedServiceResult>("azure-nextgen:datafactory/latest:getLinkedService", args ?? new GetLinkedServiceArgs(), options.WithVersion());
    }


    public sealed class GetLinkedServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The factory name.
        /// </summary>
        [Input("factoryName", required: true)]
        public string FactoryName { get; set; } = null!;

        /// <summary>
        /// The linked service name.
        /// </summary>
        [Input("linkedServiceName", required: true)]
        public string LinkedServiceName { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetLinkedServiceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLinkedServiceResult
    {
        /// <summary>
        /// Etag identifies change in the resource.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of linked service.
        /// </summary>
        public readonly Union<Outputs.AmazonMWSLinkedServiceResponse, Union<Outputs.AmazonRedshiftLinkedServiceResponse, Union<Outputs.AmazonS3LinkedServiceResponse, Union<Outputs.AzureBatchLinkedServiceResponse, Union<Outputs.AzureBlobFSLinkedServiceResponse, Union<Outputs.AzureBlobStorageLinkedServiceResponse, Union<Outputs.AzureDataExplorerLinkedServiceResponse, Union<Outputs.AzureDataLakeAnalyticsLinkedServiceResponse, Union<Outputs.AzureDataLakeStoreLinkedServiceResponse, Union<Outputs.AzureDatabricksDeltaLakeLinkedServiceResponse, Union<Outputs.AzureDatabricksLinkedServiceResponse, Union<Outputs.AzureFileStorageLinkedServiceResponse, Union<Outputs.AzureFunctionLinkedServiceResponse, Union<Outputs.AzureKeyVaultLinkedServiceResponse, Union<Outputs.AzureMLLinkedServiceResponse, Union<Outputs.AzureMLServiceLinkedServiceResponse, Union<Outputs.AzureMariaDBLinkedServiceResponse, Union<Outputs.AzureMySqlLinkedServiceResponse, Union<Outputs.AzurePostgreSqlLinkedServiceResponse, Union<Outputs.AzureSearchLinkedServiceResponse, Union<Outputs.AzureSqlDWLinkedServiceResponse, Union<Outputs.AzureSqlDatabaseLinkedServiceResponse, Union<Outputs.AzureSqlMILinkedServiceResponse, Union<Outputs.AzureStorageLinkedServiceResponse, Union<Outputs.AzureTableStorageLinkedServiceResponse, Union<Outputs.CassandraLinkedServiceResponse, Union<Outputs.CommonDataServiceForAppsLinkedServiceResponse, Union<Outputs.ConcurLinkedServiceResponse, Union<Outputs.CosmosDbLinkedServiceResponse, Union<Outputs.CosmosDbMongoDbApiLinkedServiceResponse, Union<Outputs.CouchbaseLinkedServiceResponse, Union<Outputs.CustomDataSourceLinkedServiceResponse, Union<Outputs.Db2LinkedServiceResponse, Union<Outputs.DrillLinkedServiceResponse, Union<Outputs.DynamicsAXLinkedServiceResponse, Union<Outputs.DynamicsCrmLinkedServiceResponse, Union<Outputs.DynamicsLinkedServiceResponse, Union<Outputs.EloquaLinkedServiceResponse, Union<Outputs.FileServerLinkedServiceResponse, Union<Outputs.FtpServerLinkedServiceResponse, Union<Outputs.GoogleAdWordsLinkedServiceResponse, Union<Outputs.GoogleBigQueryLinkedServiceResponse, Union<Outputs.GoogleCloudStorageLinkedServiceResponse, Union<Outputs.GreenplumLinkedServiceResponse, Union<Outputs.HBaseLinkedServiceResponse, Union<Outputs.HDInsightLinkedServiceResponse, Union<Outputs.HDInsightOnDemandLinkedServiceResponse, Union<Outputs.HdfsLinkedServiceResponse, Union<Outputs.HiveLinkedServiceResponse, Union<Outputs.HttpLinkedServiceResponse, Union<Outputs.HubspotLinkedServiceResponse, Union<Outputs.ImpalaLinkedServiceResponse, Union<Outputs.InformixLinkedServiceResponse, Union<Outputs.JiraLinkedServiceResponse, Union<Outputs.MagentoLinkedServiceResponse, Union<Outputs.MariaDBLinkedServiceResponse, Union<Outputs.MarketoLinkedServiceResponse, Union<Outputs.MicrosoftAccessLinkedServiceResponse, Union<Outputs.MongoDbAtlasLinkedServiceResponse, Union<Outputs.MongoDbLinkedServiceResponse, Union<Outputs.MongoDbV2LinkedServiceResponse, Union<Outputs.MySqlLinkedServiceResponse, Union<Outputs.NetezzaLinkedServiceResponse, Union<Outputs.ODataLinkedServiceResponse, Union<Outputs.OdbcLinkedServiceResponse, Union<Outputs.Office365LinkedServiceResponse, Union<Outputs.OracleLinkedServiceResponse, Union<Outputs.OracleServiceCloudLinkedServiceResponse, Union<Outputs.PaypalLinkedServiceResponse, Union<Outputs.PhoenixLinkedServiceResponse, Union<Outputs.PostgreSqlLinkedServiceResponse, Union<Outputs.PrestoLinkedServiceResponse, Union<Outputs.QuickBooksLinkedServiceResponse, Union<Outputs.ResponsysLinkedServiceResponse, Union<Outputs.RestServiceLinkedServiceResponse, Union<Outputs.SalesforceLinkedServiceResponse, Union<Outputs.SalesforceMarketingCloudLinkedServiceResponse, Union<Outputs.SalesforceServiceCloudLinkedServiceResponse, Union<Outputs.SapBWLinkedServiceResponse, Union<Outputs.SapCloudForCustomerLinkedServiceResponse, Union<Outputs.SapEccLinkedServiceResponse, Union<Outputs.SapHanaLinkedServiceResponse, Union<Outputs.SapOpenHubLinkedServiceResponse, Union<Outputs.SapTableLinkedServiceResponse, Union<Outputs.ServiceNowLinkedServiceResponse, Union<Outputs.SftpServerLinkedServiceResponse, Union<Outputs.SharePointOnlineListLinkedServiceResponse, Union<Outputs.ShopifyLinkedServiceResponse, Union<Outputs.SnowflakeLinkedServiceResponse, Union<Outputs.SparkLinkedServiceResponse, Union<Outputs.SqlServerLinkedServiceResponse, Union<Outputs.SquareLinkedServiceResponse, Union<Outputs.SybaseLinkedServiceResponse, Union<Outputs.TeradataLinkedServiceResponse, Union<Outputs.VerticaLinkedServiceResponse, Union<Outputs.WebLinkedServiceResponse, Union<Outputs.XeroLinkedServiceResponse, Outputs.ZohoLinkedServiceResponse>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> Properties;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetLinkedServiceResult(
            string etag,

            string name,

            Union<Outputs.AmazonMWSLinkedServiceResponse, Union<Outputs.AmazonRedshiftLinkedServiceResponse, Union<Outputs.AmazonS3LinkedServiceResponse, Union<Outputs.AzureBatchLinkedServiceResponse, Union<Outputs.AzureBlobFSLinkedServiceResponse, Union<Outputs.AzureBlobStorageLinkedServiceResponse, Union<Outputs.AzureDataExplorerLinkedServiceResponse, Union<Outputs.AzureDataLakeAnalyticsLinkedServiceResponse, Union<Outputs.AzureDataLakeStoreLinkedServiceResponse, Union<Outputs.AzureDatabricksDeltaLakeLinkedServiceResponse, Union<Outputs.AzureDatabricksLinkedServiceResponse, Union<Outputs.AzureFileStorageLinkedServiceResponse, Union<Outputs.AzureFunctionLinkedServiceResponse, Union<Outputs.AzureKeyVaultLinkedServiceResponse, Union<Outputs.AzureMLLinkedServiceResponse, Union<Outputs.AzureMLServiceLinkedServiceResponse, Union<Outputs.AzureMariaDBLinkedServiceResponse, Union<Outputs.AzureMySqlLinkedServiceResponse, Union<Outputs.AzurePostgreSqlLinkedServiceResponse, Union<Outputs.AzureSearchLinkedServiceResponse, Union<Outputs.AzureSqlDWLinkedServiceResponse, Union<Outputs.AzureSqlDatabaseLinkedServiceResponse, Union<Outputs.AzureSqlMILinkedServiceResponse, Union<Outputs.AzureStorageLinkedServiceResponse, Union<Outputs.AzureTableStorageLinkedServiceResponse, Union<Outputs.CassandraLinkedServiceResponse, Union<Outputs.CommonDataServiceForAppsLinkedServiceResponse, Union<Outputs.ConcurLinkedServiceResponse, Union<Outputs.CosmosDbLinkedServiceResponse, Union<Outputs.CosmosDbMongoDbApiLinkedServiceResponse, Union<Outputs.CouchbaseLinkedServiceResponse, Union<Outputs.CustomDataSourceLinkedServiceResponse, Union<Outputs.Db2LinkedServiceResponse, Union<Outputs.DrillLinkedServiceResponse, Union<Outputs.DynamicsAXLinkedServiceResponse, Union<Outputs.DynamicsCrmLinkedServiceResponse, Union<Outputs.DynamicsLinkedServiceResponse, Union<Outputs.EloquaLinkedServiceResponse, Union<Outputs.FileServerLinkedServiceResponse, Union<Outputs.FtpServerLinkedServiceResponse, Union<Outputs.GoogleAdWordsLinkedServiceResponse, Union<Outputs.GoogleBigQueryLinkedServiceResponse, Union<Outputs.GoogleCloudStorageLinkedServiceResponse, Union<Outputs.GreenplumLinkedServiceResponse, Union<Outputs.HBaseLinkedServiceResponse, Union<Outputs.HDInsightLinkedServiceResponse, Union<Outputs.HDInsightOnDemandLinkedServiceResponse, Union<Outputs.HdfsLinkedServiceResponse, Union<Outputs.HiveLinkedServiceResponse, Union<Outputs.HttpLinkedServiceResponse, Union<Outputs.HubspotLinkedServiceResponse, Union<Outputs.ImpalaLinkedServiceResponse, Union<Outputs.InformixLinkedServiceResponse, Union<Outputs.JiraLinkedServiceResponse, Union<Outputs.MagentoLinkedServiceResponse, Union<Outputs.MariaDBLinkedServiceResponse, Union<Outputs.MarketoLinkedServiceResponse, Union<Outputs.MicrosoftAccessLinkedServiceResponse, Union<Outputs.MongoDbAtlasLinkedServiceResponse, Union<Outputs.MongoDbLinkedServiceResponse, Union<Outputs.MongoDbV2LinkedServiceResponse, Union<Outputs.MySqlLinkedServiceResponse, Union<Outputs.NetezzaLinkedServiceResponse, Union<Outputs.ODataLinkedServiceResponse, Union<Outputs.OdbcLinkedServiceResponse, Union<Outputs.Office365LinkedServiceResponse, Union<Outputs.OracleLinkedServiceResponse, Union<Outputs.OracleServiceCloudLinkedServiceResponse, Union<Outputs.PaypalLinkedServiceResponse, Union<Outputs.PhoenixLinkedServiceResponse, Union<Outputs.PostgreSqlLinkedServiceResponse, Union<Outputs.PrestoLinkedServiceResponse, Union<Outputs.QuickBooksLinkedServiceResponse, Union<Outputs.ResponsysLinkedServiceResponse, Union<Outputs.RestServiceLinkedServiceResponse, Union<Outputs.SalesforceLinkedServiceResponse, Union<Outputs.SalesforceMarketingCloudLinkedServiceResponse, Union<Outputs.SalesforceServiceCloudLinkedServiceResponse, Union<Outputs.SapBWLinkedServiceResponse, Union<Outputs.SapCloudForCustomerLinkedServiceResponse, Union<Outputs.SapEccLinkedServiceResponse, Union<Outputs.SapHanaLinkedServiceResponse, Union<Outputs.SapOpenHubLinkedServiceResponse, Union<Outputs.SapTableLinkedServiceResponse, Union<Outputs.ServiceNowLinkedServiceResponse, Union<Outputs.SftpServerLinkedServiceResponse, Union<Outputs.SharePointOnlineListLinkedServiceResponse, Union<Outputs.ShopifyLinkedServiceResponse, Union<Outputs.SnowflakeLinkedServiceResponse, Union<Outputs.SparkLinkedServiceResponse, Union<Outputs.SqlServerLinkedServiceResponse, Union<Outputs.SquareLinkedServiceResponse, Union<Outputs.SybaseLinkedServiceResponse, Union<Outputs.TeradataLinkedServiceResponse, Union<Outputs.VerticaLinkedServiceResponse, Union<Outputs.WebLinkedServiceResponse, Union<Outputs.XeroLinkedServiceResponse, Outputs.ZohoLinkedServiceResponse>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> properties,

            string type)
        {
            Etag = etag;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
