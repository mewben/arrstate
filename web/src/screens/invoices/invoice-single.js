import React from "react"

import { useInvoice } from "@Hooks"
import { SingleProvider } from "@Providers"
import HeaderSingle from "./components/header-single"
import Preview from "./components/preview"

const InvoiceSingle = ({ invoiceID }) => {
  const renderContent = invoice => {
    return (
      <>
        <HeaderSingle invoice={invoice} />
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
