import React from "react"

import { Portal, Button } from "@Components/generic"
import { LotForm } from "@Components/popups/lot"
import { AppBar } from "@Wrappers/layout"

import { useProjects } from "@Hooks"
import { List } from "./components"

const LotList = () => {
  useProjects()
  return (
    <>
      <AppBar title="Lots">
        <Portal openByClickOn={<Button>New Lot</Button>}>
          <LotForm />
        </Portal>
      </AppBar>
      <List />
    </>
  )
}

export default LotList
