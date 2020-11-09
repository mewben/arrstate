import { useLocation } from "@reach/router"

export const useParams = () => {
  const location = useLocation()
  return new URLSearchParams(location?.search)
}
