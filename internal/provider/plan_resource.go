package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-administration/internal/client"
)

type limitItemModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.Int64  `tfsdk:"value"`
}

type pricingItemModel struct {
	SubscribeForYear     types.Int64   `tfsdk:"subscribe_for_year"`
	MonthlyPrice         types.Float64 `tfsdk:"monthly_price"`
	MonthlyPriceCurrency types.String  `tfsdk:"monthly_price_currency"`
}

type planResourceModel struct {
	ID          types.String       `tfsdk:"id"`
	Name        types.String       `tfsdk:"name"`
	LastUpdated types.String       `tfsdk:"last_updated"`
	Features    []types.String     `tfsdk:"features"`
	Limits      []limitItemModel   `tfsdk:"limits"`
	Pricing     []pricingItemModel `tfsdk:"pricing"`
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &planResource{}
	_ resource.ResourceWithConfigure   = &planResource{}
	_ resource.ResourceWithImportState = &planResource{}
)

// NewPlanResource is a helper function to simplify the provider implementation.
func NewPlanResource() resource.Resource {
	return &planResource{}
}

// planResource is the resource implementation.
type planResource struct {
	client *client.Client
}

// Metadata returns the resource type name.
func (r *planResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_billing_plan"
}

// Schema defines the schema for the resource.
func (r *planResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a plan.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Numeric identifier of the plan.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the plan.",
				Required:    true,
			},
			"features": schema.ListAttribute{
				Description: "List of features of the plan.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"limits": schema.ListNestedAttribute{
				Description: "List of limits of the plan.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "Name of limit.",
							Required:    true,
						},
						"value": schema.Int64Attribute{
							Description: "Value of limit.",
							Required:    true,
						},
					},
				},
			},
			"pricing": schema.ListNestedAttribute{
				Description: "List of pricing of the plan.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"subscribe_for_year": schema.Int64Attribute{
							Description: "Number of year of subscription.",
							Required:    true,
						},
						"monthly_price": schema.Float64Attribute{
							Description: "Monthly pricing.",
							Required:    true,
						},
						"monthly_price_currency": schema.StringAttribute{
							Description: "Monthly currency.",
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func PlanModelToPlan(plan planResourceModel) *client.Plan {
	newPlan := client.Plan{
		Name: plan.Name.ValueString(),
	}
	var features = []string{}
	for _, item := range plan.Features {
		features = append(features, item.ValueString())
	}

	var limits = []client.LimitsItem{}
	for _, item := range plan.Limits {
		limits = append(limits, client.LimitsItem{
			Name:  item.Name.ValueString(),
			Value: int(item.Value.ValueInt64()),
		})
	}

	var pricings = []client.PrincingItem{}
	for _, item := range plan.Pricing {
		pricings = append(pricings, client.PrincingItem{
			SubscribeForYear:     int(item.SubscribeForYear.ValueInt64()),
			MonthlyPrice:         item.MonthlyPrice.ValueFloat64(),
			MonthlyPriceCurrency: item.MonthlyPriceCurrency.ValueString(),
		})
	}

	newPlan.Features = features
	newPlan.Limits = limits
	newPlan.Pricing = pricings
	return &newPlan
}

func PlanToPlanModel(plan client.Plan, model *planResourceModel) {
	model.Name = types.StringValue(plan.Name)
	model.Features = []types.String{}
	for _, ftrItem := range plan.Features {
		model.Features = append(model.Features, types.StringValue(ftrItem))
	}

	model.Limits = []limitItemModel{}
	for _, limitItem := range plan.Limits {
		model.Limits = append(model.Limits, limitItemModel{
			Name:  types.StringValue(limitItem.Name),
			Value: types.Int64Value(int64(limitItem.Value)),
		})
	}

	model.Pricing = []pricingItemModel{}
	for _, pricingItem := range plan.Pricing {
		model.Pricing = append(model.Pricing, pricingItemModel{
			SubscribeForYear:     types.Int64Value(int64(pricingItem.SubscribeForYear)),
			MonthlyPrice:         types.Float64Value(pricingItem.MonthlyPrice),
			MonthlyPriceCurrency: types.StringValue(pricingItem.MonthlyPriceCurrency),
		})
	}

	model.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
}

// Create a new resource.
func (r *planResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan planResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	newPlan := PlanModelToPlan(plan)

	// Create new plan
	rplan, err := r.client.CreatePlan(*newPlan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating plan",
			"Could not create plan, unexpected error: "+err.Error(),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.ID = types.StringValue(strconv.Itoa(rplan.ID))
	PlanToPlanModel(*rplan, &plan)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *planResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state planResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from Administration
	rplan, err := r.client.GetPlan(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Administration Plan",
			"Could not read Administration plan ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite items with refreshed state
	state.ID = types.StringValue(strconv.Itoa(rplan.ID))
	PlanToPlanModel(*rplan, &state)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *planResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan planResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	newPlan := PlanModelToPlan(plan)

	// Update existing order
	_, err := r.client.UpdatePlan(plan.ID.ValueString(), *newPlan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Administration Plan",
			"Could not update plan, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated items from GetOrder as UpdateOrder items are not
	// populated.
	rplan, err := r.client.GetPlan(plan.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Administration Plan",
			"Could not read Administration plan ID "+plan.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Update resource state with updated items and timestamp
	PlanToPlanModel(*rplan, &plan)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *planResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state planResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing order
	err := r.client.DeletePlan(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Administration Plan",
			"Could not delete plan, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *planResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Configure adds the provider configured client to the resource.
func (r *planResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}
