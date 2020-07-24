import React from "react"

import { Empty, Portal, Button, Table, Th } from "@Components/generic"
import { InfiniteScroll } from "@Components/infinite-scroll"
import { map } from "@Utils/lodash"
import { useInvoices } from "@Hooks"
import { INVOICE_STATUS } from "@Enums"
import ListItem from "./list-item"

const List = ({ propertyID }) => {
  const renderContent = content => {
    return (
      <div className="p-4">
        <Table>
          <thead>
            <tr>
              <Th fullWidth>Receipt No</Th>
              <Th>Property</Th>
              <Th>Issued To</Th>
              <Th>Issued Date</Th>
              <Th>Paid Date</Th>
              <Th align="right">Total (Php)</Th>
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

  const getMethodParams = {
    status: INVOICE_STATUS.PAID,
  }
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
