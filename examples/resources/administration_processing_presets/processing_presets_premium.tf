resource "administration_processing_presets" "preset_2160p_hevc_50fps" {
  org  = "my_org"
  type = "premium"
  name = "Up to 2160p, HEVC, 50fps"
  video_medias = [
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 3840,
        height = 2160
      }
      bitrate   = 18000000,
      framerate = "50"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 1920,
        height = 1080
      }
      bitrate   = 5000000,
      framerate = "50"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 1280,
        height = 720
      }
      bitrate   = 3800000,
      framerate = "50"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 1280,
        height = 720
      }
      bitrate   = 2000000,
      framerate = "25"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 960,
        height = 540
      }
      bitrate   = 1500000,
      framerate = "25"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 768,
        height = 432
      }
      bitrate   = 1200000,
      framerate = "25"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 640,
        height = 360
      }
      bitrate   = 900000,
      framerate = "25"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 512,
        height = 288
      }
      bitrate   = 600000,
      framerate = "25"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 384,
        height = 216
      }
      bitrate   = 400000,
      framerate = "25"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 320,
        height = 180
      }
      bitrate   = 250000,
      framerate = "25"
    }
  ]

  audio_medias = [
    {
      codec      = "aac-lc",
      bitrate    = 96000,
      samplerate = "48000",
      channels   = "2.0",
      track      = "eng",
      output     = "eng"
    }
  ]
  subtitle_medias = [
    {
      track = "eng"
    }
  ]
}

resource "administration_processing_presets" "preset_1080p_hevc_50fps" {
  org  = "my_org"
  type = "premium"
  name = "Up to 1080p, HEVC, 50fps"
  video_medias = [
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 1920,
        height = 1080
      }
      bitrate   = 5000000,
      framerate = "50"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 1280,
        height = 720
      }
      bitrate   = 3800000,
      framerate = "50"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 1280,
        height = 720
      }
      bitrate   = 2000000,
      framerate = "25"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 960,
        height = 540
      }
      bitrate   = 1500000,
      framerate = "25"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 768,
        height = 432
      }
      bitrate   = 1200000,
      framerate = "25"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 640,
        height = 360
      }
      bitrate   = 900000,
      framerate = "25"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 512,
        height = 288
      }
      bitrate   = 600000,
      framerate = "25"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 384,
        height = 216
      }
      bitrate   = 400000,
      framerate = "25"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 320,
        height = 180
      }
      bitrate   = 250000,
      framerate = "25"
    }
  ]
  audio_medias = [
    {
      codec      = "aac-lc",
      bitrate    = 96000,
      samplerate = "48000",
      channels   = "2.0",
      track      = "eng",
      output     = "eng"
    }
  ]
  subtitle_medias = [
    {
      track = "eng"
    }
  ]
}
resource "administration_processing_presets" "preset_540p_hevc_25fps" {
  org  = "my_org"
  type = "premium"
  name = "Up to 540p, HEVC, 25fps"
  video_medias = [
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 960,
        height = 540
      }
      bitrate   = 1500000,
      framerate = "25"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 768,
        height = 432
      }
      bitrate   = 1200000,
      framerate = "25"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 640,
        height = 360
      }
      bitrate   = 900000,
      framerate = "25"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 512,
        height = 288
      }
      bitrate   = 600000,
      framerate = "25"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 384,
        height = 216
      }
      bitrate   = 400000,
      framerate = "25"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 320,
        height = 180
      }
      bitrate   = 250000,
      framerate = "25"
    }
  ]
  audio_medias = [
    {
      codec      = "aac-lc",
      bitrate    = 96000,
      samplerate = "48000",
      channels   = "2.0",
      track      = "eng",
      output     = "eng"
    }
  ]
  subtitle_medias = [
    {
      track = "eng"
    }
  ]
}

resource "administration_processing_presets" "preset_1080p_h264_50fps" {
  org  = "my_org"
  type = "premium"
  name = "Up to 1080p, H.264, 50fps"
  video_medias = [
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 1920,
        height = 1080
      }
      bitrate   = 6000000,
      framerate = "50"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 1280,
        height = 720
      }
      bitrate   = 5500000,
      framerate = "50"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 1280,
        height = 720
      }
      bitrate   = 3000000,
      framerate = "25"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 960,
        height = 540
      }
      bitrate   = 1800000,
      framerate = "25"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 768,
        height = 432
      }
      bitrate   = 1600000,
      framerate = "25"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 640,
        height = 360
      }
      bitrate   = 1200000,
      framerate = "25"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 512,
        height = 288
      }
      bitrate   = 800000,
      framerate = "25"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 384,
        height = 216
      }
      bitrate   = 500000,
      framerate = "25"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 320,
        height = 180
      }
      bitrate   = 300000,
      framerate = "25"
    }
  ]
  audio_medias = [
    {
      codec      = "aac-lc",
      bitrate    = 96000,
      samplerate = "48000",
      channels   = "2.0",
      track      = "eng",
      output     = "eng"
    }
  ]
  subtitle_medias = [
    {
      track = "eng"
    }
  ]
}
resource "administration_processing_presets" "preset_540p_h264_25fps" {
  org  = "my_org"
  type = "premium"
  name = "Up to 540p, H.264, 25fps"
  video_medias = [
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 960,
        height = 540
      }
      bitrate   = 1800000,
      framerate = "25"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 768,
        height = 432
      }
      bitrate   = 1600000,
      framerate = "25"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 640,
        height = 360
      }
      bitrate   = 1200000,
      framerate = "25"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 512,
        height = 288
      }
      bitrate   = 800000,
      framerate = "25"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 384,
        height = 216
      }
      bitrate   = 500000,
      framerate = "25"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 320,
        height = 180
      }
      bitrate   = 300000,
      framerate = "25"
    }
  ]
  audio_medias = [
    {
      codec      = "aac-lc",
      bitrate    = 96000,
      samplerate = "48000",
      channels   = "2.0",
      track      = "eng",
      output     = "eng"
    }
  ]
  subtitle_medias = [
    {
      track = "eng"
    }
  ]
}
resource "administration_processing_presets" "preset_2160p_hevc_60fps" {
  org  = "my_org"
  type = "premium"
  name = "Up to 2160p, HEVC, 60fps"
  video_medias = [
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 3840,
        height = 2160
      }
      bitrate   = 18000000,
      framerate = "60"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 1920,
        height = 1080
      }
      bitrate   = 5000000,
      framerate = "60"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 1280,
        height = 720
      }
      bitrate   = 3800000,
      framerate = "60"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 1280,
        height = 720
      }
      bitrate   = 2000000,
      framerate = "30"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 960,
        height = 540
      }
      bitrate   = 1500000,
      framerate = "30"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 768,
        height = 432
      }
      bitrate   = 1200000,
      framerate = "30"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 640,
        height = 360
      }
      bitrate   = 900000,
      framerate = "30"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 512,
        height = 288
      }
      bitrate   = 600000,
      framerate = "30"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 384,
        height = 216
      }
      bitrate   = 400000,
      framerate = "30"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 320,
        height = 180
      }
      bitrate   = 250000,
      framerate = "30"
    }
  ]
  audio_medias = [
    {
      codec      = "aac-lc",
      bitrate    = 96000,
      samplerate = "48000",
      channels   = "2.0",
      track      = "eng",
      output     = "eng"
    }
  ]
  subtitle_medias = [
    {
      track = "eng"
    }
  ]
}
resource "administration_processing_presets" "preset_1080p_hevc_60fps" {
  org  = "my_org"
  type = "premium"
  name = "Up to 1080p, HEVC, 60fps"
  video_medias = [
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 1920,
        height = 1080
      }
      bitrate   = 5000000,
      framerate = "60"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 1280,
        height = 720
      }
      bitrate   = 3800000,
      framerate = "60"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 1280,
        height = 720
      }
      bitrate   = 2000000,
      framerate = "30"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 960,
        height = 540
      }
      bitrate   = 1500000,
      framerate = "30"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 768,
        height = 432
      }
      bitrate   = 1200000,
      framerate = "30"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 640,
        height = 360
      }
      bitrate   = 900000,
      framerate = "30"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 512,
        height = 288
      }
      bitrate   = 600000,
      framerate = "30"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 384,
        height = 216
      }
      bitrate   = 400000,
      framerate = "30"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 320,
        height = 180
      }
      bitrate   = 250000,
      framerate = "30"
    }
  ]
  audio_medias = [
    {
      codec      = "aac-lc",
      bitrate    = 96000,
      samplerate = "48000",
      channels   = "2.0",
      track      = "eng",
      output     = "eng"
    }
  ]
  subtitle_medias = [
    {
      track = "eng"
    }
  ]
}
resource "administration_processing_presets" "preset_540p_hevc_30fps" {
  org  = "my_org"
  type = "premium"
  name = "Up to 540p, HEVC, 30fps"
  video_medias = [
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 960,
        height = 540
      }
      bitrate   = 1500000,
      framerate = "30"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 768,
        height = 432
      }
      bitrate   = 1200000,
      framerate = "30"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 640,
        height = 360
      }
      bitrate   = 900000,
      framerate = "30"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 512,
        height = 288
      }
      bitrate   = 600000,
      framerate = "30"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 384,
        height = 216
      }
      bitrate   = 400000,
      framerate = "30"
    },
    {
      codec = "hevc",
      coder = "X265",
      resolution = {
        width  = 320,
        height = 180
      }
      bitrate   = 250000,
      framerate = "30"
    }
  ]
  audio_medias = [
    {
      codec      = "aac-lc",
      bitrate    = 96000,
      samplerate = "48000",
      channels   = "2.0",
      track      = "eng",
      output     = "eng"
    }
  ]
  subtitle_medias = [
    {
      track = "eng"
    }
  ]
}
resource "administration_processing_presets" "preset_1080p_h264_60fps" {
  org  = "my_org"
  type = "premium"
  name = "Up to 1080p, H.264, 60fps"
  video_medias = [
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 1920,
        height = 1080
      }
      bitrate   = 6000000,
      framerate = "60"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 1280,
        height = 720
      }
      bitrate   = 5500000,
      framerate = "60"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 1280,
        height = 720
      }
      bitrate   = 3000000,
      framerate = "30"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 960,
        height = 540
      }
      bitrate   = 1800000,
      framerate = "30"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 768,
        height = 432
      }
      bitrate   = 1600000,
      framerate = "30"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 640,
        height = 360
      }
      bitrate   = 1200000,
      framerate = "30"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 512,
        height = 288
      }
      bitrate   = 800000,
      framerate = "30"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 384,
        height = 216
      }
      bitrate   = 500000,
      framerate = "30"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 320,
        height = 180
      }
      bitrate   = 300000,
      framerate = "30"
    }
  ]
  audio_medias = [
    {
      codec      = "aac-lc",
      bitrate    = 96000,
      samplerate = "48000",
      channels   = "2.0",
      track      = "eng",
      output     = "eng"
    }
  ]
  subtitle_medias = [
    {
      track = "eng"
    }
  ]
}
resource "administration_processing_presets" "preset_540p_h264_30fps" {
  org  = "my_org"
  type = "premium"
  name = "Up to 540p, H.264, 30fps"
  video_medias = [
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 960,
        height = 540
      }
      bitrate   = 1800000,
      framerate = "30"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 768,
        height = 432
      }
      bitrate   = 1600000,
      framerate = "30"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 640,
        height = 360
      }
      bitrate   = 1200000,
      framerate = "30"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 512,
        height = 288
      }
      bitrate   = 800000,
      framerate = "30"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 384,
        height = 216
      }
      bitrate   = 500000,
      framerate = "30"
    },
    {
      codec = "h264",
      coder = "X264",
      resolution = {
        width  = 320,
        height = 180
      }
      bitrate   = 300000,
      framerate = "30"
    }
  ]
  audio_medias = [
    {
      codec      = "aac-lc",
      bitrate    = 96000,
      samplerate = "48000",
      channels   = "2.0",
      track      = "eng",
      output     = "eng"
    }
  ]
  subtitle_medias = [
    {
      track = "eng"
    }
  ]
}
