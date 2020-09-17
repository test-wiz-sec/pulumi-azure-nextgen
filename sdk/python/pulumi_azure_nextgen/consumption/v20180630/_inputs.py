# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'BudgetTimePeriodArgs',
    'FiltersArgs',
    'NotificationArgs',
]

@pulumi.input_type
class BudgetTimePeriodArgs:
    def __init__(__self__, *,
                 start_date: pulumi.Input[str],
                 end_date: Optional[pulumi.Input[str]] = None):
        """
        The start and end date for a budget.
        :param pulumi.Input[str] start_date: The start date for the budget.
        :param pulumi.Input[str] end_date: The end date for the budget. If not provided, we default this to 10 years from the start date.
        """
        pulumi.set(__self__, "start_date", start_date)
        if end_date is not None:
            pulumi.set(__self__, "end_date", end_date)

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> pulumi.Input[str]:
        """
        The start date for the budget.
        """
        return pulumi.get(self, "start_date")

    @start_date.setter
    def start_date(self, value: pulumi.Input[str]):
        pulumi.set(self, "start_date", value)

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[pulumi.Input[str]]:
        """
        The end date for the budget. If not provided, we default this to 10 years from the start date.
        """
        return pulumi.get(self, "end_date")

    @end_date.setter
    def end_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_date", value)


@pulumi.input_type
class FiltersArgs:
    def __init__(__self__, *,
                 meters: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]]] = None):
        """
        May be used to filter budgets by resource group, resource, or meter.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] meters: The list of filters on meters (GUID), mandatory for budgets of usage category. 
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resource_groups: The list of filters on resource groups, allowed at subscription level only.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] resources: The list of filters on resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]] tags: The dictionary of filters on tags.
        """
        if meters is not None:
            pulumi.set(__self__, "meters", meters)
        if resource_groups is not None:
            pulumi.set(__self__, "resource_groups", resource_groups)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def meters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of filters on meters (GUID), mandatory for budgets of usage category. 
        """
        return pulumi.get(self, "meters")

    @meters.setter
    def meters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "meters", value)

    @property
    @pulumi.getter(name="resourceGroups")
    def resource_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of filters on resource groups, allowed at subscription level only.
        """
        return pulumi.get(self, "resource_groups")

    @resource_groups.setter
    def resource_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resource_groups", value)

    @property
    @pulumi.getter
    def resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of filters on resources.
        """
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "resources", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]]]:
        """
        The dictionary of filters on tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[str]]]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class NotificationArgs:
    def __init__(__self__, *,
                 contact_emails: pulumi.Input[Sequence[pulumi.Input[str]]],
                 enabled: pulumi.Input[bool],
                 operator: pulumi.Input[str],
                 threshold: pulumi.Input[float],
                 contact_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 contact_roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The notification associated with a budget.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contact_emails: Email addresses to send the budget notification to when the threshold is exceeded.
        :param pulumi.Input[bool] enabled: The notification is enabled or not.
        :param pulumi.Input[str] operator: The comparison operator.
        :param pulumi.Input[float] threshold: Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contact_groups: Action groups to send the budget notification to when the threshold is exceeded.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contact_roles: Contact roles to send the budget notification to when the threshold is exceeded.
        """
        pulumi.set(__self__, "contact_emails", contact_emails)
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "operator", operator)
        pulumi.set(__self__, "threshold", threshold)
        if contact_groups is not None:
            pulumi.set(__self__, "contact_groups", contact_groups)
        if contact_roles is not None:
            pulumi.set(__self__, "contact_roles", contact_roles)

    @property
    @pulumi.getter(name="contactEmails")
    def contact_emails(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Email addresses to send the budget notification to when the threshold is exceeded.
        """
        return pulumi.get(self, "contact_emails")

    @contact_emails.setter
    def contact_emails(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "contact_emails", value)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        The notification is enabled or not.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def operator(self) -> pulumi.Input[str]:
        """
        The comparison operator.
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: pulumi.Input[str]):
        pulumi.set(self, "operator", value)

    @property
    @pulumi.getter
    def threshold(self) -> pulumi.Input[float]:
        """
        Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
        """
        return pulumi.get(self, "threshold")

    @threshold.setter
    def threshold(self, value: pulumi.Input[float]):
        pulumi.set(self, "threshold", value)

    @property
    @pulumi.getter(name="contactGroups")
    def contact_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Action groups to send the budget notification to when the threshold is exceeded.
        """
        return pulumi.get(self, "contact_groups")

    @contact_groups.setter
    def contact_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "contact_groups", value)

    @property
    @pulumi.getter(name="contactRoles")
    def contact_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Contact roles to send the budget notification to when the threshold is exceeded.
        """
        return pulumi.get(self, "contact_roles")

    @contact_roles.setter
    def contact_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "contact_roles", value)

