import React from "react"

import { ReduxProvider, AuthProvider } from "@Providers"
import { LayoutWrapper } from "@Wrappers"

export const wrapRootElement = ({ element }) => {
  return (
    <ReduxProvider>
      <AuthProvider>
        <LayoutWrapper>{element}</LayoutWrapper>
      </AuthProvider>
    </ReduxProvider>
  )
}
