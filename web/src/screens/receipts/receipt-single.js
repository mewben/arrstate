import React from "react"

import { useInvoice } from "@Hooks"
import { SingleProvider } from "@Providers"
import { INVOICE_STATUS } from "@Enums"
import { Error } from "@Components/generic"
import Preview from "@Screens/invoices/components/preview"
import HeaderSingle from "./components/header-single"

const ReceiptSingle = ({ receiptID }) => {
  const renderContent = receipt => {
    console.log("receipt", receipt)
    if (receipt.status !== INVOICE_STATUS.PAID) {
      return <Error error={{ message: "Not found" }} />
    }

    return (
      <>
        <HeaderSingle receipt={receipt} />
        <div className="flex-1 overflow-y-scroll pb-28">
          <Preview invoice={receipt} isReceipt />
        </div>
      </>
    )
  }

  return (
    <SingleProvider
      getMethod={useInvoice}
      id={receiptID}
      contentRenderer={renderContent}
    />
  )
}

export default ReceiptSingle
