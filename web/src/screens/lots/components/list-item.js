import React from "react"
import { Link } from "gatsby"
import acc from "accounting"

import { Td } from "@Components/generic"

// Lots listitem
const ListItem = ({ item }) => {
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
      <Td align="right">{acc.formatNumber(item.area, 2)}</Td>
      <Td align="right">{acc.formatNumber(item.price, 2)}</Td>
      <Td align="right">{acc.formatNumber(item.priceAddon, 2)}</Td>
    </tr>
  )
}

export default ListItem
