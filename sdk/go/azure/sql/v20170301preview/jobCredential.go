// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A stored credential that can be used by a job to connect to target databases.
type JobCredential struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The credential password.
	Password pulumi.StringOutput `pulumi:"password"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The credential user name.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewJobCredential registers a new resource with the given unique name, arguments, and options.
func NewJobCredential(ctx *pulumi.Context,
	name string, args *JobCredentialArgs, opts ...pulumi.ResourceOption) (*JobCredential, error) {
	if args == nil || args.CredentialName == nil {
		return nil, errors.New("missing required argument 'CredentialName'")
	}
	if args == nil || args.JobAgentName == nil {
		return nil, errors.New("missing required argument 'JobAgentName'")
	}
	if args == nil || args.Password == nil {
		return nil, errors.New("missing required argument 'Password'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil || args.Username == nil {
		return nil, errors.New("missing required argument 'Username'")
	}
	if args == nil {
		args = &JobCredentialArgs{}
	}
	var resource JobCredential
	err := ctx.RegisterResource("azure-nextgen:sql/v20170301preview:JobCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobCredential gets an existing JobCredential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobCredentialState, opts ...pulumi.ResourceOption) (*JobCredential, error) {
	var resource JobCredential
	err := ctx.ReadResource("azure-nextgen:sql/v20170301preview:JobCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobCredential resources.
type jobCredentialState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// The credential password.
	Password *string `pulumi:"password"`
	// Resource type.
	Type *string `pulumi:"type"`
	// The credential user name.
	Username *string `pulumi:"username"`
}

type JobCredentialState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// The credential password.
	Password pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
	// The credential user name.
	Username pulumi.StringPtrInput
}

func (JobCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobCredentialState)(nil)).Elem()
}

type jobCredentialArgs struct {
	// The name of the credential.
	CredentialName string `pulumi:"credentialName"`
	// The name of the job agent.
	JobAgentName string `pulumi:"jobAgentName"`
	// The credential password.
	Password string `pulumi:"password"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The credential user name.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a JobCredential resource.
type JobCredentialArgs struct {
	// The name of the credential.
	CredentialName pulumi.StringInput
	// The name of the job agent.
	JobAgentName pulumi.StringInput
	// The credential password.
	Password pulumi.StringInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The credential user name.
	Username pulumi.StringInput
}

func (JobCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobCredentialArgs)(nil)).Elem()
}
