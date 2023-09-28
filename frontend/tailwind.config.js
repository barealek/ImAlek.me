/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {},
  },
  plugins: [
    require('tailwindcss-animated'),
  ],
  safelist: [
    "animate-delay-[150ms]",
    "animate-delay-[300ms]",
    "animate-delay-[450ms]",
    "animate-delay-[600ms]",
    "animate-delay-[750ms]",
    "animate-delay-[900ms]",
    "animate-delay-[1050ms]",
    "animate-delay-[1200ms]",
  ]
}

