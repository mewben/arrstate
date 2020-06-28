import React from "react"

import { Empty, Portal, Button, Table, Th } from "@Components/generic"
import { InfiniteScroll } from "@Components/infinite-scroll"
import { map } from "@Utils/lodash"
import { useLots } from "@Hooks"
import { LotForm } from "@Components/popups/lot"
import ListItem from "./list-item"

const List = ({ projectID }) => {
  const renderAdd = () => {
    if (!projectID) return null

    return (
      <Portal openByClickOn={<Button>Add Lot</Button>}>
        <LotForm projectID={projectID} />
      </Portal>
    )
  }

  const renderContent = content => {
    return (
      <div className="p-4">
        <div>{renderAdd()}</div>
        <Table>
          <thead>
            <tr>
              <Th>Lot No.</Th>
              <Th align="right">Area</Th>
              <Th align="right">Price</Th>
              <Th align="right">Price Addon</Th>
            </tr>
          </thead>
          <tbody className="bg-white">
            {map(content?.list, item => {
              return <ListItem key={item._id} item={item} />
            })}
          </tbody>
        </Table>
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
