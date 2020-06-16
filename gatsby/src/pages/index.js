import React from "react"
import { Redirect } from "@reach/router"

import { PrivateWrapper } from "@Wrappers"

const IndexPage = () => (
  <PrivateWrapper>
    <Redirect to="/projects" noThrow />
  </PrivateWrapper>
)

export default IndexPage
