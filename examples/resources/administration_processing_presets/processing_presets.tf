resource "administration_processing_presets" "preset_1080p_25fps" {
  org  = "orga_uuid"
  type = "standard"
  name = "Up to 1080p, H.264, 25fps"
  video_medias = [{
    codec = "h264"
    coder = "X264"
    resolution = {
      width  = 960
      height = 540
    }
    bitrate   = 2000000
    framerate = "25"
    },
  ]

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
