import React from "react"
import { Link } from "gatsby"
import { Router } from "@reach/router"

import { Loading, Error, Portal, Button } from "@Components/generic"
import { ENTITIES, INVOICE_BLOCKS } from "@Enums"
import { BlocksPreview } from "@Components/blocks-builder"
import { useInvoice } from "@Hooks"
import { AppBar, SubMenu, SubMenuItem } from "@Wrappers/layout"
import { SingleProvider } from "@Providers"
import { PayForm } from "@Components/popups/invoice"
import Preview from "./components/preview"

const InvoiceSingle = ({ invoiceID }) => {
  const renderContent = invoice => {
    console.log("invoice", invoice)
    return (
      <>
        <AppBar
          title={invoice.name}
          backTo={
            invoice?.propertyID
              ? `/properties/${invoice?.propertyID}/invoices`
              : "/invoices"
          }
        >
          <Portal openByClickOn={<Button>Pay</Button>}>
            <PayForm invoice={invoice} />
          </Portal>
        </AppBar>
        <div className="flex-1 overflow-y-scroll pb-28">
          <Preview invoice={invoice} />
        </div>
      </>
    )
  }

  return (
    <SingleProvider
      getMethod={useInvoice}
      id={invoiceID}
      contentRenderer={renderContent}
    />
  )
}

export default InvoiceSingle
