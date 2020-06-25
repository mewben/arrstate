import React from "react"
import { Link } from "gatsby"
import { Router } from "@reach/router"

import { Loading } from "@Components/generic"
import { useLot } from "@Hooks"
import { extractError } from "@Utils"
import { Header } from "./components"
import LotOverview from "./lot-overview"

const LotSingle = ({ lotID }) => {
  const { status, data, error } = useLot(lotID)
  return status === "loading" ? (
    <Loading />
  ) : status === "error" ? (
    <div>{extractError(error)}</div>
  ) : (
    <div>
      <Link to="/projects">Back to List of Lots</Link>
      <Header lot={data} />
      <Router>
        <LotOverview path="/" lot={data} />
      </Router>
    </div>
  )
}

export default LotSingle
