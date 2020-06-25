import React from "react"

import { useLots } from "@Hooks"
import { Empty } from "@Components/generic"
import { map } from "@Utils/lodash"
import { ListItem } from "./components/lots"
import { InfiniteScroll } from "@Components/infinite-scroll"

const ProjectLots = ({ project }) => {
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
      getMethodParams={project._id}
      contentRenderer={renderContent}
      emptyRenderer={renderEmpty}
    />
  )
}

export default ProjectLots
