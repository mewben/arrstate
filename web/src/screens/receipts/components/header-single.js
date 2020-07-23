import React from "react"

import { AppBar } from "@Wrappers/layout"

const HeaderSingle = ({ receipt }) => {
  return (
    <AppBar
      title={`Receipt - ${receipt.receiptNo}`}
      backTo={`/invoices/${receipt._id}`}
    />
  )
}

export default HeaderSingle
