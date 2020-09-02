import React from "react"
import { useTranslation } from "react-i18next"

import { AppBar } from "@Wrappers/layout"
import { List } from "./components"

const InvoiceList = () => {
  const { t } = useTranslation()
  return (
    <>
      <AppBar title={t("invoices.title")} />
      <List />
    </>
  )
}

export default InvoiceList
