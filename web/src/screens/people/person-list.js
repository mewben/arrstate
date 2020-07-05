import React from "react"

import { Portal, Button } from "@Components/generic"
import { PersonForm } from "@Components/popups/people"
import { AppBar } from "@Wrappers/layout"
import { List } from "./components"

const PersonList = () => {
  return (
    <>
      <AppBar title="People">
        <Portal openByClickOn={<Button>New Person</Button>}>
          <PersonForm />
        </Portal>
      </AppBar>
      <List />
    </>
  )
}

export default PersonList
