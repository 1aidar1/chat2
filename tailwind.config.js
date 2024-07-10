/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.{html,js,gohtml}"],
  theme: {
    extend: {
      scrollbar: {
        hide: 'scrollbar-width: none; -ms-overflow-style: none; &::-webkit-scrollbar { display: none; }',
      },
    },
  },
  plugins: [
    function ({ addUtilities }) {
      const newUtilities = {
        '.scrollbar-hide': {
          'scrollbar-width': 'none', /* Firefox */
          '-ms-overflow-style': 'none',  /* IE и Edge */
          '&::-webkit-scrollbar': {
            display: 'none', /* Chrome, Safari и Opera */
          },
        },
      }

      addUtilities(newUtilities)
    }
  ],}

