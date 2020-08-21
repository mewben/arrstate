import React from "react"
import { useTranslation } from "react-i18next"

import { AppBar, SubMenu, SubMenuItem } from "@Wrappers/layout"
import { Loading, Error, Portal, Button } from "@Components/generic"
import { PayForm } from "@Components/popups/invoice"
import { INVOICE_STATUS } from "@Enums"
import Status from "./status"

const HeaderSingle = ({ invoice }) => {
  const { t } = useTranslation()
  const renderTitle = () => {
    return (
      <div className="flex space-x-3 items-center">
        <h1>{invoice.name}</h1>
        <Status status={invoice.status} />
      </div>
    )
  }

  const renderAction = () => {
    switch (invoice.status) {
      case INVOICE_STATUS.PENDING:
        return (
          <Portal openByClickOn={<Button>{t("btnPay")}</Button>}>
            <PayForm invoice={invoice} />
          </Portal>
        )
      case INVOICE_STATUS.PAID:
        return (
          <Button to={`/receipts/${invoice._id}`} color="white">
            {t("btnViewReceipt")}
          </Button>
        )
      default:
        return null
    }
  }

  return (
    <AppBar
      title={renderTitle()}
      backTo={
        invoice?.propertyID
          ? `/properties/${invoice?.propertyID}/invoices`
          : "/invoices"
      }
    >
      {renderAction()}
    </AppBar>
  )
}

export default HeaderSingle
