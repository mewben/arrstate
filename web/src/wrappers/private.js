import React from "react"
import { Redirect } from "@reach/router"

import { Loading, Error } from "@Components/generic"
import { useAuth } from "@Providers"
import { useMe } from "@Hooks"
import { LayoutWrapper } from "@Wrappers"

const MeContext = React.createContext()
export const useMeContext = () => React.useContext(MeContext)

export const PrivateWrapper = ({ children }) => {
  const { isLoading, isAuthenticated } = useAuth()
  if (isLoading) {
    return null
  }

  if (!isAuthenticated) {
    return <Redirect to="/signin" noThrow />
  }

  return (
    <MeWrapper>
      <LayoutWrapper>{children}</LayoutWrapper>
    </MeWrapper>
  )
}

const MeWrapper = ({ children }) => {
  const { status, data, error } = useMe()
  // TODO: redirect to /welcome if not yet onboarded

  return status === "loading" ? (
    <Loading />
  ) : status === "error" ? (
    <Error error={error} />
  ) : (
    <MeContext.Provider value={data}>{children}</MeContext.Provider>
  )
}
