import React from "react"
import { Redirect } from "@reach/router"

import { Loading } from "@Components/generic"
import { useAuth } from "@Providers"
import { useMe } from "@Hooks"
import { extractError } from "@Utils"

export const PrivateWrapper = ({ children }) => {
  const { isLoading, isAuthenticated } = useAuth()
  if (isLoading) {
    return null
  }

  if (!isAuthenticated) {
    return <Redirect to="/signin" noThrow />
  }

  return <MeWrapper>{children}</MeWrapper>
}

const MeWrapper = ({ children }) => {
  const { status, data, error } = useMe()
  console.log("status", status)
  console.log("data", data)
  console.log("error", error)
  // TODO: redirect to /welcome if not yet onboarded

  return status === "loading" ? (
    <Loading />
  ) : status === "error" ? (
    <div>{extractError(error)}</div>
  ) : (
    children
  )
}
