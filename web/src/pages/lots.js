import React from "react"
import { Router } from "@reach/router"

import { PrivateWrapper } from "@Wrappers"
import { LotList, LotSingle } from "@Screens/lots"

const LotsPage = () => {
  return (
    <PrivateWrapper>
      <Router className="flex flex-col flex-1 overflow-hidden">
        <LotList path="/lots" />
        <LotSingle path="/lots/:lotID/*" />
      </Router>
    </PrivateWrapper>
  )
}

export default LotsPage
