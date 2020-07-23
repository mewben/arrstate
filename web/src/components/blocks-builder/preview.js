import React from "react"
import PropTypes from "prop-types"

import { useBlocks } from "@Hooks"
import { ListProvider } from "@Providers"
import { GENERIC_BLOCKS } from "@Enums"
import { map, get, isEmpty, includes } from "@Utils/lodash"
import { Content } from "./blocks/preview"
import { groupBlocks } from "@Components/blocks-builder/blocks/preview/helpers"

const BlocksPreview = ({
  entityType,
  entityID,
  blocks = [],
  renderBlock,
  blockProps = {},
  groupWrapper = {},
}) => {
  const prepareMethodParams = () => {
    return {
      ids: blocks,
      entityType,
      entityID,
    }
  }

  const renderSingleBlock = block => {
    let content = null
    switch (block.type) {
      case GENERIC_BLOCKS.CONTENT:
        content = <Content item={block} {...blockProps[block.type]} />
        break
      default:
        return null
    }

    return {
      content,
    }
  }

  const renderBlocks = blocks => {
    console.log("bloccccks", blocks)

    return map(blocks, block => {
      const data =
        (renderBlock && renderBlock(block)) || renderSingleBlock(block)
      if (!data) {
        return null
      }
      const { content } = data
      return content
    })
  }

  const prepareBlocks = ({ list: blocks }) => {
    // groupblocks
    const groupedBlocks = groupBlocks(blocks)
    console.log("groupedBlocks", groupedBlocks)

    return map(groupedBlocks, groupedBlock => {
      const Wrapper = groupWrapper[get(groupedBlock, [0, "type"])]
      if (Wrapper) {
        return <Wrapper>{renderBlocks(groupedBlock)}</Wrapper>
      } else {
        return (
          <div className="py-4 first:pt-0 last:pb-0">
            {renderBlocks(groupedBlock)}
          </div>
        )
      }
    })
  }

  const renderMain = () => {
    return (
      <ListProvider
        getMethod={useBlocks}
        getMethodParams={prepareMethodParams()}
        contentRenderer={prepareBlocks}
      />
    )
  }

  return (
    <div className="mx-auto w-full max-w-screen-md px-16 py-8">
      <div className="shadow-sm bg-white rounded-sm p-16">{renderMain()}</div>
    </div>
  )
}

BlocksPreview.propTypes = {
  entityType: PropTypes.string.isRequired,
  entityID: PropTypes.string.isRequired,
  blocks: PropTypes.arrayOf(PropTypes.string),
  renderBlock: PropTypes.func,
  blockProps: PropTypes.object,
}

export default BlocksPreview
