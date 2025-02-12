run "setup_tests" {
  module {
    source = "./tests/setup"
  }
}

run "create_presets" {
  module {
    source = "./tests/setup"
  }
  command = apply

  # first test
  assert {
    condition     = administration_processing_presets.preset_720p_25fps.type == null
    error_message = "Preset type should not be setted"
  }

  assert {
    condition     = administration_processing_presets.preset_720p_25fps.org == null
    error_message = "Preset org should not be setted"
  }

  assert {
    condition     = administration_processing_presets.preset_720p_25fps.name == "Up to 720p, H.264, 25fps"
    error_message = "Preset name is not correct"
  }

  assert {
    condition     = administration_processing_presets.preset_720p_25fps.video_medias[0].codec == "h264"
    error_message = "Video codec is not correct"
  }

  assert {
    condition     = administration_processing_presets.preset_720p_25fps.audio_medias[0].bitrate == 96000
    error_message = "Audio bitrate is not correct"
  }

  assert {
    condition     = administration_processing_presets.preset_720p_25fps.uuid == "${run.setup_tests.preset_uuid}"
    error_message = "Preset UUID is not correct"
  }

  assert {
    condition     = administration_processing_presets.preset_720p_25fps.subtitle_medias[0].track == "eng"
    error_message = "Subtitle track is not correct"
  }

  #second data test

  assert {
    condition     = administration_processing_presets.preset_with_type.type == "standard"
    error_message = "Preset type is not correct"
  }

  assert {
    condition     = administration_processing_presets.preset_with_type.name == "test_preset_standard"
    error_message = "Preset name is not correct"
  }

  assert {
    condition     = administration_processing_presets.preset_with_type.video_medias[0].codec == "h264"
    error_message = "Video codec is not correct"
  }

  # Adding assertion for audio media bitrate
  assert {
    condition     = administration_processing_presets.preset_with_type.audio_medias[0].bitrate == 96000
    error_message = "Audio bitrate is not correct"
  }

  # Adding assertion for subtitle track
  assert {
    condition     = administration_processing_presets.preset_with_type.subtitle_medias[0].track == "eng"
    error_message = "Subtitle track is not correct"
  }

  assert {
    condition     = administration_processing_presets.preset_with_type.uuid == "${run.setup_tests.preset_with_type_uuid}"
    error_message = "Preset UUID is not correct"
  }

  # third data test
  assert {
    condition     = administration_processing_presets.preset_with_org_and_type.type == "standard"
    error_message = "Preset type is not correct"
  }

  assert {
    condition     = administration_processing_presets.preset_with_org_and_type.name == "test_preset_standard_with_org"
    error_message = "Preset name is not correct"
  }

  assert {
    condition     = administration_processing_presets.preset_with_org_and_type.video_medias[0].codec == "h264"
    error_message = "Video codec is not correct"
  }

  assert {
    condition     = administration_processing_presets.preset_with_org_and_type.org == var.my_local_test_org
    error_message = "Preset org should be set to the local test org"
  }

  assert {
    condition     = administration_processing_presets.preset_with_org_and_type.uuid == "${run.setup_tests.preset_with_org_and_type_uuid}"
    error_message = "Preset UUID is not correct"
  }

}

