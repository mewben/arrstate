import React from "react"
import { Link } from "gatsby"
import { Router } from "@reach/router"

import { Loading, Error, Portal, Button } from "@Components/generic"
import { useLot } from "@Hooks"
import { AppBar, SubMenu, SubMenuItem } from "@Wrappers/layout"
import { LotForm } from "@Components/popups/lot"
import { map } from "@Utils/lodash"
import { Header } from "./components"
import LotOverview from "./lot-overview"

const LotSingle = ({ lotID }) => {
  const { status, data, error } = useLot(lotID)
  const submenu = [
    {
      label: "Overview",
      path: `/lots/${lotID}`,
    },
    {
      label: "Invoices",
      path: `/lots/${lotID}/invoices`,
    },
    {
      label: "Receipts",
      path: `/lots/${lotID}/receipts`,
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
        backTo={data?.projectID ? `/projects/${data?.projectID}/lots` : "/lots"}
        submenu={renderSubmenu()}
      >
        <Portal openByClickOn={<Button>Edit Lot</Button>}>
          <LotForm model={data} />
        </Portal>
      </AppBar>
      {/* <Header lot={data} /> */}
      <Router className="flex-1 overflow-y-scroll pb-28">
        <LotOverview path="/" lot={data} />
      </Router>
    </>
  )
}

export default LotSingle
