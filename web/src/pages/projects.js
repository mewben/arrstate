import React from "react"
import { Router } from "@reach/router"

import { PrivateWrapper } from "@Wrappers"
import { ProjectList, ProjectSingle } from "@Screens/projects"

const ProjectsPage = () => {
  return (
    <PrivateWrapper>
      <Router className="flex flex-col flex-1 overflow-hidden">
        <ProjectList path="/projects" />
        <ProjectSingle path="/projects/:projectID/*" />
      </Router>
    </PrivateWrapper>
  )
}

export default ProjectsPage
