import React from "react"
import { Link } from "gatsby"
import { Router } from "@reach/router"

import { Loading, Error } from "@Components/generic"
import { useLot } from "@Hooks"
import { Header } from "./components"
import LotOverview from "./lot-overview"

const LotSingle = ({ lotID }) => {
  const { status, data, error } = useLot(lotID)
  return status === "loading" ? (
    <Loading />
  ) : status === "error" ? (
    <Error error={error} />
  ) : (
    <div>
      <Link to="/lots">Back to List of Lots</Link>
      <Header lot={data} />
      <Router>
        <LotOverview path="/" lot={data} />
      </Router>
    </div>
  )
}

export default LotSingle
