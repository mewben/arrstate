import React from "react"

import { AppBar } from "@Wrappers/layout"
import { List } from "./components"

const ReceiptList = () => {
  return (
    <>
      <AppBar title="Receipts"></AppBar>
      <List />
    </>
  )
}

export default ReceiptList
