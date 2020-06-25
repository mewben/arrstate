import React from "react"

import { ReduxProvider, AuthProvider } from "@Providers"

export const wrapRootElement = ({ element }) => {
  return (
    <ReduxProvider>
      <AuthProvider>{element}</AuthProvider>
    </ReduxProvider>
  )
}
