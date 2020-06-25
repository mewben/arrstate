import React from "react"

import { Portal, Button } from "@Components/generic"
import { ProjectForm } from "@Components/popups/project"

import { List } from "./components"

const ProjectList = () => {
  return (
    <div>
      <div>
        <h1>Projects</h1>
        <Portal openByClickOn={<Button>New Project</Button>}>
          <ProjectForm />
        </Portal>
      </div>
      <List />
    </div>
  )
}

export default ProjectList
