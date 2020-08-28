import React from "react"

import {
  ReactQueryProvider,
  StylesProvider,
  AuthProvider,
  DatePickerProvider,
} from "@Providers"

export const wrapRootElement = ({ element }) => {
  return (
    <ReactQueryProvider>
      <StylesProvider>
        <DatePickerProvider>
          <AuthProvider>{element}</AuthProvider>
        </DatePickerProvider>
      </StylesProvider>
    </ReactQueryProvider>
  )
}
