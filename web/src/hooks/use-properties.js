import { useQuery } from "react-query"
import { requestApi } from "@Utils"

const fetchProperties = async (_, params) => {
  const { data } = await requestApi("/api/properties", "GET", { params })
  return data
}

const fetchProperty = async (_, propertyID) => {
  const { data } = await requestApi(`/api/properties/${propertyID}`)
  return data
}

export const useProperties = (params = {}) => {
  return useQuery(["properties", params], fetchProperties)
}

export const useProperty = propertyID => {
  return useQuery(["property", propertyID], fetchProperty)
}
