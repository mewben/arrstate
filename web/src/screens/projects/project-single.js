import React from "react"
import { Link } from "gatsby"
import { Router } from "@reach/router"

import { Loading, Error } from "@Components/generic"
import { useProject } from "@Hooks"
import { AppBar } from "@Wrappers/layout"
import { Header, SubMenu } from "./components"
import ProjectOverview from "./project-overview"
import ProjectLots from "./project-lots"

const ProjectSingle = ({ projectID }) => {
  const { status, data, error } = useProject(projectID)
  return status === "loading" ? (
    <Loading />
  ) : status === "error" ? (
    <Error error={error} />
  ) : (
    <>
      <AppBar title={data.name}>
        {/* <Link to="/projects">Back to List of Projects</Link>
        <Header project={data} />
        <SubMenu projectID={projectID} /> */}
      </AppBar>
      <Router>
        <ProjectOverview path="/" project={data} />
        <ProjectLots path="lots" project={data} />
      </Router>
    </>
  )
}

export default ProjectSingle
