import React from "react"

import { Portal, Button } from "@Components/generic"
import { PropertyForm } from "@Components/popups/property"

const Header = ({ property }) => {
  return (
    <div>
      <h1>{property.name}</h1>
      <Portal openByClickOn={<Button>Edit Property</Button>}>
        <PropertyForm model={property} />
      </Portal>
    </div>
  )
}

export default Header
