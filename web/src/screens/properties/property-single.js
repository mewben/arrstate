import React from "react"
import { Link } from "gatsby"
import { Router } from "@reach/router"
import { useTranslation } from "react-i18next"

import { Loading, Error, Portal, Button } from "@Components/generic"
import { useProperty } from "@Hooks"
import { AppBar } from "@Wrappers/layout"
import { PropertyForm } from "@Components/popups/property"
import PropertyOverview from "./property-overview"
import PropertyInvoices from "./property-invoices"
import PropertyReceipts from "./property-receipts"

const PropertySingle = ({ propertyID }) => {
  const { t } = useTranslation()
  const { status, data, error } = useProperty(propertyID)
  const submenu = [
    {
      label: t("properties.menu.overview"),
      path: `/properties/${propertyID}`,
    },
    {
      label: t("properties.menu.invoices"),
      path: `/properties/${propertyID}/invoices`,
    },
    {
      label: t("properties.menu.receipts"),
      path: `/properties/${propertyID}/receipts`,
    },
  ]

  return status === "loading" ? (
    <Loading />
  ) : status === "error" ? (
    <Error error={error} />
  ) : (
    <>
      <AppBar
        title={data.name}
        backTo={
          data?.projectID
            ? `/projects/${data?.projectID}/properties`
            : "/properties"
        }
        submenu={submenu}
      >
        <Portal openByClickOn={<Button>{t("properties.edit")}</Button>}>
          <PropertyForm model={data} />
        </Portal>
      </AppBar>
      <Router className="flex-1 overflow-y-scroll pb-28">
        <PropertyOverview path="/" property={data} />
        <PropertyInvoices path="/invoices" property={data} />
        <PropertyReceipts path="/receipts" property={data} />
      </Router>
    </>
  )
}

export default PropertySingle
