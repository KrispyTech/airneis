const mobileConfig = {
  api: {
    url: process.env.EXPO_PUBLIC_API_URL
  },
  images: {
    fullScreenImageXL: {
      width: "100%",
      height: 300
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
