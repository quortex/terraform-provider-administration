terraform {
  required_providers {
    administration = {
      version = "0.0.4"
      # source  = "quortex/administration"
      source = "localhost/quortex/administration"
    }
  }
}

provider "administration" {
  auth_server = "https://auth.dev.saas-dev.quortex.io"
  host        = "http://localhost:8000"
  #host = 
  client_id     = "LwJ075ut2LSyk1Hugq5auNs5fRpdhwuD"
  client_secret = "SH3QggWpWZfoO4qpuh-bUTdm2EXdTkZB9fvfAAp043hPBmRf3KEYBDKU0QOmjEio"
}

variable "my_local_test_org" {
  default = "define a local test org"
}

resource "administration_processing_presets" "preset_720p_25fps" {
  name = "Up to 720p, H.264, 25fps"

  video_medias = [{
    codec = "h264"
    coder = "X264"
    resolution = {
      width  = 1280
      height = 720
    }
    bitrate   = 4500000
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

resource "administration_processing_presets" "preset_with_type" {
  name = "test_preset_standard"
  type = "standard"

  video_medias = [{
    codec = "h264"
    coder = "X264"
    resolution = {
      width  = 1280
      height = 720
    }
    bitrate   = 4500000
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

resource "administration_processing_presets" "preset_with_org_and_type" {
  org  = var.my_local_test_org
  name = "test_preset_standard_with_org"
  type = "standard"

  video_medias = [{
    codec = "h264"
    coder = "X264"
    resolution = {
      width  = 1280
      height = 720
    }
    bitrate   = 4500000
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
