import React from "react"

import { List } from "@Screens/lots/components"

const ProjectLots = ({ project }) => {
  return <List projectID={project._id} />
}

export default ProjectLots
