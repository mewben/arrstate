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
  const { status, data, error } = useProjects()

  const { options, selectedOption } = React.useMemo(() => {
    let selectedOption = null
    const options = map(data?.list, item => {
      const option = {
        value: item._id,
        label: item.name,
      }
      if (projectID === item._id) {
        selectedOption = option
      }
      return option
    })

    return {
      options: sortBy(options, "label"),
      selectedOption,
    }
  }, [data?.list])

  return {
    status,
    options,
    selectedOption,
    error,
  }
}
