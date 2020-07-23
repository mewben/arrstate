import React from "react"
import { Link } from "gatsby"
import acc from "accounting"

import { Td } from "@Components/generic"

const ListItem = ({ item }) => {
  return (
    <tr>
      <Td wrap>
        <Link
          to={`/projects/${item._id}/properties`}
          className="font-medium text-cool-gray-700 hover:text-blue-500"
        >
          {item.name}
        </Link>
      </Td>
      <Td align="right">
        {acc.formatNumber(item.area, 2)} {item.unit}
      </Td>
    </tr>
  )
}

export default ListItem
