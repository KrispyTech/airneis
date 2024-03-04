/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{jsx,js}"],
  theme: {
    extend: {
      colors: {
        primary: "#FFEDE6",
        secondary: "#382E2A",
        tertiary: "#DBC1B8",
        clay: "#C67955",
        gray: "#E2E2E2",
        listOdd: "#FFF8F6",
        listEven: "#F4DED7",
        modal: "#F5F0EF"
      },
      lineHeight: {
        "extra-loose": "2.5"
      },
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic": "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))"
      },
      dropShadow: {
        textOutline: "0 0 1px rgba(0,0,0,1)",
        box: "2px 2px 2px rgba(0,0,0,0.8)"
      }
    }
  },
  plugins: []
}
