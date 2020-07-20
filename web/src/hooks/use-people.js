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

const fetchPerson = async (_, personID) => {
  const { data } = await requestApi(`/api/people/${personID}`)
  return data
}

export const usePeople = ({ role = [] } = {}) => {
  return useQuery(["people", { role }], fetchPeople)
}

export const useClients = () => {
  return useQuery(["people", { role: [ROLES.CLIENT] }], fetchPeople)
}

export const useAgents = () => {
  return useQuery(["people", { role: [ROLES.AGENT] }], fetchPeople)
}

export const usePerson = personID => {
  return useQuery(["people", personID], fetchPerson)
}
