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
    'GetSqlServerInstanceResult',
    'AwaitableGetSqlServerInstanceResult',
    'get_sql_server_instance',
]

@pulumi.output_type
class GetSqlServerInstanceResult:
    """
    A SqlServerInstance.
    """
    def __init__(__self__, container_resource_id=None, create_time=None, edition=None, location=None, name=None, status=None, system_data=None, tags=None, type=None, update_time=None, v_core=None, version=None):
        if container_resource_id and not isinstance(container_resource_id, str):
            raise TypeError("Expected argument 'container_resource_id' to be a str")
        pulumi.set(__self__, "container_resource_id", container_resource_id)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if edition and not isinstance(edition, str):
            raise TypeError("Expected argument 'edition' to be a str")
        pulumi.set(__self__, "edition", edition)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
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
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if v_core and not isinstance(v_core, str):
            raise TypeError("Expected argument 'v_core' to be a str")
        pulumi.set(__self__, "v_core", v_core)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="containerResourceId")
    def container_resource_id(self) -> str:
        """
        ARM Resource id of the container resource (Azure Arc for Servers)
        """
        return pulumi.get(self, "container_resource_id")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when the resource was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def edition(self) -> str:
        """
        SQL Server edition.
        """
        return pulumi.get(self, "edition")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

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
        The cloud connectivity status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Read only system data
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
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time when the resource was last updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="vCore")
    def v_core(self) -> str:
        """
        The number of logical processors used by the SQL Server instance.
        """
        return pulumi.get(self, "v_core")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        SQL Server version.
        """
        return pulumi.get(self, "version")


class AwaitableGetSqlServerInstanceResult(GetSqlServerInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSqlServerInstanceResult(
            container_resource_id=self.container_resource_id,
            create_time=self.create_time,
            edition=self.edition,
            location=self.location,
            name=self.name,
            status=self.status,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type,
            update_time=self.update_time,
            v_core=self.v_core,
            version=self.version)


def get_sql_server_instance(resource_group_name: Optional[str] = None,
                            sql_server_instance_name: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSqlServerInstanceResult:
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: The name of the Azure resource group
    :param str sql_server_instance_name: Name of SQL Server Instance
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['sqlServerInstanceName'] = sql_server_instance_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:azuredata/v20200908preview:getSqlServerInstance', __args__, opts=opts, typ=GetSqlServerInstanceResult).value

    return AwaitableGetSqlServerInstanceResult(
        container_resource_id=__ret__.container_resource_id,
        create_time=__ret__.create_time,
        edition=__ret__.edition,
        location=__ret__.location,
        name=__ret__.name,
        status=__ret__.status,
        system_data=__ret__.system_data,
        tags=__ret__.tags,
        type=__ret__.type,
        update_time=__ret__.update_time,
        v_core=__ret__.v_core,
        version=__ret__.version)
