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

__all__ = ['ScheduledQueryRule']


class ScheduledQueryRule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[Union[pulumi.InputType['AlertingActionArgs'], pulumi.InputType['LogToMetricActionArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[pulumi.InputType['ScheduleArgs']]] = None,
                 source: Optional[pulumi.Input[pulumi.InputType['SourceArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The Log Search Rule resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union[pulumi.InputType['AlertingActionArgs'], pulumi.InputType['LogToMetricActionArgs']]] action: Action needs to be taken on rule execution.
        :param pulumi.Input[str] description: The description of the Log Search rule.
        :param pulumi.Input[str] enabled: The flag which indicates whether the Log Search rule is enabled. Value should be true or false
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] rule_name: The name of the rule.
        :param pulumi.Input[pulumi.InputType['ScheduleArgs']] schedule: Schedule (Frequency, Time Window) for rule. Required for action type - AlertingAction
        :param pulumi.Input[pulumi.InputType['SourceArgs']] source: Data Source against which rule will Query Data
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
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

            if action is None:
                raise TypeError("Missing required property 'action'")
            __props__['action'] = action
            __props__['description'] = description
            __props__['enabled'] = enabled
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if rule_name is None:
                raise TypeError("Missing required property 'rule_name'")
            __props__['rule_name'] = rule_name
            __props__['schedule'] = schedule
            if source is None:
                raise TypeError("Missing required property 'source'")
            __props__['source'] = source
            __props__['tags'] = tags
            __props__['last_updated_time'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:insights/v20180416:ScheduledQueryRule"), pulumi.Alias(type_="azure-nextgen:insights/v20200501preview:ScheduledQueryRule")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ScheduledQueryRule, __self__).__init__(
            'azure-nextgen:insights/latest:ScheduledQueryRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ScheduledQueryRule':
        """
        Get an existing ScheduledQueryRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ScheduledQueryRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[Any]:
        """
        Action needs to be taken on rule execution.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the Log Search rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[str]]:
        """
        The flag which indicates whether the Log Search rule is enabled. Value should be true or false
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> pulumi.Output[str]:
        """
        Last time the rule was updated in IS08601 format.
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Azure resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Provisioning state of the scheduled query rule
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[Optional['outputs.ScheduleResponse']]:
        """
        Schedule (Frequency, Time Window) for rule. Required for action type - AlertingAction
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output['outputs.SourceResponse']:
        """
        Data Source against which rule will Query Data
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Azure resource type
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

