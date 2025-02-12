package provider

import (
	"context"
	"fmt"
	"log"

	"terraform-provider-administration/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Resolution struct {
	Width  types.Int64 `tfsdk:"width"`
	Height types.Int64 `tfsdk:"height"`
}

/* type SavcConfig struct {
 	ForceSignalLevel     types.Bool    `tfsdk:"force_signal_level"`
 	BufsizeRatio         types.Float64 `tfsdk:"bufsize_ratio"`
 	RcInitOccupancy      types.Float64 `tfsdk:"rc_init_occupancy"`
 	QualitySpeedOverride types.String  `tfsdk:"quality_speed_override"`
}.*/

type Advanced struct {
	Profile             types.String `tfsdk:"profile"`
	Level               types.String `tfsdk:"level"`
	Quality             types.String `tfsdk:"quality"`
	EncodingMode        types.String `tfsdk:"encoding_mode"`
	EncodingQuality     types.Int64  `tfsdk:"encoding_quality"`
	QualityOptimization types.String `tfsdk:"quality_optimization"`
	ClosedGop           types.Bool   `tfsdk:"closed_gop"`
	GopSize             types.Int64  `tfsdk:"gop_size"`
	GopMaxSize          types.Int64  `tfsdk:"gop_max_size"`
	Bframe              types.Bool   `tfsdk:"bframe"`
	BframeNumber        types.Int64  `tfsdk:"bframe_number"`
	KeyFrameIntervalMs  types.Int64  `tfsdk:"key_frame_interval_ms"`
	HorizontalSharpness types.Int64  `tfsdk:"horizontal_sharpness"`
	VerticalSharpness   types.Int64  `tfsdk:"vertical_sharpness"`
	LogoEnabled         types.Bool   `tfsdk:"logo_enabled"`
	// SavcConfig          SavcConfig   `tfsdk:"savc_config"`.
}

type VideoMedia struct {
	Label      types.String `tfsdk:"label"`
	Codec      types.String `tfsdk:"codec"`
	Coder      types.String `tfsdk:"coder"`
	Resolution Resolution   `tfsdk:"resolution"`
	Bitrate    types.Int64  `tfsdk:"bitrate"`
	Framerate  types.String `tfsdk:"framerate"`
	// Advanced   Advanced     `tfsdk:"advanced"`.
}

type AudioMedia struct {
	Codec            types.String `tfsdk:"codec"`
	Channels         types.String `tfsdk:"channels"`
	Bitrate          types.Int64  `tfsdk:"bitrate"`
	Samplerate       types.String `tfsdk:"samplerate"`
	Track            types.String `tfsdk:"track"`
	Output           types.String `tfsdk:"output"`
	OutputLabel      types.String `tfsdk:"output_label"`
	AudioDescription types.Bool   `tfsdk:"audio_description"`
	Label            types.String `tfsdk:"label"`
}

type SubtitleMedia struct {
	Track                types.String `tfsdk:"track"`
	Bitrate              types.Int64  `tfsdk:"bitrate"`
	Output               types.String `tfsdk:"output"`
	OutputLabel          types.String `tfsdk:"output_label"`
	DeafAndHardOfHearing types.Bool   `tfsdk:"deaf_and_hard_of_hearing"`
}

type processingPresetsResourceModel struct {
	// Params on api resources.
	Organization types.String `tfsdk:"org"`
	Type         types.String `tfsdk:"type"`
	// actual fields.
	Uuid           types.String    `tfsdk:"uuid"`
	Name           types.String    `tfsdk:"name"`
	VideoMedias    []VideoMedia    `tfsdk:"video_medias"`
	AudioMedias    []AudioMedia    `tfsdk:"audio_medias"`
	SubtitleMedias []SubtitleMedia `tfsdk:"subtitle_medias"`
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &processingPresetsResource{}
	_ resource.ResourceWithConfigure   = &processingPresetsResource{}
	_ resource.ResourceWithImportState = &processingPresetsResource{}
)

// NewPlanResource is a helper function to simplify the provider implementation.
func NewProcessingPresetsResource() resource.Resource {
	return &processingPresetsResource{}
}

// planResource is the resource implementation.
type processingPresetsResource struct {
	client *client.Client
}

// Metadata returns the resource type name.
func (r *processingPresetsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_processing_presets"
}

// Schema defines the schema for the resource.
func (r *processingPresetsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a processing presets.",
		Attributes: map[string]schema.Attribute{
			"org": schema.StringAttribute{
				Description: "Organization UUID of the processing presets needed for admin API.",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "Type of the processing presets.",
				Optional:    true,
			},
			"uuid": schema.StringAttribute{
				Description: "UUID of the processing presets.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the processing presets.",
				Required:    true,
			},
			"video_medias": schema.ListNestedAttribute{
				Description: "List of video medias of the processing presets.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"label": schema.StringAttribute{
							Description: "Label of the video media.",
							Optional:    true,
							Computed:    true,
						},
						"codec": schema.StringAttribute{
							Description: "Codec of the video media.",
							Required:    true,
						},
						"coder": schema.StringAttribute{
							Description: "Coder of the video media.",
							Optional:    true,
							Computed:    true,
						},
						"resolution": schema.SingleNestedAttribute{
							Description: "Resolution of the video media.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"width": schema.Int64Attribute{
									Description: "Width of the resolution.",
									Required:    true,
								},
								"height": schema.Int64Attribute{
									Description: "Height of the resolution.",
									Required:    true,
								},
							},
						},
						"bitrate": schema.Int64Attribute{
							Description: "Bitrate of the video media.",
							Required:    true,
						},
						"framerate": schema.StringAttribute{
							Description: "Framerate of the video media.",
							Optional:    true,
						},
						// "advanced": schema.SingleNestedAttribute{
						// 	Optional: true,
						// 	Required: false,
						// 	Attributes: map[string]schema.Attribute{
						// 		"profile": schema.StringAttribute{
						// 			Description: "Profile of the video media.",
						// 			Optional:    true,
						// 		},
						// 		"level": schema.StringAttribute{
						// 			Description: "Level of the video media.",
						// 			Optional:    true,
						// 		},
						// 		"quality": schema.StringAttribute{
						// 			Description: "Quality of the video media.",
						// 			Optional:    true,
						// 		},
						// 		"encoding_mode": schema.StringAttribute{
						// 			Description: "Encoding mode of the video media.",
						// 			Optional:    true,
						// 		},
						// 		"encoding_quality": schema.Int64Attribute{
						// 			Description: "Encoding quality of the video media.",
						// 			Optional:    true,
						// 		},
						// 		"quality_optimization": schema.StringAttribute{
						// 			Description: "Quality optimization of the video media.",
						// 			Optional:    true,
						// 		},
						// 		"closed_gop": schema.BoolAttribute{
						// 			Description: "Closed gop of the video media.",
						// 			Optional:    true,
						// 		},
						// 		"gop_size": schema.Int64Attribute{
						// 			Description: "Gop size of the video media.",
						// 			Optional:    true,
						// 		},
						// 		"gop_max_size": schema.Int64Attribute{
						// 			Description: "Gop max size of the video media.",
						// 			Optional:    true,
						// 		},
						// 		"bframe": schema.BoolAttribute{
						// 			Description: "Bframe of the video media.",
						// 			Optional:    true,
						// 		},
						// 		"bframe_number": schema.Int64Attribute{
						// 			Description: "Bframe number of the video media.",
						// 			Optional:    true,
						// 		},
						// 		"key_frame_interval_ms": schema.Int64Attribute{
						// 			Description: "Key frame interval ms of the video media.",
						// 			Optional:    true,
						// 		},
						// 		"horizontal_sharpness": schema.Int64Attribute{
						// 			Description: "Horizontal sharpness of the video media.",
						// 			Optional:    true,
						// 		},
						// 		"vertical_sharpness": schema.Int64Attribute{
						// 			Description: "Vertical sharpness of the video media.",
						// 			Optional:    true,
						// 		},
						// 		"logo_enabled": schema.BoolAttribute{
						// 			Description: "Logo enabled of the video media.",
						// 			Optional:    true,
						// 		},
						// "savc_config": schema.SingleNestedAttribute{
						// 	Optional: true,
						// 	Attributes: map[string]schema.Attribute{
						// 		"force_signal_level": schema.BoolAttribute{
						// 			Description: "Force signal of the video media.",
						// 			Optional:    true,
						// 		},
						// 		"bufsize_ratio": schema.Float64Attribute{
						// 			Description: "Bufsize ratio of the video media.",
						// 			Optional:    true,
						// 		},
						// 		"rc_init_occupancy": schema.Float64Attribute{
						// 			Description: "Rc init occupancy of the video media.",
						// 			Optional:    true,
						// 		},
						// 		"quality_speed_override": schema.StringAttribute{
						// 			Description: "Quality speed override of the video media.",
						// 			Optional:    true,
						// 		},
						// 	},
						// },
						// },.
					},
				},
			},
			// },.
			"audio_medias": schema.ListNestedAttribute{
				Description: "List of audio medias of the processing presets.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"codec": schema.StringAttribute{
							Description: "Codec of the audio media.",
							Optional:    true,
						},
						"bitrate": schema.Int64Attribute{
							Description: "Bitrate of the audio media.",
							Optional:    true,
						},
						"samplerate": schema.StringAttribute{
							Description: "Sample rate of the audio media.",
							Optional:    true,
						},
						"channels": schema.StringAttribute{
							Description: "Channels of the audio media.",
							Optional:    true,
						},
						"track": schema.StringAttribute{
							Description: "Track of the audio media.",
							Optional:    true,
						},
						"output": schema.StringAttribute{
							Description: "Output of the audio media.",
							Optional:    true,
							Computed:    true,
						},
						"output_label": schema.StringAttribute{
							Description: "Output label of the audio media.",
							Optional:    true,
							Computed:    true,
						},
						"audio_description": schema.BoolAttribute{
							Description: "Audio description of the audio media.",
							Optional:    true,
							Computed:    true,
						},
						"label": schema.StringAttribute{
							Description: "Label of the audio media.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"subtitle_medias": schema.ListNestedAttribute{
				Description: "List of subtitle medias of the processing presets.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"track": schema.StringAttribute{
							Description: "Track of the subtitle media.",
							Optional:    true,
							Computed:    true,
						},
						"bitrate": schema.Int64Attribute{
							Description: "Bitrate of the subtitle media.",
							Optional:    true,
							Computed:    true,
						},
						"output": schema.StringAttribute{
							Description: "Output of the subtitle media.",
							Optional:    true,
							Computed:    true,
						},
						"output_label": schema.StringAttribute{
							Description: "Output label of the subtitle media.",
							Optional:    true,
							Computed:    true,
						},
						"deaf_and_hard_of_hearing": schema.BoolAttribute{
							Description: "Deaf and hard of hearing of the subtitle media.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func ProcessingPresetsModelToProcessingPresetsJson(processing processingPresetsResourceModel) *client.ProcessingPresets {

	newProcessing := client.ProcessingPresets{
		Uuid: processing.Uuid.ValueString(),
		Name: processing.Name.ValueString(),
	}

	var videoMedias = []client.VideoMedia{}
	for _, videoMediaItem := range processing.VideoMedias {
		videoMedias = append(videoMedias, client.VideoMedia{
			Codec:     videoMediaItem.Codec.ValueString(),
			Coder:     videoMediaItem.Coder.ValueString(),
			Bitrate:   int(videoMediaItem.Bitrate.ValueInt64()),
			Framerate: videoMediaItem.Framerate.ValueString(),
			Resolution: client.Resolution{
				Width:  int(videoMediaItem.Resolution.Width.ValueInt64()),
				Height: int(videoMediaItem.Resolution.Height.ValueInt64()),
			},
		})
	}

	var audioMedias = []client.AudioMedia{}
	for _, audioMediaItem := range processing.AudioMedias {
		audioMedias = append(audioMedias, client.AudioMedia{
			Codec:            audioMediaItem.Codec.ValueString(),
			Channels:         audioMediaItem.Channels.ValueString(),
			Bitrate:          int(audioMediaItem.Bitrate.ValueInt64()),
			Samplerate:       audioMediaItem.Samplerate.ValueString(),
			Track:            audioMediaItem.Track.ValueString(),
			Output:           audioMediaItem.Output.ValueString(),
			OutputLabel:      audioMediaItem.OutputLabel.ValueString(),
			AudioDescription: audioMediaItem.AudioDescription.ValueBool(),
			Label:            audioMediaItem.Label.ValueString(),
		})
	}

	var subtitleMedias = []client.SubtitleMedia{}
	for _, subtitleMediaItem := range processing.SubtitleMedias {
		subtitleMedias = append(subtitleMedias, client.SubtitleMedia{
			Track:                subtitleMediaItem.Track.ValueString(),
			Bitrate:              int(subtitleMediaItem.Bitrate.ValueInt64()),
			Output:               subtitleMediaItem.Output.ValueString(),
			OutputLabel:          subtitleMediaItem.OutputLabel.ValueString(),
			DeafAndHardOfHearing: subtitleMediaItem.DeafAndHardOfHearing.ValueBool(),
		})
	}

	newProcessing.VideoMedias = videoMedias
	newProcessing.AudioMedias = audioMedias
	newProcessing.SubtitleMedias = subtitleMedias

	return &newProcessing
}

func ProcessingPresetsJsonToProcessingPresetsModel(processing client.ProcessingPresets, model *processingPresetsResourceModel) {

	model.Uuid = types.StringValue(processing.Uuid)
	model.Name = types.StringValue(processing.Name)

	model.VideoMedias = []VideoMedia{}
	for _, videoMediaItem := range processing.VideoMedias {

		model.VideoMedias = append(model.VideoMedias, VideoMedia{
			Label:     types.StringValue(videoMediaItem.Label),
			Coder:     types.StringValue(videoMediaItem.Coder),
			Codec:     types.StringValue(videoMediaItem.Codec),
			Bitrate:   types.Int64Value(int64(videoMediaItem.Bitrate)),
			Framerate: types.StringValue(videoMediaItem.Framerate),

			Resolution: Resolution{
				Width:  types.Int64Value(int64(videoMediaItem.Resolution.Width)),
				Height: types.Int64Value(int64(videoMediaItem.Resolution.Height)),
			},
		})
	}

	model.AudioMedias = []AudioMedia{}

	for _, audioMediaItem := range processing.AudioMedias {
		model.AudioMedias = append(model.AudioMedias, AudioMedia{
			Codec:            types.StringValue(audioMediaItem.Codec),
			Channels:         types.StringValue(audioMediaItem.Channels),
			Bitrate:          types.Int64Value(int64(audioMediaItem.Bitrate)),
			Samplerate:       types.StringValue(audioMediaItem.Samplerate),
			Track:            types.StringValue(audioMediaItem.Track),
			Output:           types.StringValue(audioMediaItem.Output),
			OutputLabel:      types.StringValue(audioMediaItem.OutputLabel),
			AudioDescription: types.BoolValue(audioMediaItem.AudioDescription),
			Label:            types.StringValue(audioMediaItem.Label),
		})
	}

	model.SubtitleMedias = []SubtitleMedia{}
	for _, subtitleMediaItem := range processing.SubtitleMedias {
		model.SubtitleMedias = append(model.SubtitleMedias, SubtitleMedia{
			Track:                types.StringValue(subtitleMediaItem.Track),
			Bitrate:              types.Int64Value(int64(subtitleMediaItem.Bitrate)),
			Output:               types.StringValue(subtitleMediaItem.Output),
			OutputLabel:          types.StringValue(subtitleMediaItem.OutputLabel),
			DeafAndHardOfHearing: types.BoolValue(subtitleMediaItem.DeafAndHardOfHearing),
		})
	}
}

func (r *processingPresetsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	log.Printf("Create processing presets")
	var processing processingPresetsResourceModel
	diags := req.Plan.Get(ctx, &processing)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	newProcessingPresets := ProcessingPresetsModelToProcessingPresetsJson(processing)

	rproc, err := r.client.CreateManageProcessingPresets(*newProcessingPresets, processing.Organization.ValueString(), processing.Type.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to create processing presets.",
			"Could not create processing presets, unexpected error: "+err.Error(),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values.
	ProcessingPresetsJsonToProcessingPresetsModel(*rproc, &processing)

	diags = resp.State.Set(ctx, processing)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *processingPresetsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var processing processingPresetsResourceModel
	diags := req.State.Get(ctx, &processing)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	rproc, err := r.client.GetManageProcessingPresets(processing.Uuid.ValueString(), processing.Organization.ValueString(), processing.Type.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to read processing presets.",
			"Could not read processing presets with UUID "+processing.Uuid.ValueString()+", unexpected error: "+err.Error(),
		)
		return
	}

	ProcessingPresetsJsonToProcessingPresetsModel(*rproc, &processing)

	diags = resp.State.Set(ctx, processing)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *processingPresetsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var processing processingPresetsResourceModel
	diags := req.Plan.Get(ctx, &processing)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	newProcessingPresets := ProcessingPresetsModelToProcessingPresetsJson(processing)

	_, err := r.client.UpdateManageProcessingPresets(processing.Uuid.ValueString(), *newProcessingPresets, processing.Organization.ValueString(), processing.Type.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to update processing presets.",
			"Could not update processing presets, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated items from GetOrder as UpdateOrder items are not
	// populated.
	rproc, err := r.client.GetManageProcessingPresets(processing.Uuid.ValueString(), processing.Organization.ValueString(), processing.Type.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Administration Plan",
			"Could not read processing presets with UUID "+processing.Uuid.ValueString()+", unexpected error: "+err.Error(),
		)
		return
	}
	// Map response body to schema and populate Computed attribute values.
	ProcessingPresetsJsonToProcessingPresetsModel(*rproc, &processing)

	diags = resp.State.Set(ctx, processing)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *processingPresetsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var processing processingPresetsResourceModel
	diags := req.State.Get(ctx, &processing)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteManageProcessingPresets(processing.Uuid.ValueString(), processing.Organization.ValueString(), processing.Type.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to delete processing presets.",
			"Could not delete processing presets, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *processingPresetsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to uuid attribute.
	resource.ImportStatePassthroughID(ctx, path.Root("uuid"), req, resp)
}

func (r *processingPresetsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform.
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
