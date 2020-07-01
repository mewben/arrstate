import React from "react"
import { Router } from "@reach/router"

import { PrivateWrapper } from "@Wrappers"
import { PersonList, PersonSingle } from "@Screens/people"

const PeoplePage = () => {
  return (
    <PrivateWrapper>
      <Router className="flex flex-col flex-1 overflow-hidden">
        <PersonList path="/people" />
        <PersonSingle path="/people/:personID/*" />
      </Router>
    </PrivateWrapper>
  )
}

export default PeoplePage
