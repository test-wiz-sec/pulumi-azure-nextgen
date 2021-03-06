# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetSqlDatabaseResult',
    'AwaitableGetSqlDatabaseResult',
    'get_sql_database',
]

@pulumi.output_type
class GetSqlDatabaseResult:
    """
    A sql database resource.
    """
    def __init__(__self__, collation=None, database_guid=None, location=None, max_size_bytes=None, name=None, status=None, system_data=None, tags=None, type=None):
        if collation and not isinstance(collation, str):
            raise TypeError("Expected argument 'collation' to be a str")
        pulumi.set(__self__, "collation", collation)
        if database_guid and not isinstance(database_guid, str):
            raise TypeError("Expected argument 'database_guid' to be a str")
        pulumi.set(__self__, "database_guid", database_guid)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if max_size_bytes and not isinstance(max_size_bytes, int):
            raise TypeError("Expected argument 'max_size_bytes' to be a int")
        pulumi.set(__self__, "max_size_bytes", max_size_bytes)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def collation(self) -> Optional[str]:
        """
        The collation of the database.
        """
        return pulumi.get(self, "collation")

    @property
    @pulumi.getter(name="databaseGuid")
    def database_guid(self) -> str:
        """
        The Guid of the database.
        """
        return pulumi.get(self, "database_guid")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maxSizeBytes")
    def max_size_bytes(self) -> Optional[int]:
        """
        The max size of the database expressed in bytes.
        """
        return pulumi.get(self, "max_size_bytes")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the database.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        SystemData of SqlDatabase.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetSqlDatabaseResult(GetSqlDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSqlDatabaseResult(
            collation=self.collation,
            database_guid=self.database_guid,
            location=self.location,
            max_size_bytes=self.max_size_bytes,
            name=self.name,
            status=self.status,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type)


def get_sql_database(resource_group_name: Optional[str] = None,
                     sql_database_name: Optional[str] = None,
                     workspace_name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSqlDatabaseResult:
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str sql_database_name: The name of the sql database.
    :param str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['sqlDatabaseName'] = sql_database_name
    __args__['workspaceName'] = workspace_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:synapse/v20200401preview:getSqlDatabase', __args__, opts=opts, typ=GetSqlDatabaseResult).value

    return AwaitableGetSqlDatabaseResult(
        collation=__ret__.collation,
        database_guid=__ret__.database_guid,
        location=__ret__.location,
        max_size_bytes=__ret__.max_size_bytes,
        name=__ret__.name,
        status=__ret__.status,
        system_data=__ret__.system_data,
        tags=__ret__.tags,
        type=__ret__.type)
