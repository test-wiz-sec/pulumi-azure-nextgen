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

__all__ = ['ProtectionContainer']


class ProtectionContainer(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_name: Optional[pulumi.Input[str]] = None,
                 e_tag: Optional[pulumi.Input[str]] = None,
                 fabric_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 properties: Optional[pulumi.Input[Union[pulumi.InputType['AzureSqlContainerArgs'], pulumi.InputType['AzureStorageContainerArgs'], pulumi.InputType['AzureWorkloadContainerArgs'], pulumi.InputType['DpmContainerArgs'], pulumi.InputType['GenericContainerArgs'], pulumi.InputType['IaaSVMContainerArgs'], pulumi.InputType['MabContainerArgs']]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vault_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Base class for container with backup items. Containers with specific workloads are derived from this class.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_name: Name of the container to be registered.
        :param pulumi.Input[str] e_tag: Optional ETag.
        :param pulumi.Input[str] fabric_name: Fabric name associated with the container.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[Union[pulumi.InputType['AzureSqlContainerArgs'], pulumi.InputType['AzureStorageContainerArgs'], pulumi.InputType['AzureWorkloadContainerArgs'], pulumi.InputType['DpmContainerArgs'], pulumi.InputType['GenericContainerArgs'], pulumi.InputType['IaaSVMContainerArgs'], pulumi.InputType['MabContainerArgs']]] properties: ProtectionContainerResource properties
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the recovery services vault is present.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[str] vault_name: The name of the recovery services vault.
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

            if container_name is None:
                raise TypeError("Missing required property 'container_name'")
            __props__['container_name'] = container_name
            __props__['e_tag'] = e_tag
            if fabric_name is None:
                raise TypeError("Missing required property 'fabric_name'")
            __props__['fabric_name'] = fabric_name
            __props__['location'] = location
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            if vault_name is None:
                raise TypeError("Missing required property 'vault_name'")
            __props__['vault_name'] = vault_name
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:recoveryservices/v20161201:ProtectionContainer")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ProtectionContainer, __self__).__init__(
            'azure-nextgen:recoveryservices/latest:ProtectionContainer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ProtectionContainer':
        """
        Get an existing ProtectionContainer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ProtectionContainer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="eTag")
    def e_tag(self) -> pulumi.Output[Optional[str]]:
        """
        Optional ETag.
        """
        return pulumi.get(self, "e_tag")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name associated with the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output[Any]:
        """
        ProtectionContainerResource properties
        """
        return pulumi.get(self, "properties")

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
        Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

