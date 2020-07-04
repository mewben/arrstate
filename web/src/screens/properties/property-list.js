import React from "react"

import { Portal, Button } from "@Components/generic"
import { PropertyForm } from "@Components/popups/property"
import { AppBar } from "@Wrappers/layout"

import { useProjects } from "@Hooks"
import { List } from "./components"

const PropertyList = () => {
  useProjects()
  return (
    <>
      <AppBar title="Properties">
        <Portal openByClickOn={<Button>New Property</Button>}>
          <PropertyForm />
        </Portal>
      </AppBar>
      <List />
    </>
  )
}

export default PropertyList
