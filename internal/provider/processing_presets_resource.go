package provider

import (
	"context"
	"fmt"
	"time"

	"terraform-provider-administration/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Resolution struct {
	Width  types.Int64 `tfsdk:"width"`
	Height types.Int64 `tfsdk:"height"`
}

type SavcConfig struct {
	ForceSignalLevel     types.Bool    `tfsdk:"force_signal_level"`
	BufsizeRatio         types.Float64 `tfsdk:"bufsize_ratio"`
	RcInitOccupancy      types.Float64 `tfsdk:"rc_init_occupancy"`
	QualitySpeedOverride types.String  `tfsdk:"quality_speed_override"`
}

// Advanced -
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
	Maxrate             types.Int64  `tfsdk:"maxrate"`
	KeyFrameIntervalMs  types.Int64  `tfsdk:"key_frame_interval_ms"`
	HorizontalSharpness types.Int64  `tfsdk:"horizontal_sharpness"`
	VerticalSharpness   types.Int64  `tfsdk:"vertical_sharpness"`
	LogoEnabled         types.Bool   `tfsdk:"logo_enabled"`
	SavcConfig          SavcConfig   `tfsdk:"savc_config"`
}

type VideoMedia struct {
	Label      types.String `tfsdk:"label"`
	Codec      types.String `tfsdk:"codec"`
	Coder      types.String `tfsdk:"coder"`
	Resolution Resolution   `tfsdk:"resolution"`
	Bitrate    types.Int64  `tfsdk:"bitrate"`
	Framerate  types.String `tfsdk:"framerate"`
	Advanced   Advanced     `tfsdk:"advanced"`
}

type AudioMedia struct {
	Codec            types.String `tfsdk:"codec"`
	Label            types.String `tfsdk:"label"`
	Bitrate          types.Int64  `tfsdk:"bitrate"`
	Samplerate       types.String `tfsdk:"samplerate"`
	Channels         types.String `tfsdk:"channels"`
	Track            types.String `tfsdk:"track"`
	Output           types.String `tfsdk:"output"`
	AudioDescription types.Bool   `tfsdk:"audio_description"`
}

type SubtitleMedia struct {
	Track                types.String `tfsdk:"track"`
	Output               types.String `tfsdk:"output"`
	DeafAndHardOfHearing types.Bool   `tfsdk:"deaf_and_hard_of_hearing"`
}

type processingPresetsResourceModel struct {
	Uuid           types.String    `tfsdk:"uuid"`
	Name           types.String    `tfsdk:"name"`
	Identifier     types.String    `tfsdk:"identifier"`
	Published      types.Bool      `tfsdk:"published"`
	PoolUuid       types.String    `tfsdk:"pool_uuid"`
	VideoMedias    []VideoMedia    `tfsdk:"video_medias"`
	AudioMedias    []AudioMedia    `tfsdk:"audio_medias"`
	SubtitleMedias []SubtitleMedia `tfsdk:"subtitle_medias"`
	Labels         []types.String  `tfsdk:"labels"`
	ModifiedAt     types.String    `tfsdk:"modified_at"`
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
			"uuid": schema.StringAttribute{
				Description: "UUID of the processing presets.",
				Computed:    true,
			},
			"modified_at": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the processing presets.",
				Required:    true,
			},
			"identifier": schema.StringAttribute{
				Description: "Identifier of the processing presets.",
				Computed:    true,
			},
			"published": schema.BoolAttribute{
				Description: "Published status of the processing presets.",
				Computed:    true,
			},
			"pool_uuid": schema.StringAttribute{
				Description: "Pool UUID of the processing presets.",
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
						},
						"codec": schema.StringAttribute{
							Description: "Codec of the video media.",
							Required:    true,
						},
						"coder": schema.StringAttribute{
							Description: "Coder of the video media.",
							Required:    true,
						},
						"resolution": schema.SingleNestedAttribute{
							Description: "Resolution of the video media.",
							Required:    true,
							Attributes: map[string]schema.Attribute{
								"width": schema.Int64Attribute{
									Description: "Width of the resolution.",
									Optional:    true,
								},
								"height": schema.Int64Attribute{
									Description: "Height of the resolution.",
									Optional:    true,
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
						"advanced": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"profile": schema.StringAttribute{
									Description: "Profile of the video media.",
									Optional:    true,
								},
								"level": schema.StringAttribute{
									Description: "Level of the video media.",
									Optional:    true,
								},
								"quality": schema.StringAttribute{
									Description: "Quality of the video media.",
									Optional:    true,
								},
								"encoding_mode": schema.StringAttribute{
									Description: "Encoding mode of the video media.",
									Optional:    true,
								},
								"encoding_quality": schema.Int64Attribute{
									Description: "Encoding quality of the video media.",
									Optional:    true,
								},
								"quality_optimization": schema.StringAttribute{
									Description: "Quality optimization of the video media.",
									Optional:    true,
								},
								"closed_gop": schema.BoolAttribute{
									Description: "Closed gop of the video media.",
									Optional:    true,
								},
								"gop_size": schema.Int64Attribute{
									Description: "Gop size of the video media.",
									Optional:    true,
								},
								"gop_max_size": schema.Int64Attribute{
									Description: "Gop max size of the video media.",
									Optional:    true,
								},
								"bframe": schema.BoolAttribute{
									Description: "Bframe of the video media.",
									Optional:    true,
								},
								"bframe_number": schema.Int64Attribute{
									Description: "Bframe number of the video media.",
									Optional:    true,
								},
								"key_frame_interval_ms": schema.Int64Attribute{
									Description: "Key frame interval ms of the video media.",
									Optional:    true,
								},
								"horizontal_sharpness": schema.Int64Attribute{
									Description: "Horizontal sharpness of the video media.",
									Optional:    true,
								},
								"vertical_sharpness": schema.Int64Attribute{
									Description: "Vertical sharpness of the video media.",
									Optional:    true,
								},
								"logo_enabled": schema.BoolAttribute{
									Description: "Logo enabled of the video media.",
									Optional:    true,
								},
								"savc_config": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"force_signal": schema.BoolAttribute{
											Description: "Force signal of the video media.",
											Optional:    true,
										},
										"bufsize_ratio": schema.Float64Attribute{
											Description: "Bufsize ratio of the video media.",
											Optional:    true,
										},
										"rc_init_occupancy": schema.Float64Attribute{
											Description: "Rc init occupancy of the video media.",
											Optional:    true,
										},
										"quality_speed_override": schema.StringAttribute{
											Description: "Quality speed override of the video media.",
											Optional:    true,
										},
									},
								},
							},
						},
					},
				},
			},
			"audio_medias": schema.ListNestedAttribute{
				Description: "List of audio medias of the processing presets.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"label": schema.StringAttribute{
							Description: "Label of the audio media.",
							Optional:    true,
						},
						"codec": schema.StringAttribute{
							Description: "Codec of the audio media.",
							Optional:    true,
						},
						"bitrate": schema.Int64Attribute{
							Description: "Bitrate of the audio media.",
							Optional:    true,
						},
						"samplerate": schema.Int64Attribute{
							Description: "Sample rate of the audio media.",
							Optional:    true,
						},
						"channels": schema.Int64Attribute{
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
						},
						"audio_description": schema.BoolAttribute{
							Description: "Audio description of the audio media.",
							Optional:    true,
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
						},
						"output": schema.StringAttribute{
							Description: "Output of the subtitle media.",
							Optional:    true,
						},
						"deaf_and_hard_of_hearing": schema.BoolAttribute{
							Description: "Deaf and hard of hearing of the subtitle media.",
							Optional:    true,
						},
					},
				},
			},
			"labels": schema.ListAttribute{
				Description: "List of labels of the processing presets.",
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func ProcessingPresetsModelToProcessingPresets(processing processingPresetsResourceModel) *client.ProcessingPresets {
	newProcessing := client.ProcessingPresets{
		Uuid:      processing.Uuid.ValueString(),
		Name:      processing.Name.ValueString(),
		Published: processing.Published.ValueBool(),
		PoolUuid:  processing.PoolUuid.ValueString(),
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

			// advanceditem: videoMediaItem.Advanced,

			Advanced: client.Advanced{
				Profile:             videoMediaItem.Advanced.Profile.ValueString(),
				Level:               videoMediaItem.Advanced.Level.ValueString(),
				Quality:             videoMediaItem.Advanced.Quality.ValueString(),
				EncodingMode:        videoMediaItem.Advanced.EncodingMode.ValueString(),
				EncodingQuality:     int(videoMediaItem.Advanced.EncodingQuality.ValueInt64()),
				QualityOptimization: videoMediaItem.Advanced.QualityOptimization.ValueString(),
				ClosedGop:           videoMediaItem.Advanced.ClosedGop.ValueBool(),
				GopSize:             int(videoMediaItem.Advanced.GopSize.ValueInt64()),
				GopMaxSize:          int(videoMediaItem.Advanced.GopMaxSize.ValueInt64()),
				Bframe:              videoMediaItem.Advanced.Bframe.ValueBool(),
				BframeNumber:        int(videoMediaItem.Advanced.BframeNumber.ValueInt64()),
				KeyFrameIntervalMs:  int(videoMediaItem.Advanced.KeyFrameIntervalMs.ValueInt64()),
				HorizontalSharpness: int(videoMediaItem.Advanced.HorizontalSharpness.ValueInt64()),
				VerticalSharpness:   int(videoMediaItem.Advanced.VerticalSharpness.ValueInt64()),
				LogoEnabled:         videoMediaItem.Advanced.LogoEnabled.ValueBool(),

				SavcConfig: client.SavcConfig{
					ForceSignalLevel:     videoMediaItem.Advanced.SavcConfig.ForceSignalLevel.ValueBool(),
					BufsizeRatio:         videoMediaItem.Advanced.SavcConfig.BufsizeRatio.ValueFloat64(),
					RcInitOccupancy:      videoMediaItem.Advanced.SavcConfig.RcInitOccupancy.ValueFloat64(),
					QualitySpeedOverride: videoMediaItem.Advanced.SavcConfig.QualitySpeedOverride.ValueString(),
				},
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
			AudioDescription: audioMediaItem.AudioDescription.ValueBool(),
		})
	}

	var subtitleMedias = []client.SubtitleMedia{}
	for _, subtitleMediaItem := range processing.SubtitleMedias {
		subtitleMedias = append(subtitleMedias, client.SubtitleMedia{
			Track:                subtitleMediaItem.Track.ValueString(),
			Output:               subtitleMediaItem.Output.ValueString(),
			DeafAndHardOfHearing: subtitleMediaItem.DeafAndHardOfHearing.ValueBool(),
		})
	}

	newProcessing.VideoMedias = videoMedias
	newProcessing.AudioMedias = audioMedias
	newProcessing.SubtitleMedias = subtitleMedias
	newProcessing.Labels = []string{}
	for _, labelItem := range processing.Labels {
		newProcessing.Labels = append(newProcessing.Labels, labelItem.ValueString())
	}

	return &newProcessing
}

func ProcessingPresetsToProcessingPresetsModel(processing client.ProcessingPresets, model *processingPresetsResourceModel) {
	model.Uuid = types.StringValue(processing.Uuid)
	model.Name = types.StringValue(processing.Name)
	model.Published = types.BoolValue(processing.Published)
	model.PoolUuid = types.StringValue(processing.PoolUuid)

	model.VideoMedias = []VideoMedia{}
	for _, videoMediaItem := range processing.VideoMedias {
		model.VideoMedias = append(model.VideoMedias, VideoMedia{
			Label:     types.StringValue(videoMediaItem.Label),
			Codec:     types.StringValue(videoMediaItem.Codec),
			Coder:     types.StringValue(videoMediaItem.Coder),
			Bitrate:   types.Int64Value(int64(videoMediaItem.Bitrate)),
			Framerate: types.StringValue(videoMediaItem.Framerate),

			Resolution: Resolution{
				Width:  types.Int64Value(int64(videoMediaItem.Resolution.Width)),
				Height: types.Int64Value(int64(videoMediaItem.Resolution.Height)),
			},

			Advanced: Advanced{
				Profile:             types.StringValue(videoMediaItem.Advanced.Profile),
				Level:               types.StringValue(videoMediaItem.Advanced.Level),
				Quality:             types.StringValue(videoMediaItem.Advanced.Quality),
				EncodingMode:        types.StringValue(videoMediaItem.Advanced.EncodingMode),
				EncodingQuality:     types.Int64Value(int64(videoMediaItem.Advanced.EncodingQuality)),
				QualityOptimization: types.StringValue(videoMediaItem.Advanced.QualityOptimization),
				ClosedGop:           types.BoolValue(videoMediaItem.Advanced.ClosedGop),
				GopSize:             types.Int64Value(int64(videoMediaItem.Advanced.GopSize)),
				GopMaxSize:          types.Int64Value(int64(videoMediaItem.Advanced.GopMaxSize)),
				Bframe:              types.BoolValue(videoMediaItem.Advanced.Bframe),
				BframeNumber:        types.Int64Value(int64(videoMediaItem.Advanced.BframeNumber)),
				KeyFrameIntervalMs:  types.Int64Value(int64(videoMediaItem.Advanced.KeyFrameIntervalMs)),
				HorizontalSharpness: types.Int64Value(int64(videoMediaItem.Advanced.HorizontalSharpness)),
				VerticalSharpness:   types.Int64Value(int64(videoMediaItem.Advanced.VerticalSharpness)),
				LogoEnabled:         types.BoolValue(videoMediaItem.Advanced.LogoEnabled),
				SavcConfig: SavcConfig{
					ForceSignalLevel:     types.BoolValue(videoMediaItem.Advanced.SavcConfig.ForceSignalLevel),
					BufsizeRatio:         types.Float64Value(videoMediaItem.Advanced.SavcConfig.BufsizeRatio),
					RcInitOccupancy:      types.Float64Value(videoMediaItem.Advanced.SavcConfig.RcInitOccupancy),
					QualitySpeedOverride: types.StringValue(videoMediaItem.Advanced.SavcConfig.QualitySpeedOverride),
				},
			},
		})
	}

	model.AudioMedias = []AudioMedia{}

	for _, audioMediaItem := range processing.AudioMedias {
		model.AudioMedias = append(model.AudioMedias, AudioMedia{
			Label:            types.StringValue(audioMediaItem.Label),
			Codec:            types.StringValue(audioMediaItem.Codec),
			Channels:         types.StringValue(audioMediaItem.Channels),
			Bitrate:          types.Int64Value(int64(audioMediaItem.Bitrate)),
			Samplerate:       types.StringValue(audioMediaItem.Samplerate),
			Track:            types.StringValue(audioMediaItem.Track),
			Output:           types.StringValue(audioMediaItem.Output),
			AudioDescription: types.BoolValue(audioMediaItem.AudioDescription),
		})
	}

	model.SubtitleMedias = []SubtitleMedia{}
	for _, subtitleMediaItem := range processing.SubtitleMedias {
		model.SubtitleMedias = append(model.SubtitleMedias, SubtitleMedia{
			Track:                types.StringValue(subtitleMediaItem.Track),
			Output:               types.StringValue(subtitleMediaItem.Output),
			DeafAndHardOfHearing: types.BoolValue(subtitleMediaItem.DeafAndHardOfHearing),
		})
	}

	model.Labels = []types.String{}
	for _, labelItem := range processing.Labels {
		model.Labels = append(model.Labels, types.StringValue(labelItem))
	}

	model.ModifiedAt = types.StringValue(time.Now().Format(time.RFC850))
}

func (r *processingPresetsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var processing processingPresetsResourceModel
	diags := req.Plan.Get(ctx, &processing)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	newProcessingPresets := ProcessingPresetsModelToProcessingPresets(processing)

	rproc, err := r.client.CreateManageProcessingPresets(*newProcessingPresets)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to create processing presets.",
			"Could not create processing presets, unexpected error: "+err.Error(),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	processing.Identifier = types.StringValue(rproc.Identifier)
	processing.Name = types.StringValue(rproc.Name)
	processing.Uuid = types.StringValue(rproc.Uuid)

	ProcessingPresetsToProcessingPresetsModel(*rproc, &processing)

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

	rproc, err := r.client.GetManageProcessingPresets(processing.Uuid.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to read processing presets.",
			"Could not read processing presets, unexpected error: "+err.Error(),
		)
		return
	}

	ProcessingPresetsToProcessingPresetsModel(*rproc, &processing)

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

	newProcessingPresets := ProcessingPresetsModelToProcessingPresets(processing)

	rproc, err := r.client.UpdateManageProcessingPresets(processing.Uuid.ValueString(), *newProcessingPresets)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to update processing presets.",
			"Could not update processing presets, unexpected error: "+err.Error(),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	processing.Identifier = types.StringValue(rproc.Identifier)
	processing.Name = types.StringValue(rproc.Name)
	processing.Uuid = types.StringValue(rproc.Uuid)

	ProcessingPresetsToProcessingPresetsModel(*rproc, &processing)

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

	err := r.client.DeleteManageProcessingPresets(processing.Uuid.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to delete processing presets.",
			"Could not delete processing presets, unexpected error: "+err.Error(),
		)
		return
	}
}

// TODO do i need that ?
func (r *processingPresetsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// TODO do i need that ?
func (r *processingPresetsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
