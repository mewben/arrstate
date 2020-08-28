import React from "react"

import {
  ReactQueryProvider,
  StylesProvider,
  AuthProvider,
  DatePickerProvider,
} from "@Providers"
import { Loading } from "@Components/generic"
import "@Providers/i18n"
// import { ReactQueryDevtools } from "react-query-devtools"

export const wrapRootElement = ({ element }) => {
  return (
    <React.Suspense fallback={<Loading />}>
      <ReactQueryProvider>
        <StylesProvider>
          <DatePickerProvider>
            <AuthProvider>{element}</AuthProvider>
          </DatePickerProvider>
        </StylesProvider>
        {/* <ReactQueryDevtools /> */}
      </ReactQueryProvider>
    </React.Suspense>
  )
}
