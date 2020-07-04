import React from "react"
import { Router } from "@reach/router"

import { PrivateWrapper } from "@Wrappers"
import { PropertyList, PropertySingle } from "@Screens/properties"

const PropertiesPage = () => {
  return (
    <PrivateWrapper>
      <Router className="flex flex-col flex-1 overflow-hidden">
        <PropertyList path="/properties" />
        <PropertySingle path="/properties/:propertyID/*" />
      </Router>
    </PrivateWrapper>
  )
}

export default PropertiesPage
