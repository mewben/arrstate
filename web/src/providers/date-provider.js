import React from "react"
import { LocalizationProvider } from "@material-ui/pickers"
import LuxonAdapter from "@material-ui/pickers/adapter/luxon"

export const DateProvider = ({ children }) => {
  return (
    <LocalizationProvider dateAdapter={LuxonAdapter}>
      {children}
    </LocalizationProvider>
  )
}
