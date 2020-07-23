import React from "react"

import { AppBar } from "@Wrappers/layout"
import { List } from "./components"

const InvoiceList = () => {
  return (
    <>
      <AppBar title="Invoices"></AppBar>
      <List />
    </>
  )
}

export default InvoiceList
