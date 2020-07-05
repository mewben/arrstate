import React from "react"
import { createMuiTheme, ThemeProvider } from "@material-ui/core/styles"
import CssBaseline from "@material-ui/core/CssBaseline"
import Button from "@material-ui/core/Button"
import defaultTheme from "tailwindcss/defaultTheme"
import "typeface-inter"
import "./main.css"

const theme = createMuiTheme({
  palette: {
    divider: "#e5e7eb",
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
  MuiAutocomplete: {
    inputRoot: {
      '&[class*="MuiInput-root"]': {
        paddingBottom: "0.125rem",
        paddingTop: "0.125rem",
        paddingLeft: "0.125rem",
        "& $input": {
          paddingTop: "0.5rem",
          paddingBottom: "0.5rem",
          paddingLeft: "0.75rem",
          paddingRight: "0.75rem",
        },
        "& $input:first-child": {
          paddingTop: "0.5rem",
          paddingBottom: "0.5rem",
          paddingLeft: "0.75rem",
          paddingRight: "0.75rem",
        },
      },
      '&[class*="MuiInput-root"][class*="MuiInput-marginDense"]': {
        "& $input": {
          paddingTop: "0.5rem",
          paddingBottom: "0.5rem",
          paddingLeft: "0.75rem",
          paddingRight: "0.75rem",
        },
        "& $input:first-child": {
          paddingTop: "0.5rem",
          paddingBottom: "0.5rem",
          paddingLeft: "0.75rem",
          paddingRight: "0.75rem",
        },
      },
    },
    input: {
      fontSize: "0.875rem",
      lineHeight: "1.25rem",
    },
    endAdornment: {
      right: 32,
      "& .MuiIconButton-root": {
        color: "#9fa6b2", // gray-400
      },
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
