import { useQuery } from "react-query"
import { requestApi } from "@Utils"

const fetchLots = async (_, params) => {
  const { data } = await requestApi("/api/lots", "GET", { params })
  return data
}

const fetchLot = async (_, lotID) => {
  const { data } = await requestApi(`/api/lots/${lotID}`)
  return data
}

export const useLots = (params = {}) => {
  return useQuery(["lots", params], fetchLots)
}

export const useLot = lotID => {
  return useQuery(["lot", lotID], fetchLot)
}
