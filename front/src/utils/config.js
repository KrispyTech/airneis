const config = {
  api: {
    url: typeof window === "undefined" ? process.env.API_URL : process.env.NEXT_PUBLIC_API_URL
  }
}

export default config
