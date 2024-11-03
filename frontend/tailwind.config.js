/** @type {import('tailwindcss').Config} */
export default {
  content: ["./dist/index.html", "./src/**/*.{vue,html,js,ts,jsx,tsx,svelte}",],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["dark"]
  }
}

