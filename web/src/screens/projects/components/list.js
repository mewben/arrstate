import React from "react"

import { Loading, Empty } from "@Components/generic"
import { extractError } from "@Utils"
import { map } from "@Utils/lodash"
import { useProjects } from "@Hooks"
import ListItem from "./list-item"

const List = () => {
  const { status, data, error } = useProjects()
  return status === "loading" ? (
    <Loading />
  ) : status === "error" ? (
    <div>{extractError(error)}</div>
  ) : !data?.total ? (
    <Empty />
  ) : (
    <div>
      {map(data?.list, item => {
        return <ListItem key={item._id} item={item} />
      })}
    </div>
  )
}

export default List
