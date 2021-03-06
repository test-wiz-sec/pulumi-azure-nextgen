# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Server']


class Server(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 administrator_login: Optional[pulumi.Input[str]] = None,
                 administrator_login_password: Optional[pulumi.Input[str]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 create_mode: Optional[pulumi.Input[str]] = None,
                 delegated_subnet_arguments: Optional[pulumi.Input[pulumi.InputType['ServerPropertiesDelegatedSubnetArgumentsArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 ha_enabled: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['IdentityArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 maintenance_window: Optional[pulumi.Input[pulumi.InputType['MaintenanceWindowArgs']]] = None,
                 point_in_time_utc: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['SkuArgs']]] = None,
                 source_server_name: Optional[pulumi.Input[str]] = None,
                 storage_profile: Optional[pulumi.Input[pulumi.InputType['StorageProfileArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents a server.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] administrator_login: The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
        :param pulumi.Input[str] administrator_login_password: The administrator login password (required for server creation).
        :param pulumi.Input[str] availability_zone: availability Zone information of the server.
        :param pulumi.Input[str] create_mode: The mode to create a new PostgreSQL server.
        :param pulumi.Input[str] display_name: The display name of a server.
        :param pulumi.Input[str] ha_enabled: stand by count value can be either enabled or disabled
        :param pulumi.Input[pulumi.InputType['IdentityArgs']] identity: The Azure Active Directory identity of the server.
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[pulumi.InputType['MaintenanceWindowArgs']] maintenance_window: Maintenance window of a server.
        :param pulumi.Input[str] point_in_time_utc: Restore point creation time (ISO8601 format), specifying the time to restore from.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] server_name: The name of the server.
        :param pulumi.Input[pulumi.InputType['SkuArgs']] sku: The SKU (pricing tier) of the server.
        :param pulumi.Input[str] source_server_name: The source PostgreSQL server name to restore from.
        :param pulumi.Input[pulumi.InputType['StorageProfileArgs']] storage_profile: Storage profile of a server.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[str] version: PostgreSQL Server version.
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

            __props__['administrator_login'] = administrator_login
            __props__['administrator_login_password'] = administrator_login_password
            __props__['availability_zone'] = availability_zone
            __props__['create_mode'] = create_mode
            __props__['delegated_subnet_arguments'] = delegated_subnet_arguments
            __props__['display_name'] = display_name
            __props__['ha_enabled'] = ha_enabled
            __props__['identity'] = identity
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['maintenance_window'] = maintenance_window
            __props__['point_in_time_utc'] = point_in_time_utc
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if server_name is None:
                raise TypeError("Missing required property 'server_name'")
            __props__['server_name'] = server_name
            __props__['sku'] = sku
            __props__['source_server_name'] = source_server_name
            __props__['storage_profile'] = storage_profile
            __props__['tags'] = tags
            __props__['version'] = version
            __props__['byok_enforcement'] = None
            __props__['fully_qualified_domain_name'] = None
            __props__['ha_state'] = None
            __props__['name'] = None
            __props__['public_network_access'] = None
            __props__['standby_availability_zone'] = None
            __props__['state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:dbforpostgresql/v20200214privatepreview:Server")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Server, __self__).__init__(
            'azure-nextgen:dbforpostgresql/v20200214preview:Server',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Server':
        """
        Get an existing Server resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Server(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="administratorLogin")
    def administrator_login(self) -> pulumi.Output[Optional[str]]:
        """
        The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
        """
        return pulumi.get(self, "administrator_login")

    @property
    @pulumi.getter(name="administratorLoginPassword")
    def administrator_login_password(self) -> pulumi.Output[Optional[str]]:
        """
        The administrator login password (required for server creation).
        """
        return pulumi.get(self, "administrator_login_password")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[Optional[str]]:
        """
        availability Zone information of the server.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="byokEnforcement")
    def byok_enforcement(self) -> pulumi.Output[str]:
        """
        Status showing whether the data encryption is enabled with customer-managed keys.
        """
        return pulumi.get(self, "byok_enforcement")

    @property
    @pulumi.getter(name="createMode")
    def create_mode(self) -> pulumi.Output[Optional[str]]:
        """
        The mode to create a new PostgreSQL server.
        """
        return pulumi.get(self, "create_mode")

    @property
    @pulumi.getter(name="delegatedSubnetArguments")
    def delegated_subnet_arguments(self) -> pulumi.Output[Optional['outputs.ServerPropertiesResponseDelegatedSubnetArguments']]:
        return pulumi.get(self, "delegated_subnet_arguments")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        The display name of a server.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="fullyQualifiedDomainName")
    def fully_qualified_domain_name(self) -> pulumi.Output[str]:
        """
        The fully qualified domain name of a server.
        """
        return pulumi.get(self, "fully_qualified_domain_name")

    @property
    @pulumi.getter(name="haEnabled")
    def ha_enabled(self) -> pulumi.Output[Optional[str]]:
        """
        stand by count value can be either enabled or disabled
        """
        return pulumi.get(self, "ha_enabled")

    @property
    @pulumi.getter(name="haState")
    def ha_state(self) -> pulumi.Output[str]:
        """
        A state of a HA server that is visible to user.
        """
        return pulumi.get(self, "ha_state")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.IdentityResponse']]:
        """
        The Azure Active Directory identity of the server.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> pulumi.Output[Optional['outputs.MaintenanceWindowResponse']]:
        """
        Maintenance window of a server.
        """
        return pulumi.get(self, "maintenance_window")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pointInTimeUTC")
    def point_in_time_utc(self) -> pulumi.Output[Optional[str]]:
        """
        Restore point creation time (ISO8601 format), specifying the time to restore from.
        """
        return pulumi.get(self, "point_in_time_utc")

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> pulumi.Output[str]:
        """
        public network access is enabled or not
        """
        return pulumi.get(self, "public_network_access")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.SkuResponse']]:
        """
        The SKU (pricing tier) of the server.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="sourceServerName")
    def source_server_name(self) -> pulumi.Output[Optional[str]]:
        """
        The source PostgreSQL server name to restore from.
        """
        return pulumi.get(self, "source_server_name")

    @property
    @pulumi.getter(name="standbyAvailabilityZone")
    def standby_availability_zone(self) -> pulumi.Output[str]:
        """
        availability Zone information of the server.
        """
        return pulumi.get(self, "standby_availability_zone")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        A state of a server that is visible to user.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="storageProfile")
    def storage_profile(self) -> pulumi.Output[Optional['outputs.StorageProfileResponse']]:
        """
        Storage profile of a server.
        """
        return pulumi.get(self, "storage_profile")

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
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[Optional[str]]:
        """
        PostgreSQL Server version.
        """
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

