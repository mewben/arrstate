import React from "react"
import { SWRConfig } from "swr"

import { privateFetcher } from "@Utils/request"

const config = {
  errorRetryCount: 1,
  fetcher: privateFetcher,
}

export const SwrProvider = ({ children }) => {
  return <SWRConfig value={config}>{children}</SWRConfig>
}
