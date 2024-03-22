/** @type {import('next').NextConfig} */

// We have to use the require syntax to import the i18n object from the next-i18next.config.js file.
const { i18n } = require("./next-i18next.config")
const nextConfig = {
  reactStrictMode: true,
  eslint: {
    ignoreDuringBuilds: true
  },
  i18n
}

module.exports = nextConfig
