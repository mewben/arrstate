import React from "react"
import { navigate, useLocation } from "@reach/router"
import { useMutation } from "react-query"

import { useAuth } from "@Providers"
import { Loading, Error } from "@Components/generic"
import { requestApi } from "@Utils"

// this catches the deviceCode after signup
const CbPage = () => {
  const [loading, setLoading] = React.useState(false)
  const { authSignIn } = useAuth()
  const location = useLocation()
  const params = new URLSearchParams(location.search)
  const [mutate, { reset, error }] = useMutation(formData => {
    return requestApi("/auth/signin", "POST", {
      data: formData,
      noToken: true,
    })
  })

  const deviceCode = params.get("deviceCode")

  React.useEffect(() => {
    const signin = async () => {
      if (deviceCode) {
        // login here to get token
        const res = await mutate({
          grant_type: "device_code",
          deviceCode,
        })

        if (res) {
          authSignIn(res.data.token)
          navigate("/", { replace: true })
        }
        setLoading(false)
      }
    }
    signin()
  }, [deviceCode])

  return loading ? (
    <Loading />
  ) : !!error ? (
    <Error error={error} />
  ) : (
    "Redirecting..."
  )
}

export default CbPage
