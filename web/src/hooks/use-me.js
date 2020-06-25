import { useQuery } from "react-query"
import { requestApi } from "@Utils"

const fetchMe = async () => {
  const { data } = await requestApi("/api/me", "GET")
  return data
}

export const useMe = () => {
  return useQuery("me", fetchMe)
}
