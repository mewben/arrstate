import React from "react"
import { Link } from "gatsby"
import { Router } from "@reach/router"

import { Loading, Error, Portal, Button } from "@Components/generic"
import { useProperty } from "@Hooks"
import { AppBar, SubMenu, SubMenuItem } from "@Wrappers/layout"
import { PropertyForm } from "@Components/popups/property"
import { map } from "@Utils/lodash"
import PropertyOverview from "./property-overview"

const PropertySingle = ({ propertyID }) => {
  const { status, data, error } = useProperty(propertyID)
  const submenu = [
    {
      label: "Overview",
      path: `/properties/${propertyID}`,
    },
    {
      label: "Invoices",
      path: `/properties/${propertyID}/invoices`,
    },
    {
      label: "Receipts",
      path: `/properties/${propertyID}/receipts`,
    },
  ]

  const renderSubmenu = () => {
    return (
      <SubMenu>
        {map(submenu, (item, i) => {
          return (
            <SubMenuItem key={i} to={item.path}>
              {item.label}
            </SubMenuItem>
          )
        })}
      </SubMenu>
    )
  }

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
        submenu={renderSubmenu()}
      >
        <Portal openByClickOn={<Button>Edit Property</Button>}>
          <PropertyForm model={data} />
        </Portal>
      </AppBar>
      <Router className="flex-1 overflow-y-scroll pb-28">
        <PropertyOverview path="/" property={data} />
      </Router>
    </>
  )
}

export default PropertySingle
