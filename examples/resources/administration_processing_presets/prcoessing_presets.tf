resource "administration_processing_presets" "preset_1080p_25fps" {
  name      = "Up to 1080p, H.264, 25fps"
  pool_uuid = "pool_6lnvgxmv"
  # published = true

  video_medias = [{
    label = "test"
    codec = "h264"
    coder = "X264"
    resolution = {
      width  = 1280
      height = 720
    }
    bitrate   = 4500000
    framerate = "25"
    advanced = {

      # savc_config = {
      #   force_signal_level = false
      # }
    }
  }]

  audio_medias = [{
    codec      = "aac-lc"
    bitrate    = 96000
    samplerate = "48000"
    channels   = "2.0"
    track      = "eng"
    output     = "eng"
  }]

  subtitle_medias = [{
    track = "eng"
  }]
}
