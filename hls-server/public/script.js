document.addEventListener("DOMContentLoaded", function () {
  const video = document.getElementById("video");

  if (Hls.isSupported()) {
    const hls = new Hls();

    hls.loadSource("http://localhost:8080/hls/playlist.m3u8");
    hls.attachMedia(video);

    hls.on(Hls.Events.MANIFEST_PARSED, function () {
      video.play().catch((error) => {
        console.error("Error trying to play video:", error);
      });
    });

    hls.on(Hls.Events.ERROR, function (event, data) {
      if (data.fatal) {
        switch (data.fatal) {
          case Hls.ErrorTypes.NETWORK_ERROR:
            console.error("Network error:", data);
            break;
          case Hls.ErrorTypes.MEDIA_ERROR:
            console.error("Media error:", data);
            break;
          case Hls.ErrorTypes.OTHER_ERROR:
            console.error("Other error:", data);
            break;
          default:
            console.error("Unrecognized error:", data);
            break;
        }
      }
    });
  } else {
    console.error("HLS.js is not supported in this browser.");
    video.src = "http://localhost:8080/hls/playlist.m3u8"; // Fallback for browsers that support native HLS
  }
});
