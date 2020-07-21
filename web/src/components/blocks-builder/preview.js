import React from "react"

import { useBlocks } from "@Hooks"
import { ListProvider } from "@Providers"

const BlocksPreview = ({ blocks = [] }) => {
  const prepareMethodParams = () => {
    return {
      ids: blocks,
    }
  }

  const renderBlocks = ({ list: blocks }) => {
    console.log("bloccccks", blocks)
    return <div>renderBlocks</div>
  }

  return (
    <ListProvider
      getMethod={useBlocks}
      getMethodParams={prepareMethodParams()}
      contentRenderer={renderBlocks}
    />
  )
}

export default BlocksPreview
