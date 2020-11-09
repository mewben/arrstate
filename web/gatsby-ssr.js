import React from "react"

import {
  ReactQueryProvider,
  SwrProvider,
  StylesProvider,
  AuthProvider,
  DateProvider,
} from "@Providers"

export const wrapRootElement = ({ element }) => {
  return (
    <SwrProvider>
      <ReactQueryProvider>
        <StylesProvider>
          <DateProvider>
            <AuthProvider>{element}</AuthProvider>
          </DateProvider>
        </StylesProvider>
      </ReactQueryProvider>
    </SwrProvider>
  )
}
