import React from "react"

import {
  ReactQueryProvider,
  SwrProvider,
  StylesProvider,
  AuthProvider,
  DateProvider,
} from "@Providers"
import { Loading } from "@Components/generic"
import "@Providers/i18n"
// import { ReactQueryDevtools } from "react-query-devtools"

export const wrapRootElement = ({ element }) => {
  return (
    <React.Suspense fallback={<Loading />}>
      <SwrProvider>
        <ReactQueryProvider>
          <StylesProvider>
            <DateProvider>
              <AuthProvider>{element}</AuthProvider>
            </DateProvider>
          </StylesProvider>
        </ReactQueryProvider>
      </SwrProvider>
    </React.Suspense>
  )
}

/*
export const wrapRootElement2 = ({ element }) => {
  return (
    <React.Suspense fallback={<Loading />}>
      <ReactQueryProvider>
        <StylesProvider>
          <DateProvider>
            <AuthProvider>{element}</AuthProvider>
          </DateProvider>
        </StylesProvider>
      </ReactQueryProvider>
    </React.Suspense>
  )
}
*/
