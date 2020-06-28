import React from "react"

import { ReduxProvider, StylesProvider, AuthProvider } from "@Providers"

export const wrapRootElement = ({ element }) => {
  return (
    <ReduxProvider>
      <StylesProvider>
        <AuthProvider>{element}</AuthProvider>
      </StylesProvider>
    </ReduxProvider>
  )
}
