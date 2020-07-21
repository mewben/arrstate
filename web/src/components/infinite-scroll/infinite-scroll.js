import React from "react"
import PropTypes from "prop-types"

import { Loading, Error } from "@Components/generic"
import { ListProvider } from "@Providers"

// TODO: make this into an infinity scroll
const InfiniteScroll = ({
  getMethod,
  getMethodParams,
  contentRenderer,
  emptyRenderer,
}) => {
  const renderContent = data => {
    return (
      <div className="overflow-y-scroll pb-28">{contentRenderer(data)}</div>
    )
  }
  return (
    <ListProvider
      getMethod={getMethod}
      getMethodParams={getMethodParams}
      contentRenderer={renderContent}
      emptyRenderer={emptyRenderer}
    />
  )
}

InfiniteScroll.propTypes = {
  getMethod: PropTypes.func.isRequired,
  getMethodParams: PropTypes.any,
  contentRenderer: PropTypes.any.isRequired,
  emptyRenderer: PropTypes.any,
}

export default InfiniteScroll
