import React from "react"

import { List } from "@Screens/invoices/components"

const PropertyInvoices = ({ property }) => {
  return <List propertyID={property._id} />
}

export default PropertyInvoices
