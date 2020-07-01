import React from "react"

import { Empty, Portal, Button, Table, Th } from "@Components/generic"
import { InfiniteScroll } from "@Components/infinite-scroll"
import { map } from "@Utils/lodash"
import { useClients } from "@Hooks"
import { ClientForm } from "@Components/popups/people"
import ListItem from "./list-item"

const List = () => {
  const renderAdd = () => {
    return (
      <Portal openByClickOn={<Button>Add Client</Button>}>
        <ClientForm />
      </Portal>
    )
  }

  const renderContent = content => {
    return (
      <div className="p-4">
        <Table>
          <thead>
            <tr>
              <Th>Name</Th>
            </tr>
          </thead>
          <tbody className="bg-white">
            {map(content?.list, (item, index) => {
              return <ListItem key={item._id} item={item} index={index} />
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

  return (
    <InfiniteScroll
      getMethod={useClients}
      contentRenderer={renderContent}
      emptyRenderer={renderEmpty}
    />
  )
}

export default List
