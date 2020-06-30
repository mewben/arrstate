import React from "react"
import { Router } from "@reach/router"

import { PrivateWrapper } from "@Wrappers"
import { AgentList, AgentSingle } from "@Screens/agents"

const AgentsPage = () => {
  return (
    <PrivateWrapper>
      <Router className="flex flex-col flex-1 overflow-hidden">
        <AgentList path="/agents" />
        <AgentSingle path="/agents/:agentID/*" />
      </Router>
    </PrivateWrapper>
  )
}

export default AgentsPage
