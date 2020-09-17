# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'ActivityLogAlertActionGroupResponse',
    'ActivityLogAlertActionListResponse',
    'ActivityLogAlertAllOfConditionResponse',
    'ActivityLogAlertLeafConditionResponse',
    'AutomationRunbookReceiverResponse',
    'AzureAppPushReceiverResponse',
    'EmailReceiverResponse',
    'ItsmReceiverResponse',
    'SmsReceiverResponse',
    'WebhookReceiverResponse',
]

@pulumi.output_type
class ActivityLogAlertActionGroupResponse(dict):
    """
    A pointer to an Azure Action Group.
    """
    def __init__(__self__, *,
                 action_group_id: str,
                 webhook_properties: Optional[Mapping[str, str]] = None):
        """
        A pointer to an Azure Action Group.
        :param str action_group_id: The resourceId of the action group. This cannot be null or empty.
        :param Mapping[str, str] webhook_properties: the dictionary of custom properties to include with the post operation. These data are appended to the webhook payload.
        """
        pulumi.set(__self__, "action_group_id", action_group_id)
        if webhook_properties is not None:
            pulumi.set(__self__, "webhook_properties", webhook_properties)

    @property
    @pulumi.getter(name="actionGroupId")
    def action_group_id(self) -> str:
        """
        The resourceId of the action group. This cannot be null or empty.
        """
        return pulumi.get(self, "action_group_id")

    @property
    @pulumi.getter(name="webhookProperties")
    def webhook_properties(self) -> Optional[Mapping[str, str]]:
        """
        the dictionary of custom properties to include with the post operation. These data are appended to the webhook payload.
        """
        return pulumi.get(self, "webhook_properties")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ActivityLogAlertActionListResponse(dict):
    """
    A list of activity log alert actions.
    """
    def __init__(__self__, *,
                 action_groups: Optional[Sequence['outputs.ActivityLogAlertActionGroupResponse']] = None):
        """
        A list of activity log alert actions.
        :param Sequence['ActivityLogAlertActionGroupResponseArgs'] action_groups: The list of activity log alerts.
        """
        if action_groups is not None:
            pulumi.set(__self__, "action_groups", action_groups)

    @property
    @pulumi.getter(name="actionGroups")
    def action_groups(self) -> Optional[Sequence['outputs.ActivityLogAlertActionGroupResponse']]:
        """
        The list of activity log alerts.
        """
        return pulumi.get(self, "action_groups")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ActivityLogAlertAllOfConditionResponse(dict):
    """
    An Activity Log alert condition that is met when all its member conditions are met.
    """
    def __init__(__self__, *,
                 all_of: Sequence['outputs.ActivityLogAlertLeafConditionResponse']):
        """
        An Activity Log alert condition that is met when all its member conditions are met.
        :param Sequence['ActivityLogAlertLeafConditionResponseArgs'] all_of: The list of activity log alert conditions.
        """
        pulumi.set(__self__, "all_of", all_of)

    @property
    @pulumi.getter(name="allOf")
    def all_of(self) -> Sequence['outputs.ActivityLogAlertLeafConditionResponse']:
        """
        The list of activity log alert conditions.
        """
        return pulumi.get(self, "all_of")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ActivityLogAlertLeafConditionResponse(dict):
    """
    An Activity Log alert condition that is met by comparing an activity log field and value.
    """
    def __init__(__self__, *,
                 equals: str,
                 field: str):
        """
        An Activity Log alert condition that is met by comparing an activity log field and value.
        :param str equals: The field value will be compared to this value (case-insensitive) to determine if the condition is met.
        :param str field: The name of the field that this condition will examine. The possible values for this field are (case-insensitive): 'resourceId', 'category', 'caller', 'level', 'operationName', 'resourceGroup', 'resourceProvider', 'status', 'subStatus', 'resourceType', or anything beginning with 'properties.'.
        """
        pulumi.set(__self__, "equals", equals)
        pulumi.set(__self__, "field", field)

    @property
    @pulumi.getter
    def equals(self) -> str:
        """
        The field value will be compared to this value (case-insensitive) to determine if the condition is met.
        """
        return pulumi.get(self, "equals")

    @property
    @pulumi.getter
    def field(self) -> str:
        """
        The name of the field that this condition will examine. The possible values for this field are (case-insensitive): 'resourceId', 'category', 'caller', 'level', 'operationName', 'resourceGroup', 'resourceProvider', 'status', 'subStatus', 'resourceType', or anything beginning with 'properties.'.
        """
        return pulumi.get(self, "field")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AutomationRunbookReceiverResponse(dict):
    """
    The Azure Automation Runbook notification receiver.
    """
    def __init__(__self__, *,
                 automation_account_id: str,
                 is_global_runbook: bool,
                 runbook_name: str,
                 webhook_resource_id: str,
                 name: Optional[str] = None,
                 service_uri: Optional[str] = None):
        """
        The Azure Automation Runbook notification receiver.
        :param str automation_account_id: The Azure automation account Id which holds this runbook and authenticate to Azure resource.
        :param bool is_global_runbook: Indicates whether this instance is global runbook.
        :param str runbook_name: The name for this runbook.
        :param str webhook_resource_id: The resource id for webhook linked to this runbook.
        :param str name: Indicates name of the webhook.
        :param str service_uri: The URI where webhooks should be sent.
        """
        pulumi.set(__self__, "automation_account_id", automation_account_id)
        pulumi.set(__self__, "is_global_runbook", is_global_runbook)
        pulumi.set(__self__, "runbook_name", runbook_name)
        pulumi.set(__self__, "webhook_resource_id", webhook_resource_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if service_uri is not None:
            pulumi.set(__self__, "service_uri", service_uri)

    @property
    @pulumi.getter(name="automationAccountId")
    def automation_account_id(self) -> str:
        """
        The Azure automation account Id which holds this runbook and authenticate to Azure resource.
        """
        return pulumi.get(self, "automation_account_id")

    @property
    @pulumi.getter(name="isGlobalRunbook")
    def is_global_runbook(self) -> bool:
        """
        Indicates whether this instance is global runbook.
        """
        return pulumi.get(self, "is_global_runbook")

    @property
    @pulumi.getter(name="runbookName")
    def runbook_name(self) -> str:
        """
        The name for this runbook.
        """
        return pulumi.get(self, "runbook_name")

    @property
    @pulumi.getter(name="webhookResourceId")
    def webhook_resource_id(self) -> str:
        """
        The resource id for webhook linked to this runbook.
        """
        return pulumi.get(self, "webhook_resource_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Indicates name of the webhook.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceUri")
    def service_uri(self) -> Optional[str]:
        """
        The URI where webhooks should be sent.
        """
        return pulumi.get(self, "service_uri")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AzureAppPushReceiverResponse(dict):
    """
    The Azure mobile App push notification receiver.
    """
    def __init__(__self__, *,
                 email_address: str,
                 name: str):
        """
        The Azure mobile App push notification receiver.
        :param str email_address: The email address registered for the Azure mobile app.
        :param str name: The name of the Azure mobile app push receiver. Names must be unique across all receivers within an action group.
        """
        pulumi.set(__self__, "email_address", email_address)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="emailAddress")
    def email_address(self) -> str:
        """
        The email address registered for the Azure mobile app.
        """
        return pulumi.get(self, "email_address")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Azure mobile app push receiver. Names must be unique across all receivers within an action group.
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EmailReceiverResponse(dict):
    """
    An email receiver.
    """
    def __init__(__self__, *,
                 email_address: str,
                 name: str,
                 status: str):
        """
        An email receiver.
        :param str email_address: The email address of this receiver.
        :param str name: The name of the email receiver. Names must be unique across all receivers within an action group.
        :param str status: The receiver status of the e-mail.
        """
        pulumi.set(__self__, "email_address", email_address)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="emailAddress")
    def email_address(self) -> str:
        """
        The email address of this receiver.
        """
        return pulumi.get(self, "email_address")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the email receiver. Names must be unique across all receivers within an action group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The receiver status of the e-mail.
        """
        return pulumi.get(self, "status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ItsmReceiverResponse(dict):
    """
    An Itsm receiver.
    """
    def __init__(__self__, *,
                 connection_id: str,
                 name: str,
                 region: str,
                 ticket_configuration: str,
                 workspace_id: str):
        """
        An Itsm receiver.
        :param str connection_id: Unique identification of ITSM connection among multiple defined in above workspace.
        :param str name: The name of the Itsm receiver. Names must be unique across all receivers within an action group.
        :param str region: Region in which workspace resides. Supported values:'centralindia','japaneast','southeastasia','australiasoutheast','uksouth','westcentralus','canadacentral','eastus','westeurope'
        :param str ticket_configuration: JSON blob for the configurations of the ITSM action. CreateMultipleWorkItems option will be part of this blob as well.
        :param str workspace_id: OMS LA instance identifier.
        """
        pulumi.set(__self__, "connection_id", connection_id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "ticket_configuration", ticket_configuration)
        pulumi.set(__self__, "workspace_id", workspace_id)

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> str:
        """
        Unique identification of ITSM connection among multiple defined in above workspace.
        """
        return pulumi.get(self, "connection_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Itsm receiver. Names must be unique across all receivers within an action group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        Region in which workspace resides. Supported values:'centralindia','japaneast','southeastasia','australiasoutheast','uksouth','westcentralus','canadacentral','eastus','westeurope'
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="ticketConfiguration")
    def ticket_configuration(self) -> str:
        """
        JSON blob for the configurations of the ITSM action. CreateMultipleWorkItems option will be part of this blob as well.
        """
        return pulumi.get(self, "ticket_configuration")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> str:
        """
        OMS LA instance identifier.
        """
        return pulumi.get(self, "workspace_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SmsReceiverResponse(dict):
    """
    An SMS receiver.
    """
    def __init__(__self__, *,
                 country_code: str,
                 name: str,
                 phone_number: str,
                 status: str):
        """
        An SMS receiver.
        :param str country_code: The country code of the SMS receiver.
        :param str name: The name of the SMS receiver. Names must be unique across all receivers within an action group.
        :param str phone_number: The phone number of the SMS receiver.
        :param str status: The status of the receiver.
        """
        pulumi.set(__self__, "country_code", country_code)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "phone_number", phone_number)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="countryCode")
    def country_code(self) -> str:
        """
        The country code of the SMS receiver.
        """
        return pulumi.get(self, "country_code")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the SMS receiver. Names must be unique across all receivers within an action group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="phoneNumber")
    def phone_number(self) -> str:
        """
        The phone number of the SMS receiver.
        """
        return pulumi.get(self, "phone_number")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the receiver.
        """
        return pulumi.get(self, "status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class WebhookReceiverResponse(dict):
    """
    A webhook receiver.
    """
    def __init__(__self__, *,
                 name: str,
                 service_uri: str):
        """
        A webhook receiver.
        :param str name: The name of the webhook receiver. Names must be unique across all receivers within an action group.
        :param str service_uri: The URI where webhooks should be sent.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "service_uri", service_uri)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the webhook receiver. Names must be unique across all receivers within an action group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceUri")
    def service_uri(self) -> str:
        """
        The URI where webhooks should be sent.
        """
        return pulumi.get(self, "service_uri")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

