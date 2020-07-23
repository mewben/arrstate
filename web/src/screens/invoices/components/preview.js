import React from "react"

import { BlocksPreview } from "@Components/blocks-builder"
import { ENTITIES, INVOICE_BLOCKS } from "@Enums"
import { Intro, Item, ItemWrapper, Summary } from "./blocks/preview"

const Preview = ({ invoice }) => {
  const renderBlock = block => {
    let content = null
    switch (block.type) {
      case INVOICE_BLOCKS.INTRO:
        content = <Intro block={block} invoice={invoice} />
        break
      case INVOICE_BLOCKS.ITEM:
        content = <Item block={block} />
        break
      case INVOICE_BLOCKS.SUMMARY:
        content = <Summary block={block} invoice={invoice} />
        break
      default:
        break
    }
    return { content }
  }

  return (
    <BlocksPreview
      blocks={invoice?.blocks}
      entityType={ENTITIES.INVOICE}
      entityID={invoice?._id}
      renderBlock={renderBlock}
      groupWrapper={{
        [INVOICE_BLOCKS.ITEM]: ItemWrapper,
      }}
    />
  )
}

export default Preview
