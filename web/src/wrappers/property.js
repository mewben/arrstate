import React from "react"

import { useProperty } from "@Hooks"
import { Loading, Error } from "@Components/generic"

export const PropertyWrapper = ({ propertyID, children }) => {
  const { status, data, error } = useProperty(propertyID)

  return status === "loading" ? (
    <Loading />
  ) : status === "error" || !data ? (
    <Error error={error} />
  ) : (
    children({ property: data })
  )
}
