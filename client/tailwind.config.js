/** @type {import('tailwindcss').Config} */
export default {
  content: ["./web/components/**/*.{templ,go}", "./web/routes/**/*.{templ,go}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
};
