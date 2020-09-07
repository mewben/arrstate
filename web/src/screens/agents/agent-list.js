import React from "react"

import { Portal, Button } from "@Components/generic"
import { AgentForm } from "@Components/popups/person"
import { AppBar } from "@Wrappers/layout"

import { useAgents } from "@Hooks"
import { List } from "./components"

const AgentList = () => {
  useAgents()
  return (
    <>
      <AppBar title="Agents">
        <Portal openByClickOn={<Button>New Agent</Button>}>
          <AgentForm />
        </Portal>
      </AppBar>
      <List />
    </>
  )
}

export default AgentList
