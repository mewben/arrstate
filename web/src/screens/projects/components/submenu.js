import React from "react"
import { Link } from "gatsby"

const SubMenu = ({ projectID }) => {
  return (
    <div>
      <Link to={`/projects/${projectID}`}>Overview</Link>
      <Link to={`/projects/${projectID}/lots`}>Lots</Link>
      <Link to={`/projects/${projectID}/agents`}>Agents</Link>
    </div>
  )
}

export default SubMenu
