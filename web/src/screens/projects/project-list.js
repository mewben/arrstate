import React from "react"

import { Portal, Button } from "@Components/generic"
import { ProjectForm } from "@Components/popups/project"

import { AppBar } from "@Wrappers/layout"
import { List } from "./components"

const ProjectList = () => {
  return (
    <>
      <AppBar title="Projects">
        <Portal openByClickOn={<Button>New Project</Button>}>
          <ProjectForm />
        </Portal>
      </AppBar>
      <List />
    </>
  )
}

export default ProjectList
