# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'AppServiceCertificateResponse',
    'CertificateDetailsResponse',
    'CertificateOrderCertificateResponse',
]

@pulumi.output_type
class AppServiceCertificateResponse(dict):
    """
    Key Vault container for a certificate that is purchased through Azure.
    """
    def __init__(__self__, *,
                 provisioning_state: str,
                 key_vault_id: Optional[str] = None,
                 key_vault_secret_name: Optional[str] = None):
        """
        Key Vault container for a certificate that is purchased through Azure.
        :param str provisioning_state: Status of the Key Vault secret.
        :param str key_vault_id: Key Vault resource Id.
        :param str key_vault_secret_name: Key Vault secret name.
        """
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if key_vault_id is not None:
            pulumi.set(__self__, "key_vault_id", key_vault_id)
        if key_vault_secret_name is not None:
            pulumi.set(__self__, "key_vault_secret_name", key_vault_secret_name)

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Status of the Key Vault secret.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> Optional[str]:
        """
        Key Vault resource Id.
        """
        return pulumi.get(self, "key_vault_id")

    @property
    @pulumi.getter(name="keyVaultSecretName")
    def key_vault_secret_name(self) -> Optional[str]:
        """
        Key Vault secret name.
        """
        return pulumi.get(self, "key_vault_secret_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CertificateDetailsResponse(dict):
    """
    Certificate Details
    """
    def __init__(__self__, *,
                 location: str,
                 id: Optional[str] = None,
                 issuer: Optional[str] = None,
                 kind: Optional[str] = None,
                 name: Optional[str] = None,
                 not_after: Optional[str] = None,
                 not_before: Optional[str] = None,
                 raw_data: Optional[str] = None,
                 serial_number: Optional[str] = None,
                 signature_algorithm: Optional[str] = None,
                 subject: Optional[str] = None,
                 tags: Optional[Mapping[str, str]] = None,
                 thumbprint: Optional[str] = None,
                 type: Optional[str] = None,
                 version: Optional[int] = None):
        """
        Certificate Details
        :param str location: Resource Location
        :param str id: Resource Id
        :param str issuer: Issuer
        :param str kind: Kind of resource
        :param str name: Resource Name
        :param str not_after: Valid to
        :param str not_before: Valid from
        :param str raw_data: Raw certificate data
        :param str serial_number: Serial Number
        :param str signature_algorithm: Signature Algorithm
        :param str subject: Subject
        :param Mapping[str, str] tags: Resource tags
        :param str thumbprint: Thumbprint
        :param str type: Resource type
        :param int version: Version
        """
        pulumi.set(__self__, "location", location)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if issuer is not None:
            pulumi.set(__self__, "issuer", issuer)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if not_after is not None:
            pulumi.set(__self__, "not_after", not_after)
        if not_before is not None:
            pulumi.set(__self__, "not_before", not_before)
        if raw_data is not None:
            pulumi.set(__self__, "raw_data", raw_data)
        if serial_number is not None:
            pulumi.set(__self__, "serial_number", serial_number)
        if signature_algorithm is not None:
            pulumi.set(__self__, "signature_algorithm", signature_algorithm)
        if subject is not None:
            pulumi.set(__self__, "subject", subject)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if thumbprint is not None:
            pulumi.set(__self__, "thumbprint", thumbprint)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource Location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Resource Id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def issuer(self) -> Optional[str]:
        """
        Issuer
        """
        return pulumi.get(self, "issuer")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind of resource
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Resource Name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notAfter")
    def not_after(self) -> Optional[str]:
        """
        Valid to
        """
        return pulumi.get(self, "not_after")

    @property
    @pulumi.getter(name="notBefore")
    def not_before(self) -> Optional[str]:
        """
        Valid from
        """
        return pulumi.get(self, "not_before")

    @property
    @pulumi.getter(name="rawData")
    def raw_data(self) -> Optional[str]:
        """
        Raw certificate data
        """
        return pulumi.get(self, "raw_data")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> Optional[str]:
        """
        Serial Number
        """
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter(name="signatureAlgorithm")
    def signature_algorithm(self) -> Optional[str]:
        """
        Signature Algorithm
        """
        return pulumi.get(self, "signature_algorithm")

    @property
    @pulumi.getter
    def subject(self) -> Optional[str]:
        """
        Subject
        """
        return pulumi.get(self, "subject")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def thumbprint(self) -> Optional[str]:
        """
        Thumbprint
        """
        return pulumi.get(self, "thumbprint")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> Optional[int]:
        """
        Version
        """
        return pulumi.get(self, "version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CertificateOrderCertificateResponse(dict):
    """
    Class representing the Key Vault container for certificate purchased through Azure
    """
    def __init__(__self__, *,
                 location: str,
                 id: Optional[str] = None,
                 key_vault_id: Optional[str] = None,
                 key_vault_secret_name: Optional[str] = None,
                 kind: Optional[str] = None,
                 name: Optional[str] = None,
                 provisioning_state: Optional[str] = None,
                 tags: Optional[Mapping[str, str]] = None,
                 type: Optional[str] = None):
        """
        Class representing the Key Vault container for certificate purchased through Azure
        :param str location: Resource Location
        :param str id: Resource Id
        :param str key_vault_id: Key Vault Csm resource Id
        :param str key_vault_secret_name: Key Vault secret name
        :param str kind: Kind of resource
        :param str name: Resource Name
        :param str provisioning_state: Status of the Key Vault secret
        :param Mapping[str, str] tags: Resource tags
        :param str type: Resource type
        """
        pulumi.set(__self__, "location", location)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if key_vault_id is not None:
            pulumi.set(__self__, "key_vault_id", key_vault_id)
        if key_vault_secret_name is not None:
            pulumi.set(__self__, "key_vault_secret_name", key_vault_secret_name)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if provisioning_state is not None:
            pulumi.set(__self__, "provisioning_state", provisioning_state)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource Location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        Resource Id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="keyVaultId")
    def key_vault_id(self) -> Optional[str]:
        """
        Key Vault Csm resource Id
        """
        return pulumi.get(self, "key_vault_id")

    @property
    @pulumi.getter(name="keyVaultSecretName")
    def key_vault_secret_name(self) -> Optional[str]:
        """
        Key Vault secret name
        """
        return pulumi.get(self, "key_vault_secret_name")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind of resource
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Resource Name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        Status of the Key Vault secret
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


