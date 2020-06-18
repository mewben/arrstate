import React from "react"

import { AuthProvider } from "@Providers"
import { LayoutWrapper } from "@Wrappers"

export const wrapRootElement = ({ element }) => {
  return (
    <AuthProvider>
      <LayoutWrapper>{element}</LayoutWrapper>
    </AuthProvider>
  )
}
