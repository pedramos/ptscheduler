/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["*/*.templ}", "*.templ"],
  theme: {
    extend: {},
  },
    plugins: [
        require('flowbite/plugin')
    ]
}

