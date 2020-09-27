import React from "react"

import { Td } from "@Components/generic"

const ListItem = ({ item }) => {
  return (
    <tr>
      <Td wrap>{item.title}</Td>
      <Td>{item.ext}</Td>
      <Td align="right">{item.size}</Td>
    </tr>
  )
}

export default ListItem
