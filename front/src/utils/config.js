const config = {
  api: {
    url: typeof window === "undefined" ? process.env.API_URL : process.env.NEXT_PUBLIC_API_URL
    // I did this to not break Barnabe's page of test
  },
  images: {
    productPage: {
      carousel: {
        width: 550,
        height: 400
      }
    },
    productCard: {
      imageWidth: 350,
      imageHeight: 200
    },
    categoryCard: {
      imageWidth: 350,
      imageHeight: 200
    },
    backoffice: {
      carousel: {
        width: 350,
        height: 200
      }
    }
  }
}

export default config
