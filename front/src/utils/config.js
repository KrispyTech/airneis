const config = {
  api: {
    url: typeof window === "undefined" ? process.env.API_URL : process.env.NEXT_PUBLIC_API_URL
    // I did this to not break Barnabe's page of test
  }
}

export default config
