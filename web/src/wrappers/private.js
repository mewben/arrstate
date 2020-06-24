import React from "react"
import { Redirect } from "@reach/router"

import { useAuth } from "@Providers"

export const PrivateWrapper = ({ children }) => {
  const { isLoading, isAuthenticated } = useAuth()
  if (isLoading) {
    return null
  }

  if (!isAuthenticated) {
    return <Redirect to="/signin" noThrow />
  }

  return children
}

const MeWrapper = ({ children }) => {
  // fetch me
}
