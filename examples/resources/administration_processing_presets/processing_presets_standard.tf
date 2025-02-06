resource "administration_processing_presets" "preset_1080p_25fps" {
  org  = "orga_d7Gpg42L"
  type = "standard"
  name = "Up to 1080p, H.264, 25fps"

  video_medias = [
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 1920,
        height = 1080
      },
      bitrate   = 7800000,
      framerate = "25"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 1280,
        height = 720
      },
      bitrate   = 4500000,
      framerate = "25"
      }, {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 960
        height = 540
      }
      bitrate   = 2000000
      framerate = "25"

    },
    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 768
        height = 432
      }
      bitrate   = 730000
      framerate = "25"
    },
    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 416
        height = 234
      }
      bitrate   = 145000
      framerate = "25"
    }
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

resource "administration_processing_presets" "preset_720p_25fps" {
  org  = "orga_d7Gpg42L"
  type = "standard"
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

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 960
        height = 540
      }
      bitrate   = 2000000
      framerate = "25"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 768
        height = 432
      }
      bitrate   = 730000
      framerate = "25"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 416
        height = 234
      }
      bitrate   = 145000
      framerate = "25"
    }
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

resource "administration_processing_presets" "preset_540p_25fps" {
  org  = "orga_d7Gpg42L"
  type = "standard"
  name = "Up to 540p, H.264, 25fps"

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

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 768
        height = 432
      }
      bitrate   = 730000
      framerate = "25"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 416
        height = 234
      }
      bitrate   = 145000
      framerate = "25"
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

resource "administration_processing_presets" "preset_1080p_30fps" {
  org  = "orga_d7Gpg42L"
  type = "standard"
  name = "Up to 1080p, H.264, 30fps"


  video_medias = [{
    codec = "h264"
    coder = "X264"
    resolution = {
      width  = 1920
      height = 1080
    }
    bitrate   = 9360000
    framerate = "30"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 1280
        height = 720
      }
      bitrate   = 5400000
      framerate = "30"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 960
        height = 540
      }
      bitrate   = 2400000
      framerate = "30"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 768
        height = 432
      }
      bitrate   = 876000
      framerate = "30"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 416
        height = 234
      }
      bitrate   = 174000
      framerate = "30"
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

resource "administration_processing_presets" "preset_720p_30fps" {
  org  = "orga_d7Gpg42L"
  type = "standard"
  name = "Up to 720p, H.264, 30fps"


  video_medias = [{
    codec = "h264"
    coder = "X264"
    resolution = {
      width  = 1280
      height = 720
    }
    bitrate   = 5400000
    framerate = "30"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 960
        height = 540
      }
      bitrate   = 2400000
      framerate = "30"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 768
        height = 432
      }
      bitrate   = 876000
      framerate = "30"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 416
        height = 234
      }
      bitrate   = 174000
      framerate = "30"
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

resource "administration_processing_presets" "preset_540p_30fps" {
  org  = "orga_d7Gpg42L"
  type = "standard"
  name = "Up to 540p, H.264, 30fps"


  video_medias = [{
    codec = "h264"
    coder = "X264"
    resolution = {
      width  = 960
      height = 540
    }
    bitrate   = 2400000
    framerate = "30"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 768
        height = 432
      }
      bitrate   = 876000
      framerate = "30"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 416
        height = 234
      }
      bitrate   = 174000
      framerate = "30"
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

resource "administration_processing_presets" "preset_1080p_50fps" {
  org  = "orga_d7Gpg42L"
  type = "standard"
  name = "Up to 1080p, H.264, 50fps"


  video_medias = [{
    codec = "h264"
    coder = "X264"
    resolution = {
      width  = 1920
      height = 1080
    }
    bitrate   = 7800000
    framerate = "50"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 1280
        height = 720
      }
      bitrate   = 4500000
      framerate = "50"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 960
        height = 540
      }
      bitrate   = 2000000
      framerate = "50"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 768
        height = 432
      }
      bitrate   = 730000
      framerate = "25"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 416
        height = 234
      }
      bitrate   = 145000
      framerate = "25"
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

resource "administration_processing_presets" "preset_720p_50fps" {
  org  = "orga_d7Gpg42L"
  type = "standard"
  name = "Up to 720p, H.264, 50fps"


  video_medias = [{
    codec = "h264"
    coder = "X264"
    resolution = {
      width  = 1280
      height = 720
    }
    bitrate   = 4500000
    framerate = "50"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 960
        height = 540
      }
      bitrate   = 2000000
      framerate = "50"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 768
        height = 432
      }
      bitrate   = 730000
      framerate = "25"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 416
        height = 234
      }
      bitrate   = 145000
      framerate = "25"
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

resource "administration_processing_presets" "preset_540p_50fps" {
  org  = "orga_d7Gpg42L"
  type = "standard"
  name = "Up to 540p, H.264, 50fps"

  video_medias = [
    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 960
        height = 540
      }
      bitrate   = 2000000
      framerate = "50"
    },
    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 768
        height = 432
      }
      bitrate   = 730000
      framerate = "25"
    },
    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 416
        height = 234
      }
      bitrate   = 145000
      framerate = "25"
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

resource "administration_processing_presets" "preset_1080p_60fps" {
  org  = "orga_d7Gpg42L"
  type = "standard"
  name = "Up to 1080p, H.264, 60fps"


  video_medias = [{
    codec = "h264"
    coder = "X264"
    resolution = {
      width  = 1920
      height = 1080
    }
    bitrate   = 9360000
    framerate = "60"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 1280
        height = 720
      }
      bitrate   = 5400000
      framerate = "60"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 960
        height = 540
      }
      bitrate   = 2400000
      framerate = "60"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 768
        height = 432
      }
      bitrate   = 876000
      framerate = "30"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 416
        height = 234
      }
      bitrate   = 174000
      framerate = "30"
  }]

  audio_medias = [{
    codec      = "aac-lc"
    bitrate    = 96000
    samplerate = 48000
    channels   = "2.0"
    track      = "eng"
    output     = "eng"
  }]

  subtitle_medias = [{
    track = "eng"
  }]
}

resource "administration_processing_presets" "preset_720p_60fps" {
  org  = "orga_d7Gpg42L"
  type = "standard"
  name = "Up to 720p, H.264, 60fps"
  video_medias = [{
    codec = "h264"
    coder = "X264"
    resolution = {
      width  = 1280
      height = 720
    }
    bitrate   = 5400000
    framerate = "60"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 960
        height = 540
      }
      bitrate   = 2400000
      framerate = "60"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 768
        height = 432
      }
      bitrate   = 876000
      framerate = 30
    },
    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 416
        height = 234
      }
      bitrate   = 174000
      framerate = "30"
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


resource "administration_processing_presets" "preset_540p_60fps" {
  org  = "orga_d7Gpg42L"
  type = "standard"
  name = "Up to 540p, H.264, 60fps"

  video_medias = [{
    codec = "h264"
    coder = "X264"
    resolution = {
      width  = 960
      height = 540
    }
    bitrate   = 2400000
    framerate = "60"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 768
        height = 432
      }
      bitrate   = 876000
      framerate = "30"
    },

    {
      codec = "h264"
      coder = "X264"
      resolution = {
        width  = 416
        height = 234
      }
      bitrate   = 174000
      framerate = "30"
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

resource "administration_processing_presets" "preset_radio_128kbps" {
  org  = "orga_d7Gpg42L"
  type = "standard"
  name = "Radio, up to 128kbps"

  video_medias = []

  audio_medias = [{
    codec      = "aac-lc"
    bitrate    = 96000
    samplerate = "48000"
    channels   = "2.0"
    track      = "eng"
    output     = "eng"
    },
    {
      codec      = "aac-lc"
      bitrate    = 128000
      samplerate = "48000"
      channels   = "2.0"
      track      = "eng"
      output     = "eng"
  }]

  subtitle_medias = [{}]
}
