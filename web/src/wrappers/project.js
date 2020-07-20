import React from "react"

import { useProject } from "@Hooks"
import { Loading, Error } from "@Components/generic"

export const ProjectWrapper = ({ projectID, children }) => {
  const { status, data, error } = useProject(projectID)

  return status === "loading" ? (
    <Loading />
  ) : status === "error" || !data ? (
    <Error error={error} />
  ) : (
    children({ project: data })
  )
}
