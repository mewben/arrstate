import React from "react"

import { useAuth } from "@Providers"
import { navigate } from "gatsby"

const SignoutPage = () => {
  const { authSignout } = useAuth()

  React.useEffect(() => {
    authSignout()
    navigate("/", { replace: true })
  }, [])

  return <div>Signing out...</div>
}

export default SignoutPage
