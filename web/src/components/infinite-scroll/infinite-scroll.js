import React from "react"
import PropTypes from "prop-types"

import { Loading, Error } from "@Components/generic"

// TODO: make this into an infinity scroll
const InfiniteScroll = ({
  getMethod,
  getMethodParams,
  contentRenderer,
  emptyRenderer,
}) => {
  const { status, data, error } = getMethod(getMethodParams)
  console.log("status:", status)
  console.log("data:", data)
  console.log("error:", error)
  return status === "loading" ? (
    <Loading />
  ) : status === "error" ? (
    <Error error={error} />
  ) : !data?.total ? (
    emptyRenderer()
  ) : (
    contentRenderer(data)
  )
}

InfiniteScroll.propTypes = {
  getMethod: PropTypes.func.isRequired,
  getMethodParams: PropTypes.any,
  contentRenderer: PropTypes.any.isRequired,
  emptyRenderer: PropTypes.any,
}

export default InfiniteScroll
