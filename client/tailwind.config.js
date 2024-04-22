/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./web/components/**/*.{templ,go}",
    "./web/routes/**/*.{templ,go}",
    "./handlers/**/*.go",
  ],
  daisyui: {
    themes: [
      {
        "dark-theme": {
          primary: "#CFCA74", //primary
          "primary-content": "#343200", //onPrimary
          primaryContainer: "#4B4900", //primaryContainer
          onPrimaryContainer: "#EBE68D", //onPrimaryContaienr
          secondary: "#CCC8A4", //secondary
          "secondary-content": "#333118", //onSecondary
          secondaryContainer: "#4A482C", //secondaryContainer
          onSecondaryContainer: "#E8E4BE", //onSecondaryContainer
          accent: "#A5D0BB", //tertiary
          "accent-content": "#0D3728", //onTertiary
          accentContainer: "#264E3E", //tertiaryContainer
          onAccentContainer: "#C0ECD6", //onTertiaryContainer
          error: "#FFB4AB", //error
          "error-content": "#690005", // onError
          errorContainer: "#93000A", //errorContainer
          onErrorContainer: "#FFDAD6", //onErrorContainer
          "base-100": "#14140C", //background
          "base-content": "#E6E2D5", //onBackground
          surface: "#14140C", //surface
          onSurface: "#E6E2D5", //onSurface
          surfaceVariant: "#49473A", //surfaceVariant
          onSurfaceVariant: "#CAC7B5", //onSurfaceVariant
          outline: "#949181", //outline
          outlineVariant: "#49473A", //outlineVariant
          shadow: "#000000", //shadow
          scrim: "#000000", //scrim
          inverseSurface: "#E6E2D5", //inverseSurface
          inverseOnSurface: "#323128", //inverseOnSurface
          inversePrimary: "#646116", //inversePrimary
          primaryFixed: "#EBE68D", //primaryFixed
          onPrimaryFixed: "#1E1C00", //onPrimaryFixed
          primaryFixedDim: "#CFCA74", //primaryFixedDim
          onPrimaryFixedVariant: "#4B4900", //onPrimaryFixedDim
          secondaryFixed: "#E8E4BE", //secondaryFixed
          onSecondaryFixed: "#1D1C05", //onSecondaryFixed
          secondaryFixedDim: "#CCC8A4", //secondaryFixedDim
          onSecondaryFixedVariant: "#4A482C", //onSecondaryFixedVariant
          tertiaryFixed: "#C0ECD6", //tertiaryFixed
          onTertiaryFixed: "#002115", //onTertiaryFixed
          tertiaryFixedDim: "#A5D0BB", //tertiaryFixedDim
          onTertiaryFixedVariant: "#264E3E", //onTertiaryFixedVariant
          surfaceDim: "#14140C", //surfaceDim
          surfaceBright: "#3B3930", //surfaceBright
          surfaceContainerLowest: "#0F0E07", //surfaceContainerLowest
          surfaceContainerLow: "#1D1C14", //surfaceContainerLow
          surfaceContainer: "#212018", //surfaceContainer
          surfaceContainerHigh: "#2B2A22", //surfaceContainerHigh
          surfaceContainerHighest: "#36352C", //surfaceContainerHighest
        },
      },
    ],
  },
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
};
