import React from "react"
import { ReactQueryConfigProvider } from "react-query"

const queryConfig = { queries: { refetchOnWindowFocus: false } }
export const ReactQueryProvider = ({ children }) => {
  return (
    <ReactQueryConfigProvider config={queryConfig}>
      {children}
    </ReactQueryConfigProvider>
  )
}
