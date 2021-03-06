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

__all__ = ['JobSchedule']


class JobSchedule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 automation_account_name: Optional[pulumi.Input[str]] = None,
                 job_schedule_id: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 run_on: Optional[pulumi.Input[str]] = None,
                 runbook: Optional[pulumi.Input[pulumi.InputType['RunbookAssociationPropertyArgs']]] = None,
                 schedule: Optional[pulumi.Input[pulumi.InputType['ScheduleAssociationPropertyArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Definition of the job schedule.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account.
        :param pulumi.Input[str] job_schedule_id: The job schedule name.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: Gets or sets a list of job properties.
        :param pulumi.Input[str] resource_group_name: Name of an Azure Resource group.
        :param pulumi.Input[str] run_on: Gets or sets the hybrid worker group that the scheduled job should run on.
        :param pulumi.Input[pulumi.InputType['RunbookAssociationPropertyArgs']] runbook: Gets or sets the runbook.
        :param pulumi.Input[pulumi.InputType['ScheduleAssociationPropertyArgs']] schedule: Gets or sets the schedule.
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

            if automation_account_name is None:
                raise TypeError("Missing required property 'automation_account_name'")
            __props__['automation_account_name'] = automation_account_name
            if job_schedule_id is None:
                raise TypeError("Missing required property 'job_schedule_id'")
            __props__['job_schedule_id'] = job_schedule_id
            __props__['parameters'] = parameters
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['run_on'] = run_on
            if runbook is None:
                raise TypeError("Missing required property 'runbook'")
            __props__['runbook'] = runbook
            if schedule is None:
                raise TypeError("Missing required property 'schedule'")
            __props__['schedule'] = schedule
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:automation/latest:JobSchedule")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(JobSchedule, __self__).__init__(
            'azure-nextgen:automation/v20151031:JobSchedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'JobSchedule':
        """
        Get an existing JobSchedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return JobSchedule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="jobScheduleId")
    def job_schedule_id(self) -> pulumi.Output[Optional[str]]:
        """
        Gets or sets the id of job schedule.
        """
        return pulumi.get(self, "job_schedule_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Gets the name of the variable.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Gets or sets the parameters of the job schedule.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="runOn")
    def run_on(self) -> pulumi.Output[Optional[str]]:
        """
        Gets or sets the hybrid worker group that the scheduled job should run on.
        """
        return pulumi.get(self, "run_on")

    @property
    @pulumi.getter
    def runbook(self) -> pulumi.Output[Optional['outputs.RunbookAssociationPropertyResponse']]:
        """
        Gets or sets the runbook.
        """
        return pulumi.get(self, "runbook")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[Optional['outputs.ScheduleAssociationPropertyResponse']]:
        """
        Gets or sets the schedule.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

