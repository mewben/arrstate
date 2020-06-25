import { useQuery } from "react-query"
import { requestApi } from "@Utils"

const fetchLots = async (_, projectID) => {
  const { data } = await requestApi(`/api/lots/${projectID}`)
  return data
}

export const useLots = projectID => {
  return useQuery(["lots", projectID], fetchLots)
}
