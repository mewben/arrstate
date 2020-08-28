import React from "react"
import { LocalizationProvider } from "@material-ui/pickers"
import DayjsUtils from "@material-ui/pickers/adapter/dayjs"

export const DatePickerProvider = ({ children }) => {
  return (
    <LocalizationProvider dateAdapter={DayjsUtils}>
      {children}
    </LocalizationProvider>
  )
}
