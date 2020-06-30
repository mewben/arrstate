import { useQuery } from "react-query"
import { requestApi } from "@Utils"
import { ROLES } from "@Enums"

const fetchPeople = async (_, params) => {
  const { data } = await requestApi("/api/people", "GET", { params })
  return data
}

export const useClients = () => {
  return useQuery(["people", { role: [ROLES.CLIENT] }], fetchPeople)
}

export const useAgents = () => {
  return useQuery(["people", { role: [ROLES.AGENT] }], fetchPeople)
}
