import React from "react"

import { List } from "@Screens/properties/components"

const ProjectProperties = ({ project }) => {
  return <List projectID={project._id} />
}

export default ProjectProperties
