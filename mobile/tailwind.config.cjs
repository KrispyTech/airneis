module.exports = {
  content: ["./App.{js,jsx,ts,tsx}", "./src/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      colors: {
        primary: "#FFEDE6",
        secondary: "#382E2B",
        lightGray: "#E2E2E2",
        darkGray: "#707070"
      },
      dropShadow: {
        textOutline: "0 0 1px rgba(0,0,0,1)",
        box: "2px 2px 2px rgba(0,0,0,0.8)"
      },
      borderWidth: {
        6: "6px"
      }
    }
  },
  plugins: []
}
