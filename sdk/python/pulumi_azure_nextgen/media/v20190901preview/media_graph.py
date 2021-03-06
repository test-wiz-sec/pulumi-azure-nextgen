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

__all__ = ['MediaGraph']


class MediaGraph(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 media_graph_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sinks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MediaGraphAssetSinkArgs']]]]] = None,
                 sources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MediaGraphRtspSourceArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The Media Graph.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The Media Services account name.
        :param pulumi.Input[str] description: Media Graph  description
        :param pulumi.Input[str] media_graph_name: The Media Graph name.
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the Azure subscription.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MediaGraphAssetSinkArgs']]]] sinks: Media Graph sinks
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['MediaGraphRtspSourceArgs']]]] sources: Media Graph sources
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['description'] = description
            if media_graph_name is None:
                raise TypeError("Missing required property 'media_graph_name'")
            __props__['media_graph_name'] = media_graph_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sinks is None:
                raise TypeError("Missing required property 'sinks'")
            __props__['sinks'] = sinks
            if sources is None:
                raise TypeError("Missing required property 'sources'")
            __props__['sources'] = sources
            __props__['created'] = None
            __props__['last_modified'] = None
            __props__['name'] = None
            __props__['state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:media/v20200201preview:MediaGraph")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(MediaGraph, __self__).__init__(
            'azure-nextgen:media/v20190901preview:MediaGraph',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'MediaGraph':
        """
        Get an existing MediaGraph resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return MediaGraph(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[str]:
        """
        Date the Media Graph was created
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Media Graph  description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> pulumi.Output[str]:
        """
        Date the Media Graph was last modified
        """
        return pulumi.get(self, "last_modified")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def sinks(self) -> pulumi.Output[Sequence['outputs.MediaGraphAssetSinkResponse']]:
        """
        Media Graph sinks
        """
        return pulumi.get(self, "sinks")

    @property
    @pulumi.getter
    def sources(self) -> pulumi.Output[Sequence['outputs.MediaGraphRtspSourceResponse']]:
        """
        Media Graph sources
        """
        return pulumi.get(self, "sources")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Media Graph state
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

