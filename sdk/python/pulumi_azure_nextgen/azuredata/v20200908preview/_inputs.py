# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'OnPremisePropertyArgs',
]

@pulumi.input_type
class OnPremisePropertyArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 public_signing_key: pulumi.Input[str],
                 signing_certificate_thumbprint: Optional[pulumi.Input[str]] = None):
        """
        Properties from the on premise data controller
        :param pulumi.Input[str] id: A globally unique ID identifying the associated on premise cluster
        :param pulumi.Input[str] public_signing_key: Certificate that contains the on premise cluster public key used to verify signing
        :param pulumi.Input[str] signing_certificate_thumbprint: Unique thumbprint returned to customer to verify the certificate being uploaded
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "public_signing_key", public_signing_key)
        if signing_certificate_thumbprint is not None:
            pulumi.set(__self__, "signing_certificate_thumbprint", signing_certificate_thumbprint)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        A globally unique ID identifying the associated on premise cluster
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="publicSigningKey")
    def public_signing_key(self) -> pulumi.Input[str]:
        """
        Certificate that contains the on premise cluster public key used to verify signing
        """
        return pulumi.get(self, "public_signing_key")

    @public_signing_key.setter
    def public_signing_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "public_signing_key", value)

    @property
    @pulumi.getter(name="signingCertificateThumbprint")
    def signing_certificate_thumbprint(self) -> Optional[pulumi.Input[str]]:
        """
        Unique thumbprint returned to customer to verify the certificate being uploaded
        """
        return pulumi.get(self, "signing_certificate_thumbprint")

    @signing_certificate_thumbprint.setter
    def signing_certificate_thumbprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "signing_certificate_thumbprint", value)


