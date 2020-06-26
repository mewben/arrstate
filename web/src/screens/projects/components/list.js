import React from "react"

import { Empty, Portal, Button } from "@Components/generic"
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

  return (
    <InfiniteScroll
      getMethod={useProjects}
      contentRenderer={renderContent}
      emptyRenderer={renderEmpty}
    />
  )
}

export default List
