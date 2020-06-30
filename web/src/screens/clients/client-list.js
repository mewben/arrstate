import React from "react"

import { Portal, Button } from "@Components/generic"
import { ClientForm } from "@Components/popups/people"
import { AppBar } from "@Wrappers/layout"

import { useClients } from "@Hooks"
import { List } from "./components"

const ClientList = () => {
  useClients()
  return (
    <>
      <AppBar title="Clients">
        <Portal openByClickOn={<Button>New Client</Button>}>
          <ClientForm />
        </Portal>
      </AppBar>
      <List />
    </>
  )
}

export default ClientList
