import React from "react"
import { Router } from "@reach/router"

import { PrivateWrapper } from "@Wrappers"
import { InvoiceList, InvoiceSingle } from "@Screens/invoices"

const InvoicesPage = () => {
  return (
    <PrivateWrapper>
      <Router className="flex flex-col flex-1 overflow-hidden">
        <InvoiceList path="/invoices" />
        <InvoiceSingle path="/invoices/:invoiceID/*" />
      </Router>
    </PrivateWrapper>
  )
}

export default InvoicesPage
