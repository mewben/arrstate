import React from "react"
import { useQuery } from "react-query"
import { requestApi } from "@Utils"

const fetchCountries = async () => {
  const { data } = await requestApi("/api/businesses/countries")
  return data
}

export const useCountries = () => {
  return useQuery(["countries"], fetchCountries)
}
