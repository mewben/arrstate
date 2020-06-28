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

export const StylesProvider = ({ children }) => {
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      {children}
    </ThemeProvider>
  )
}
