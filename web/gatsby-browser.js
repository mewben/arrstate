import React from "react"

import { ReactQueryProvider, StylesProvider, AuthProvider } from "@Providers"
import { Loading } from "@Components/generic"
import "@Providers/i18n"

export const wrapRootElement = ({ element }) => {
  return (
    <React.Suspense fallback={<Loading />}>
      <ReactQueryProvider>
        <StylesProvider>
          <AuthProvider>{element}</AuthProvider>
        </StylesProvider>
      </ReactQueryProvider>
    </React.Suspense>
  )
}
