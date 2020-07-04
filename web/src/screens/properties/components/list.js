import React from "react"

import { Empty, Portal, Button, Table, Th } from "@Components/generic"
import { InfiniteScroll } from "@Components/infinite-scroll"
import { map } from "@Utils/lodash"
import { useProperties } from "@Hooks"
import { PropertyForm } from "@Components/popups/property"
import ListItem from "./list-item"

const List = ({ projectID }) => {
  const renderAdd = () => {
    if (!projectID) return null

    return (
      <Portal openByClickOn={<Button>Add Property</Button>}>
        <PropertyForm projectID={projectID} />
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
              <Th>Property No.</Th>
              {!projectID && <Th>Project</Th>}
              <Th align="right">Area</Th>
              <Th align="right">Price</Th>
              <Th align="right">Price Addon</Th>
            </tr>
          </thead>
          <tbody className="bg-white">
            {map(content?.list, item => {
              return (
                <ListItem key={item._id} item={item} projectID={projectID} />
              )
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
      getMethod={useProperties}
      getMethodParams={getMethodParams}
      contentRenderer={renderContent}
      emptyRenderer={renderEmpty}
    />
  )
}

export default List
