import { useQuery } from "react-query"
import { requestApi } from "@Utils"

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
