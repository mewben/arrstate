import React from "react"
import { Router } from "@reach/router"

import { PrivateWrapper } from "@Wrappers"
import { ReceiptList, ReceiptSingle } from "@Screens/receipts"

const ReceiptsPage = () => {
  return (
    <PrivateWrapper>
      <Router className="flex flex-col flex-1 overflow-hidden">
        <ReceiptList path="/receipts" />
        <ReceiptSingle path="/receipts/:receiptID/*" />
      </Router>
    </PrivateWrapper>
  )
}

export default ReceiptsPage
