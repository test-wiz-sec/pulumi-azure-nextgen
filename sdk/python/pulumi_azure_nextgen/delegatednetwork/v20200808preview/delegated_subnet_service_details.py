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

__all__ = ['DelegatedSubnetServiceDetails']


class DelegatedSubnetServiceDetails(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 controller_details: Optional[pulumi.Input[pulumi.InputType['ControllerDetailsArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_name_: Optional[pulumi.Input[str]] = None,
                 subnet_details: Optional[pulumi.Input[pulumi.InputType['SubnetDetailsArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents an instance of a orchestrator.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ControllerDetailsArgs']] controller_details: controller details
        :param pulumi.Input[str] location: Location of the resource.
        :param pulumi.Input[str] resource_group_name: The name of the Azure Resource group of which a given DelegatedNetwork resource is part. This name must be at least 1 character in length, and no more than 90.
        :param pulumi.Input[str] resource_name_: The name of the resource. It must be a minimum of 3 characters, and a maximum of 63.
        :param pulumi.Input[pulumi.InputType['SubnetDetailsArgs']] subnet_details: orchestrator details
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The resource tags.
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

            __props__['controller_details'] = controller_details
            __props__['location'] = location
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_name_ is None:
                raise TypeError("Missing required property 'resource_name_'")
            __props__['resource_name'] = resource_name_
            __props__['subnet_details'] = subnet_details
            __props__['tags'] = tags
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['resource_guid'] = None
            __props__['type'] = None
        super(DelegatedSubnetServiceDetails, __self__).__init__(
            'azure-nextgen:delegatednetwork/v20200808preview:DelegatedSubnetServiceDetails',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DelegatedSubnetServiceDetails':
        """
        Get an existing DelegatedSubnetServiceDetails resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return DelegatedSubnetServiceDetails(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="controllerDetails")
    def controller_details(self) -> pulumi.Output[Optional['outputs.ControllerDetailsResponse']]:
        """
        controller details
        """
        return pulumi.get(self, "controller_details")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The current state of dnc delegated subnet resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceGuid")
    def resource_guid(self) -> pulumi.Output[str]:
        """
        Resource guid.
        """
        return pulumi.get(self, "resource_guid")

    @property
    @pulumi.getter(name="subnetDetails")
    def subnet_details(self) -> pulumi.Output[Optional['outputs.SubnetDetailsResponse']]:
        """
        orchestrator details
        """
        return pulumi.get(self, "subnet_details")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

