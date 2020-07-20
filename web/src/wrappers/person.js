import React from "react"

import { usePerson } from "@Hooks"
import { Loading, Error } from "@Components/generic"

export const PersonWrapper = ({ personID, children }) => {
  const { status, data, error } = usePerson(personID)

  return status === "loading" ? (
    <Loading />
  ) : status === "error" ? (
    <Error error={error} />
  ) : (
    children({ person: data })
  )
}
