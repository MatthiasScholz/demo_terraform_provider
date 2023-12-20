package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = (*demotpgenProvider)(nil)

func New() func() provider.Provider {
	return func() provider.Provider {
		return &demotpgenProvider{}
	}
}

type demotpgenProvider struct{}

func (p *demotpgenProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {

}

func (p *demotpgenProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

func (p *demotpgenProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "demotpgen"
}

func (p *demotpgenProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *demotpgenProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}
