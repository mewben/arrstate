import { useQuery } from "react-query"
import { requestApi } from "@Utils"

const fetchBlocks = async (_, payload) => {
  const { data } = await requestApi("/api/blocks/get", "POST", {
    data: payload,
  })
  return data
}

export const useBlocks = (payload = {}) => {
  return useQuery(["blocks", payload], fetchBlocks)
}
