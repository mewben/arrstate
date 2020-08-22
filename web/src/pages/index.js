import React from "react"

import { PrivateWrapper } from "@Wrappers"
import { Dashboard } from "@Screens/dashboard"

const IndexPage = () => (
  <PrivateWrapper>
    <Dashboard />
  </PrivateWrapper>
)

export default IndexPage
