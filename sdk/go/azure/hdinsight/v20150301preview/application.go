// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150301preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The HDInsight cluster application
type Application struct {
	pulumi.CustomResourceState

	// The ETag for the application
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the application.
	Properties ApplicationPropertiesResponseOutput `pulumi:"properties"`
	// The tags for the application.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil || args.ApplicationName == nil {
		return nil, errors.New("missing required argument 'ApplicationName'")
	}
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ApplicationArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:hdinsight/v20180601preview:Application"),
		},
	})
	opts = append(opts, aliases)
	var resource Application
	err := ctx.RegisterResource("azure-nextgen:hdinsight/v20150301preview:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azure-nextgen:hdinsight/v20150301preview:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// The ETag for the application
	Etag *string `pulumi:"etag"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The properties of the application.
	Properties *ApplicationPropertiesResponse `pulumi:"properties"`
	// The tags for the application.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type ApplicationState struct {
	// The ETag for the application
	Etag pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The properties of the application.
	Properties ApplicationPropertiesResponsePtrInput
	// The tags for the application.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// The constant value for the application name.
	ApplicationName string `pulumi:"applicationName"`
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// The ETag for the application
	Etag *string `pulumi:"etag"`
	// The properties of the application.
	Properties *ApplicationProperties `pulumi:"properties"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags for the application.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// The constant value for the application name.
	ApplicationName pulumi.StringInput
	// The name of the cluster.
	ClusterName pulumi.StringInput
	// The ETag for the application
	Etag pulumi.StringPtrInput
	// The properties of the application.
	Properties ApplicationPropertiesPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The tags for the application.
	Tags pulumi.StringMapInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}
