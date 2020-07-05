import { useQuery } from "react-query"
import { requestApi, buildParams } from "@Utils"
import { ROLES } from "@Enums"

const fetchPeople = async (_, params) => {
  const q = buildParams(params)
  const { data } = await requestApi("/api/people", "GET", {
    params: q,
  })
  return data
}

export const usePeople = () => {
  return useQuery(["people"], fetchPeople)
}

export const useClients = () => {
  return useQuery(["people", { role: [ROLES.CLIENT] }], fetchPeople)
}

export const useAgents = () => {
  return useQuery(["people", { role: [ROLES.AGENT] }], fetchPeople)
}
