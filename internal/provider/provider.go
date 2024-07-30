package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"terraform-provider-administration/internal/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &administrationProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &administrationProvider{
			version: version,
		}
	}
}

// administrationProvider is the provider implementation.
type administrationProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// administrationProviderModel maps provider schema data to a Go type.
type administrationProviderModel struct {
	AuthServer   types.String `tfsdk:"auth_server"`
	Host         types.String `tfsdk:"host"`
	ClientId     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
}

// Metadata returns the provider type name.
func (p *administrationProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "administration"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *administrationProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with Administration.",
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Description: "Host for Administration API. May also be provided via ADMINISTRATION_HOST environment variable.",
				Optional:    true,
			},
			"auth_server": schema.StringAttribute{
				Description: "Auth server for Administration API. May also be provided via ADMINISTRATION_AUTH_SERVER environment variable.",
				Optional:    true,
			},
			"client_id": schema.StringAttribute{
				Description: "ClientId for Administration API. May also be provided via ADMINISTRATION_CLIENT_ID environment variable.",
				Required:    true,
			},
			"client_secret": schema.StringAttribute{
				Description: "ClientSecret for Administration API. May also be provided via ADMINISTRATION_CLIENT_SECRET environment variable.",
				Required:    true,
				Sensitive:   true,
			},
		},
	}
}

func (p *administrationProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	tflog.Info(ctx, "Configuring Administration client")
	// Retrieve provider data from configuration
	var config administrationProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.
	if config.AuthServer.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("auth_server"),
			"Unknown Administration API Auth Server",
			"The provider cannot create the Administration API client as there is an unknown configuration value for the Administration API auth_server. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the ADMINISTRATION_AUTH_SERVER environment variable.",
		)
	}

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Administration API Host",
			"The provider cannot create the Administration API client as there is an unknown configuration value for the Administration API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the ADMINISTRATION_HOST environment variable.",
		)
	}

	if config.ClientId.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_id"),
			"Unknown Administration API ClientId",
			"The provider cannot create the Administration API client as there is an unknown configuration value for the Administration API client_id. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the ADMINISTRATION_CLIENT_ID environment variable.",
		)
	}

	if config.ClientSecret.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_secret"),
			"Unknown Administration API ClientSecret",
			"The provider cannot create the Administration API client as there is an unknown configuration value for the Administration API client_secret. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the ADMINISTRATION_CLIENT_SECRET environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	auth_server := os.Getenv("ADMINISTRATION_AUTH_SERVER")
	host := os.Getenv("ADMINISTRATION_HOST")
	client_id := os.Getenv("ADMINISTRATION_CLIENT_ID")
	client_secret := os.Getenv("ADMINISTRATION_CLIENT_SECRET")

	if !config.AuthServer.IsNull() {
		auth_server = config.AuthServer.ValueString()
	}

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.ClientId.IsNull() {
		client_id = config.ClientId.ValueString()
	}

	if !config.ClientSecret.IsNull() {
		client_secret = config.ClientSecret.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.
	if auth_server == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("auth_server"),
			"Missing Administration API Auth Server",
			"The provider cannot create the Administration API client as there is a missing or empty value for the Administration API auth_server. "+
				"Set the auth_server value in the configuration or use the ADMINISTRATION_AUTH_SERVER environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Administration API Host",
			"The provider cannot create the Administration API client as there is a missing or empty value for the Administration API host. "+
				"Set the host value in the configuration or use the ADMINISTRATION_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if client_id == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_id"),
			"Missing Administration API ClientId",
			"The provider cannot create the Administration API client as there is a missing or empty value for the Administration API client_id. "+
				"Set the client_id value in the configuration or use the ADMINISTRATION_CLIENT_ID environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if client_secret == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_secret"),
			"Missing Administration API ClientSecret",
			"The provider cannot create the Administration API client as there is a missing or empty value for the Administration API client_secret. "+
				"Set the client_secret value in the configuration or use the ADMINISTRATION_CLIENT_SECRET environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "administration_auth_server", auth_server)
	ctx = tflog.SetField(ctx, "administration_host", host)
	ctx = tflog.SetField(ctx, "administration_client_id", client_id)
	ctx = tflog.SetField(ctx, "administration_client_secret", client_secret)

	tflog.Debug(ctx, "Creating Administration client")

	// Create a new Administration client using the configuration values
	client, err := client.NewClient(&auth_server, &host, &client_id, &client_secret)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Administration API Client",
			"An unexpected error occurred when creating the Administration API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Administration Client Error: "+err.Error(),
		)
		return
	}

	// Make the Administration client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Info(ctx, "Configured Administration client", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *administrationProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *administrationProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewPlanResource,
	}
}
