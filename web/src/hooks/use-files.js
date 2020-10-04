import { useQuery } from "react-query"
import { requestApi } from "@Utils"

export const useFiles = (params = {}) => {
  return useQuery(["files", params], async (_, params) => {
    const { data } = await requestApi("/api/files", "GET", { params })
    return data
  })
}
