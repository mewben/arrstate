import React from "react"
import { Router } from "@reach/router"

import { PrivateWrapper } from "@Wrappers"
import { ClientList, ClientSingle } from "@Screens/clients"

const ClientsPage = () => {
  return (
    <PrivateWrapper>
      <Router className="flex flex-col flex-1 overflow-hidden">
        <ClientList path="/clients" />
        <ClientSingle path="/clients/:clientID/*" />
      </Router>
    </PrivateWrapper>
  )
}

export default ClientsPage
