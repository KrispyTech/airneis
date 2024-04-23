const mobileConfig = {
  api: {
    url: process.env.EXPO_PUBLIC_API_URL
  },
  dropShadow: { textOutline: { shadowOffset: { width: "0px", height: "0px" }, shadowOpacity: "100%" } },
  images: {
    fullScreenImageXL: {
      width: "100%",
      height: 300
    },
    fullScreenImageLG: {
      width: "100%",
      height: 250
    },
    fullScreenImageMD: {
      width: "100%",
      height: 200
    }
  },
  icons: {
    lgIcon: {
      size: 30,
      strokeWidth: 2
    },
    xlIcon: {
      size: 33,
      strokeWidth: 2
    }
  }
}

export default mobileConfig
