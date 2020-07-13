import React from "react"
import { useQuery } from "react-query"
import { requestApi } from "@Utils"
import { map, sortBy, keyBy } from "@Utils/lodash"

const fetchProjects = async () => {
  const { data } = await requestApi("/api/projects")
  return data
}

const fetchProject = async (_, projectID) => {
  if (!projectID) return null
  const { data } = await requestApi(`/api/projects/${projectID}`)
  return data
}

export const useProjects = () => {
  return useQuery(["projects"], fetchProjects)
}

export const useProject = projectID => {
  return useQuery(["project", projectID], fetchProject)
}

// export const useProjectsMap = () => {
//   const { status, data, error, isFetching } = useProjects()

//   const projectsMap = React.useMemo(() => {
//     return keyBy(data?.list, "_id")
//   }, [data?.list])

//   return {
//     status: isFetching ? "loading" : status,
//     projectsMap,
//     error,
//   }
// }

// returns projectOptions and finds the selected option by projectID
export const useProjectOptions = () => {
  const { status, data, error, isFetching } = useProjects()

  const options = React.useMemo(() => {
    const options = map(data?.list, item => {
      return {
        value: item._id,
        label: item.name,
      }
    })
    return sortBy(options, "label")
  }, [data?.list])

  return {
    status: isFetching ? "loading" : status,
    options,
    error,
  }
}
