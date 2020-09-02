import React from "react"
import { Redirect } from "@reach/router"
import { useQuery } from "react-query"

import { Loading, Error } from "@Components/generic"
import { useAuth } from "@Providers"
import { LayoutWrapper } from "@Wrappers"

const CurrentContext = React.createContext()
export const useCurrentContext = () => React.useContext(CurrentContext)

export const PrivateWrapper = ({ children }) => {
  const { isLoading, isAuthenticated } = useAuth()
  if (isLoading) {
    return null
  }

  if (!isAuthenticated) {
    return <Redirect to="/signin" noThrow />
  }

  return (
    <CurrentWrapper>
      <LayoutWrapper>{children}</LayoutWrapper>
    </CurrentWrapper>
  )
}

const CurrentWrapper = ({ children }) => {
  // /api/users/current
  // /api/businesses/current
  // /api/people/current
  const { status: status1, data: data1, error: error1 } = useQuery(
    ["currentUser", { path: "/api/users/current" }],
    null,
    { staleTime: Infinity }
  )

  const { status: status2, data: data2, error: error2 } = useQuery(
    ["currentBusiness", { path: "/api/businesses/current" }],
    null,
    { staleTime: Infinity }
  )

  const { status: status3, data: data3, error: error3 } = useQuery(
    ["currentPerson", { path: "/api/people/current" }],
    null,
    { staleTime: Infinity }
  )

  if (status1 === "loading" || status2 === "loading" || status3 === "loading") {
    return <Loading />
  }

  if (status1 === "error" || status2 === "error" || status3 === "error") {
    return <Error error={error1 || error2 || error3} />
  }

  return (
    <CurrentContext.Provider
      value={{
        currentUser: data1,
        currentBusiness: data2,
        currentPerson: data3,
      }}
    >
      {children}
    </CurrentContext.Provider>
  )
}
