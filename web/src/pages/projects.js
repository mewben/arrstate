import React from "react"
import { Router } from "@reach/router"

import { PrivateWrapper } from "@Wrappers"
import { useAuth } from "@Providers"
import { ProjectList, ProjectSingle } from "@Screens/projects"

const ProjectsPage = () => {
  const { authSignout } = useAuth()
  return (
    <PrivateWrapper>
      <Router>
        <ProjectList path="/projects" />
        <ProjectSingle path="/projects/:projectID/*" />
      </Router>
    </PrivateWrapper>
  )
}

export default ProjectsPage
