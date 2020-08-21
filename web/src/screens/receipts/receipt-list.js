import React from "react"
import { useTranslation } from "react-i18next"

import { AppBar } from "@Wrappers/layout"
import { List } from "./components"

const ReceiptList = () => {
  const { t } = useTranslation()
  return (
    <>
      <AppBar title={t("receipts.title")} />
      <List />
    </>
  )
}

export default ReceiptList
