import React from "react"

import { List } from "@Screens/receipts/components"

const PropertyReceipts = ({ property }) => {
  return <List propertyID={property._id} />
}

export default PropertyReceipts
