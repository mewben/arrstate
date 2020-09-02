import React from "react"
import { createMuiTheme, ThemeProvider } from "@material-ui/core/styles"
import CssBaseline from "@material-ui/core/CssBaseline"
import Button from "@material-ui/core/Button"
import defaultTheme from "tailwindcss/defaultTheme"
import colors from "@tailwindcss/ui/colors"
import "typeface-inter"
import "./main.css"

const theme = createMuiTheme({
  palette: {
    divider: colors["cool-gray"][200], // "#e5e7eb",
  },
  typography: {
    fontFamily: ["Inter var", ...defaultTheme.fontFamily.sans].join(","),
  },
})

theme.overrides = {
  MuiFormControl: {
    root: {
      border: "inherit",
      borderWidth: 1,
    },
  },
  MuiInputBase: {
    input: {
      height: "inherit",
    },
  },
  MuiAutocomplete: {
    root: {
      "&$focused .form-input": {
        boxShadow: `0 0 0 3px rgba(164, 202, 254, 0.45)`,
        outline: "none",
        borderColor: colors.blue[300],
      },
    },
    endAdornment: {
      right: 8,
      "& .MuiIconButton-root": {
        color: colors.gray[400],
      },
    },
    popupIndicator: {
      padding: 4,
    },
    inputRoot: {
      '&[class*="MuiInput-root"]': {
        paddingBottom: 0,
        padding: 0,
        "& $input": {
          padding: 0,
        },
      },
      '&[class*="MuiInput-root"][class*="MuiInput-marginDense"]': {
        "& $input": {
          paddingTop: "0.5rem",
          paddingRight: "0.75rem",
          paddingBottom: "0.5rem",
          paddingLeft: "0.75rem",
        },
        "& $input:first-child": {
          paddingTop: "0.5rem",
          paddingRight: "0.75rem",
          paddingBottom: "0.5rem",
          paddingLeft: "0.75rem",
        },
      },
    },
    input: {
      fontSize: "0.875rem",
      lineHeight: "1.25rem",
    },
  },
}

export const StylesProvider = ({ children }) => {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      {children}
    </ThemeProvider>
  )
}
