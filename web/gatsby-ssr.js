import React from "react"

import {
  ReactQueryProvider,
  StylesProvider,
  AuthProvider,
  DateProvider,
} from "@Providers"

export const wrapRootElement = ({ element }) => {
  return (
    <ReactQueryProvider>
      <StylesProvider>
        <DateProvider>
          <AuthProvider>{element}</AuthProvider>
        </DateProvider>
      </StylesProvider>
    </ReactQueryProvider>
  )
}
