import React from "react"

import { Empty } from "@Components/generic"
import { InfiniteScroll } from "@Components/infinite-scroll"
import { map } from "@Utils/lodash"
import { useLots } from "@Hooks"
import ListItem from "./list-item"

const List = () => {
  const renderContent = content => {
    return (
      <div>
        {map(content?.list, item => {
          return <ListItem key={item._id} item={item} />
        })}
      </div>
    )
  }

  const renderEmpty = () => {
    return (
      <div>
        <Empty />
      </div>
    )
  }

  return (
    <InfiniteScroll
      getMethod={useLots}
      contentRenderer={renderContent}
      emptyRenderer={renderEmpty}
    />
  )
}

export default List
