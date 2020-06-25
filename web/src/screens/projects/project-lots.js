import React from "react"

import { useLots } from "@Hooks"
import { Empty, Portal, Button } from "@Components/generic"
import { InfiniteScroll } from "@Components/infinite-scroll"
import { LotForm } from "@Components/popups/lot"
import { map } from "@Utils/lodash"
import { ListItem } from "./components/lots"

const ProjectLots = ({ project }) => {
  const renderAdd = () => {
    return (
      <Portal openByClickOn={<Button>Add Lot</Button>}>
        <LotForm project={project} />
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
      getMethod={useLots}
      getMethodParams={{ projectID: project._id }}
      contentRenderer={renderContent}
      emptyRenderer={renderEmpty}
    />
  )
}

export default ProjectLots
