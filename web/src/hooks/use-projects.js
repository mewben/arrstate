import React from "react"
import { useQuery } from "react-query"
import { requestApi } from "@Utils"
import { map, sortBy } from "@Utils/lodash"

const fetchProjects = async () => {
  const { data } = await requestApi("/api/projects")
  return data
}

const fetchProject = async (_, projectID) => {
  const { data } = await requestApi(`/api/projects/${projectID}`)
  return data
}

export const useProjects = () => {
  return useQuery(["projects"], fetchProjects)
}

export const useProject = projectID => {
  return useQuery(["project", projectID], fetchProject)
}

// returns projectOptions and finds the selected option by projectID
export const useProjectOptions = projectID => {
  const { status, data, error, isFetching } = useProjects()

  const options = React.useMemo(() => {
    const options = map(data?.list, item => {
      return {
        value: item._id,
        label: item.name,
      }
    })
    return sortBy(options, "label")
  }, [projectID, data?.list])

  return {
    status: isFetching ? "loading" : status,
    options,
    error,
  }
}
