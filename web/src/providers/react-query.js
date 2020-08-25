import React from "react"
import { ReactQueryConfigProvider } from "react-query"

import { requestApi } from "@Utils"

const defaultQueryFn = async (key, { path, params }) => {
  const { data } = await requestApi(path, "GET", {
    params,
  })
  return data
}

const queryConfig = {
  queries: { queryFn: defaultQueryFn, refetchOnWindowFocus: false, retry: 0 },
}
export const ReactQueryProvider = ({ children }) => {
  return (
    <ReactQueryConfigProvider config={queryConfig}>
      {children}
    </ReactQueryConfigProvider>
  )
}
