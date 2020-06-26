import React from "react"

import { Empty, Portal, Button } from "@Components/generic"
import { InfiniteScroll } from "@Components/infinite-scroll"
import { map } from "@Utils/lodash"
import { useLots } from "@Hooks"
import { LotForm } from "@Components/popups/lot"
import ListItem from "./list-item"

const List = ({ projectID }) => {
  const renderAdd = () => {
    return (
      <Portal openByClickOn={<Button>Add Lot</Button>}>
        <LotForm projectID={projectID} />
      </Portal>
    )
  }

  const renderContent = content => {
    return (
      <div>
        <div>{renderAdd()}</div>
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
        {renderAdd()}
      </div>
    )
  }

  const getMethodParams = {}
  if (projectID === null || !!projectID) {
    getMethodParams["projectID"] = projectID
  }

  return (
    <InfiniteScroll
      getMethod={useLots}
      getMethodParams={getMethodParams}
      contentRenderer={renderContent}
      emptyRenderer={renderEmpty}
    />
  )
}

export default List
