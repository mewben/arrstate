import React from "react"
import { Router } from "@reach/router"

import { PrivateWrapper } from "@Wrappers"
import { useAuth } from "@Providers"
import { LotList, LotSingle } from "@Screens/lots"

const LotsPage = () => {
  const { authSignout } = useAuth()
  return (
    <PrivateWrapper>
      <Router>
        <LotList path="/lots" />
        <LotSingle path="/lots/:lotID/*" />
      </Router>
      <hr />
      <button onClick={() => authSignout()}>Signout</button>
    </PrivateWrapper>
  )
}

export default LotsPage
