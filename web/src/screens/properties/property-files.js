import React from "react"

import { ENTITIES } from "@Enums"
import { List } from "@Screens/files/components"

const PropertyFiles = ({ property }) => {
  return <List entityType={ENTITIES.PROPERTY} entityID={property._id} />
}

export default PropertyFiles
