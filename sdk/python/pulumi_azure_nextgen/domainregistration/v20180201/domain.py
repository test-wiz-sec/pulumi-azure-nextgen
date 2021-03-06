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

__all__ = ['Domain']


class Domain(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_code: Optional[pulumi.Input[str]] = None,
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 consent: Optional[pulumi.Input[pulumi.InputType['DomainPurchaseConsentArgs']]] = None,
                 contact_admin: Optional[pulumi.Input[pulumi.InputType['ContactArgs']]] = None,
                 contact_billing: Optional[pulumi.Input[pulumi.InputType['ContactArgs']]] = None,
                 contact_registrant: Optional[pulumi.Input[pulumi.InputType['ContactArgs']]] = None,
                 contact_tech: Optional[pulumi.Input[pulumi.InputType['ContactArgs']]] = None,
                 dns_type: Optional[pulumi.Input[str]] = None,
                 dns_zone_id: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 privacy: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_dns_type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Information about a domain.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_renew: <code>true</code> if the domain should be automatically renewed; otherwise, <code>false</code>.
        :param pulumi.Input[pulumi.InputType['DomainPurchaseConsentArgs']] consent: Legal agreement consent.
        :param pulumi.Input[pulumi.InputType['ContactArgs']] contact_admin: Administrative contact.
        :param pulumi.Input[pulumi.InputType['ContactArgs']] contact_billing: Billing contact.
        :param pulumi.Input[pulumi.InputType['ContactArgs']] contact_registrant: Registrant contact.
        :param pulumi.Input[pulumi.InputType['ContactArgs']] contact_tech: Technical contact.
        :param pulumi.Input[str] dns_type: Current DNS type
        :param pulumi.Input[str] dns_zone_id: Azure DNS Zone to use
        :param pulumi.Input[str] domain_name: Name of the domain.
        :param pulumi.Input[str] kind: Kind of resource.
        :param pulumi.Input[str] location: Resource Location.
        :param pulumi.Input[bool] privacy: <code>true</code> if domain privacy is enabled for this domain; otherwise, <code>false</code>.
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[str] target_dns_type: Target DNS type (would be used for migration)
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

            __props__['auth_code'] = auth_code
            __props__['auto_renew'] = auto_renew
            if consent is None:
                raise TypeError("Missing required property 'consent'")
            __props__['consent'] = consent
            if contact_admin is None:
                raise TypeError("Missing required property 'contact_admin'")
            __props__['contact_admin'] = contact_admin
            if contact_billing is None:
                raise TypeError("Missing required property 'contact_billing'")
            __props__['contact_billing'] = contact_billing
            if contact_registrant is None:
                raise TypeError("Missing required property 'contact_registrant'")
            __props__['contact_registrant'] = contact_registrant
            if contact_tech is None:
                raise TypeError("Missing required property 'contact_tech'")
            __props__['contact_tech'] = contact_tech
            __props__['dns_type'] = dns_type
            __props__['dns_zone_id'] = dns_zone_id
            if domain_name is None:
                raise TypeError("Missing required property 'domain_name'")
            __props__['domain_name'] = domain_name
            __props__['kind'] = kind
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['privacy'] = privacy
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['target_dns_type'] = target_dns_type
            __props__['created_time'] = None
            __props__['domain_not_renewable_reasons'] = None
            __props__['expiration_time'] = None
            __props__['last_renewed_time'] = None
            __props__['managed_host_names'] = None
            __props__['name'] = None
            __props__['name_servers'] = None
            __props__['provisioning_state'] = None
            __props__['ready_for_dns_record_management'] = None
            __props__['registration_status'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:domainregistration/latest:Domain"), pulumi.Alias(type_="azure-nextgen:domainregistration/v20150401:Domain"), pulumi.Alias(type_="azure-nextgen:domainregistration/v20150801:Domain"), pulumi.Alias(type_="azure-nextgen:domainregistration/v20190801:Domain"), pulumi.Alias(type_="azure-nextgen:domainregistration/v20200601:Domain")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Domain, __self__).__init__(
            'azure-nextgen:domainregistration/v20180201:Domain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Domain':
        """
        Get an existing Domain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Domain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authCode")
    def auth_code(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "auth_code")

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> pulumi.Output[Optional[bool]]:
        """
        <code>true</code> if the domain should be automatically renewed; otherwise, <code>false</code>.
        """
        return pulumi.get(self, "auto_renew")

    @property
    @pulumi.getter
    def consent(self) -> pulumi.Output['outputs.DomainPurchaseConsentResponse']:
        """
        Legal agreement consent.
        """
        return pulumi.get(self, "consent")

    @property
    @pulumi.getter(name="contactAdmin")
    def contact_admin(self) -> pulumi.Output['outputs.ContactResponse']:
        """
        Administrative contact.
        """
        return pulumi.get(self, "contact_admin")

    @property
    @pulumi.getter(name="contactBilling")
    def contact_billing(self) -> pulumi.Output['outputs.ContactResponse']:
        """
        Billing contact.
        """
        return pulumi.get(self, "contact_billing")

    @property
    @pulumi.getter(name="contactRegistrant")
    def contact_registrant(self) -> pulumi.Output['outputs.ContactResponse']:
        """
        Registrant contact.
        """
        return pulumi.get(self, "contact_registrant")

    @property
    @pulumi.getter(name="contactTech")
    def contact_tech(self) -> pulumi.Output['outputs.ContactResponse']:
        """
        Technical contact.
        """
        return pulumi.get(self, "contact_tech")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> pulumi.Output[str]:
        """
        Domain creation timestamp.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="dnsType")
    def dns_type(self) -> pulumi.Output[Optional[str]]:
        """
        Current DNS type
        """
        return pulumi.get(self, "dns_type")

    @property
    @pulumi.getter(name="dnsZoneId")
    def dns_zone_id(self) -> pulumi.Output[Optional[str]]:
        """
        Azure DNS Zone to use
        """
        return pulumi.get(self, "dns_zone_id")

    @property
    @pulumi.getter(name="domainNotRenewableReasons")
    def domain_not_renewable_reasons(self) -> pulumi.Output[Sequence[str]]:
        """
        Reasons why domain is not renewable.
        """
        return pulumi.get(self, "domain_not_renewable_reasons")

    @property
    @pulumi.getter(name="expirationTime")
    def expiration_time(self) -> pulumi.Output[str]:
        """
        Domain expiration timestamp.
        """
        return pulumi.get(self, "expiration_time")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[str]]:
        """
        Kind of resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="lastRenewedTime")
    def last_renewed_time(self) -> pulumi.Output[str]:
        """
        Timestamp when the domain was renewed last time.
        """
        return pulumi.get(self, "last_renewed_time")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource Location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedHostNames")
    def managed_host_names(self) -> pulumi.Output[Sequence['outputs.HostNameResponse']]:
        """
        All hostnames derived from the domain and assigned to Azure resources.
        """
        return pulumi.get(self, "managed_host_names")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameServers")
    def name_servers(self) -> pulumi.Output[Sequence[str]]:
        """
        Name servers.
        """
        return pulumi.get(self, "name_servers")

    @property
    @pulumi.getter
    def privacy(self) -> pulumi.Output[Optional[bool]]:
        """
        <code>true</code> if domain privacy is enabled for this domain; otherwise, <code>false</code>.
        """
        return pulumi.get(self, "privacy")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Domain provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="readyForDnsRecordManagement")
    def ready_for_dns_record_management(self) -> pulumi.Output[bool]:
        """
        <code>true</code> if Azure can assign this domain to App Service apps; otherwise, <code>false</code>. This value will be <code>true</code> if domain registration status is active and 
         it is hosted on name servers Azure has programmatic access to.
        """
        return pulumi.get(self, "ready_for_dns_record_management")

    @property
    @pulumi.getter(name="registrationStatus")
    def registration_status(self) -> pulumi.Output[str]:
        """
        Domain registration status.
        """
        return pulumi.get(self, "registration_status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetDnsType")
    def target_dns_type(self) -> pulumi.Output[Optional[str]]:
        """
        Target DNS type (would be used for migration)
        """
        return pulumi.get(self, "target_dns_type")

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

