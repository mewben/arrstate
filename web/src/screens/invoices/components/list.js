import React from "react"

import { Empty, Portal, Button, Table, Th } from "@Components/generic"
import { InfiniteScroll } from "@Components/infinite-scroll"
import { map } from "@Utils/lodash"
import { useInvoices } from "@Hooks"
import ListItem from "./list-item"

const List = ({ propertyID }) => {
  const renderContent = content => {
    return (
      <div className="p-4">
        <Table>
          <thead>
            <tr>
              <Th fullWidth>Invoice No.</Th>
              {!propertyID && <Th>Property</Th>}
              <Th>Issued To</Th>
              <Th>Issue Date</Th>
              <Th>Due Date</Th>
              <Th align="right">Total (Php)</Th>
              <Th>Status</Th>
            </tr>
          </thead>
          <tbody className="bg-white">
            {map(content?.list, item => {
              return (
                <ListItem key={item._id} item={item} propertyID={propertyID} />
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
      </div>
    )
  }

  const getMethodParams = {}
  if (propertyID === null || !!propertyID) {
    getMethodParams["propertyID"] = propertyID
  }

  return (
    <InfiniteScroll
      getMethod={useInvoices}
      getMethodParams={getMethodParams}
      contentRenderer={renderContent}
      emptyRenderer={renderEmpty}
    />
  )
}

export default List
