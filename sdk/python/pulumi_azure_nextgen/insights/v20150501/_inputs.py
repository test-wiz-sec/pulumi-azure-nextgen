# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'ApplicationInsightsComponentAnalyticsItemPropertiesArgs',
    'WebTestGeolocationArgs',
    'WebTestPropertiesConfigurationArgs',
]

@pulumi.input_type
class ApplicationInsightsComponentAnalyticsItemPropertiesArgs:
    def __init__(__self__, *,
                 function_alias: Optional[pulumi.Input[str]] = None):
        """
        A set of properties that can be defined in the context of a specific item type. Each type may have its own properties.
        :param pulumi.Input[str] function_alias: A function alias, used when the type of the item is Function
        """
        if function_alias is not None:
            pulumi.set(__self__, "function_alias", function_alias)

    @property
    @pulumi.getter(name="functionAlias")
    def function_alias(self) -> Optional[pulumi.Input[str]]:
        """
        A function alias, used when the type of the item is Function
        """
        return pulumi.get(self, "function_alias")

    @function_alias.setter
    def function_alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_alias", value)


@pulumi.input_type
class WebTestGeolocationArgs:
    def __init__(__self__, *,
                 location: Optional[pulumi.Input[str]] = None):
        """
        Geo-physical location to run a web test from. You must specify one or more locations for the test to run from.
        :param pulumi.Input[str] location: Location ID for the webtest to run from.
        """
        if location is not None:
            pulumi.set(__self__, "location", location)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Location ID for the webtest to run from.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)


@pulumi.input_type
class WebTestPropertiesConfigurationArgs:
    def __init__(__self__, *,
                 web_test: Optional[pulumi.Input[str]] = None):
        """
        An XML configuration specification for a WebTest.
        :param pulumi.Input[str] web_test: The XML specification of a WebTest to run against an application.
        """
        if web_test is not None:
            pulumi.set(__self__, "web_test", web_test)

    @property
    @pulumi.getter(name="webTest")
    def web_test(self) -> Optional[pulumi.Input[str]]:
        """
        The XML specification of a WebTest to run against an application.
        """
        return pulumi.get(self, "web_test")

    @web_test.setter
    def web_test(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "web_test", value)


