import React from "react"
import { Link } from "gatsby"
import acc from "accounting"

import { useProject } from "@Hooks"
import { Td } from "@Components/generic"

// Lots listitem
const ListItem = ({ item, projectID }) => {
  return (
    <tr>
      <Td>
        <Link
          to={`/lots/${item._id}`}
          className="font-medium text-cool-gray-700 hover:text-blue-500"
        >
          {item.name}
        </Link>
      </Td>
      {!projectID && (
        <Td>{!!item.projectID && <Project id={item.projectID} />}</Td>
      )}
      <Td align="right">{acc.formatNumber(item.area, 2)}</Td>
      <Td align="right">{acc.formatNumber(item.price, 2)}</Td>
      <Td align="right">{acc.formatNumber(item.priceAddon, 2)}</Td>
    </tr>
  )
}

const Project = ({ id }) => {
  const { status, data } = useProject(id)
  if (status !== "success") return null
  return (
    <Link to={`/projects/${id}`} className="text-gray-700 hover:text-blue-500">
      {data.name}
    </Link>
  )
}

export default ListItem
