import React from "react"
import { Link } from "gatsby"
import { Router } from "@reach/router"

import { Loading } from "@Components/generic"
import { useProject } from "@Hooks"
import { extractError } from "@Utils"
import { Header, SubMenu } from "./components"
import ProjectOverview from "./project-overview"
import ProjectLots from "./project-lots"

const ProjectSingle = ({ projectID }) => {
  const { status, data, error } = useProject(projectID)
  return status === "loading" ? (
    <Loading />
  ) : status === "error" ? (
    <div>{extractError(error)}</div>
  ) : (
    <div>
      <Link to="/projects">Back to List of Projects</Link>
      <Header project={data} />
      <SubMenu projectID={projectID} />
      <Router>
        <ProjectOverview path="/" project={data} />
        <ProjectLots path="lots" project={data} />
      </Router>
    </div>
  )
}

export default ProjectSingle
