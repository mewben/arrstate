import React from "react"

import { PrivateWrapper } from "@Wrappers"
import { useAuth } from "@Providers"

const ProjectsPage = () => {
  const { authSignout } = useAuth()
  return (
    <PrivateWrapper>
      Projects
      <button onClick={() => authSignout()}>Signout</button>
    </PrivateWrapper>
  )
}

export default ProjectsPage
