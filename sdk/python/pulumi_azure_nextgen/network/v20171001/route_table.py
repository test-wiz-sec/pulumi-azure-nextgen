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

__all__ = ['RouteTable']


class RouteTable(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_bgp_route_propagation: Optional[pulumi.Input[bool]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 provisioning_state: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 route_table_name: Optional[pulumi.Input[str]] = None,
                 routes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouteArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Route table resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_bgp_route_propagation: Gets or sets whether to disable the routes learned by BGP on that route table. True means disable.
        :param pulumi.Input[str] etag: Gets a unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] provisioning_state: The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] route_table_name: The name of the route table.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RouteArgs']]]] routes: Collection of routes contained within a route table.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
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

            __props__['disable_bgp_route_propagation'] = disable_bgp_route_propagation
            __props__['etag'] = etag
            __props__['id'] = id
            __props__['location'] = location
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if route_table_name is None:
                raise TypeError("Missing required property 'route_table_name'")
            __props__['route_table_name'] = route_table_name
            __props__['routes'] = routes
            __props__['tags'] = tags
            __props__['name'] = None
            __props__['subnets'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:network/latest:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20150501preview:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20150615:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20160330:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20160601:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20160901:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20161201:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20170301:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20170601:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20170801:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20170901:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20171101:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20180101:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20180201:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20180401:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20180601:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20180701:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20180801:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20181001:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20181101:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20181201:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20190201:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20190401:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20190601:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20190701:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20190801:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20190901:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20191101:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20191201:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20200301:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20200401:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20200501:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20200601:RouteTable"), pulumi.Alias(type_="azure-nextgen:network/v20200701:RouteTable")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(RouteTable, __self__).__init__(
            'azure-nextgen:network/v20171001:RouteTable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RouteTable':
        """
        Get an existing RouteTable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return RouteTable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="disableBgpRoutePropagation")
    def disable_bgp_route_propagation(self) -> pulumi.Output[Optional[bool]]:
        """
        Gets or sets whether to disable the routes learned by BGP on that route table. True means disable.
        """
        return pulumi.get(self, "disable_bgp_route_propagation")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[Optional[str]]:
        """
        Gets a unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

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
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[Optional[str]]:
        """
        The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def routes(self) -> pulumi.Output[Optional[Sequence['outputs.RouteResponse']]]:
        """
        Collection of routes contained within a route table.
        """
        return pulumi.get(self, "routes")

    @property
    @pulumi.getter
    def subnets(self) -> pulumi.Output[Sequence['outputs.SubnetResponse']]:
        """
        A collection of references to subnets.
        """
        return pulumi.get(self, "subnets")

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
        Resource type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

