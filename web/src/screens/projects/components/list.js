import React from "react"

import { Empty, Portal, Button, Table, Th } from "@Components/generic"
import { InfiniteScroll } from "@Components/infinite-scroll"
import { map } from "@Utils/lodash"
import { useProjects } from "@Hooks"
import { ProjectForm } from "@Components/popups/project"
import ListItem from "./list-item"

const List = () => {
  const renderAdd = () => {
    return (
      <Portal openByClickOn={<Button>Add Project</Button>}>
        <ProjectForm />
      </Portal>
    )
  }

  const renderContent = content => {
    return (
      <div className="p-4">
        <Table>
          <thead>
            <tr>
              <Th>Project Name</Th>
              <Th align="right">Area</Th>
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
      getMethod={useProjects}
      contentRenderer={renderContent}
      emptyRenderer={renderEmpty}
    />
  )
}

export default List
