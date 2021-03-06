# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs

__all__ = ['SqlManagedInstance']


class SqlManagedInstance(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin: Optional[pulumi.Input[str]] = None,
                 data_controller_id: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 instance_endpoint: Optional[pulumi.Input[str]] = None,
                 k8s_raw: Optional[Any] = None,
                 last_uploaded_date: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sql_managed_instance_name: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 v_core: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A SqlManagedInstance.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin: The instance admin user
        :param pulumi.Input[str] data_controller_id: null
        :param pulumi.Input[str] end_time: The instance end time
        :param pulumi.Input[str] instance_endpoint: The on premise instance endpoint
        :param Any k8s_raw: The raw kubernetes information
        :param pulumi.Input[str] last_uploaded_date: Last uploaded date from on premise cluster. Defaults to current date time
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[str] resource_group_name: The name of the Azure resource group
        :param pulumi.Input[str] sql_managed_instance_name: The name of SQL Managed Instances
        :param pulumi.Input[str] start_time: The instance start time
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[str] v_core: The instance vCore
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['admin'] = admin
            __props__['data_controller_id'] = data_controller_id
            __props__['end_time'] = end_time
            __props__['instance_endpoint'] = instance_endpoint
            __props__['k8s_raw'] = k8s_raw
            __props__['last_uploaded_date'] = last_uploaded_date
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sql_managed_instance_name is None:
                raise TypeError("Missing required property 'sql_managed_instance_name'")
            __props__['sql_managed_instance_name'] = sql_managed_instance_name
            __props__['start_time'] = start_time
            __props__['tags'] = tags
            __props__['v_core'] = v_core
            __props__['name'] = None
            __props__['system_data'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:azuredata/v20190724preview:SqlManagedInstance")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(SqlManagedInstance, __self__).__init__(
            'azure-nextgen:azuredata/v20200908preview:SqlManagedInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SqlManagedInstance':
        """
        Get an existing SqlManagedInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return SqlManagedInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def admin(self) -> pulumi.Output[Optional[str]]:
        """
        The instance admin user
        """
        return pulumi.get(self, "admin")

    @property
    @pulumi.getter(name="dataControllerId")
    def data_controller_id(self) -> pulumi.Output[Optional[str]]:
        """
        null
        """
        return pulumi.get(self, "data_controller_id")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[Optional[str]]:
        """
        The instance end time
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="instanceEndpoint")
    def instance_endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        The on premise instance endpoint
        """
        return pulumi.get(self, "instance_endpoint")

    @property
    @pulumi.getter(name="k8sRaw")
    def k8s_raw(self) -> pulumi.Output[Optional[Any]]:
        """
        The raw kubernetes information
        """
        return pulumi.get(self, "k8s_raw")

    @property
    @pulumi.getter(name="lastUploadedDate")
    def last_uploaded_date(self) -> pulumi.Output[Optional[str]]:
        """
        Last uploaded date from on premise cluster. Defaults to current date time
        """
        return pulumi.get(self, "last_uploaded_date")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[Optional[str]]:
        """
        The instance start time
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Read only system data
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vCore")
    def v_core(self) -> pulumi.Output[Optional[str]]:
        """
        The instance vCore
        """
        return pulumi.get(self, "v_core")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

