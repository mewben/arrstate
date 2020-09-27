import React from "react"

import { ENTITIES } from "@Enums"
import { List } from "@Screens/files/components"

const ProjectFiles = ({ project }) => {
  return <List entityType={ENTITIES.PROJECT} entityID={project._id} />
}

export default ProjectFiles
