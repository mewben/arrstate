import React from "react"

import { ReactQueryProvider, StylesProvider, AuthProvider } from "@Providers"

export const wrapRootElement = ({ element }) => {
  return (
    <ReactQueryProvider>
      <StylesProvider>
        <AuthProvider>{element}</AuthProvider>
      </StylesProvider>
    </ReactQueryProvider>
  )
}
