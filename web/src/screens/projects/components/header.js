import React from "react"

import { Portal, Button } from "@Components/generic"
import { ProjectForm } from "@Components/popups/project"

const Header = ({ project }) => {
  return (
    <div>
      <h1>{project.name}</h1>
      <Portal openByClickOn={<Button>Edit Project</Button>}>
        <ProjectForm model={project} />
      </Portal>
    </div>
  )
}

export default Header
