import React from "react"
import filesize from "filesize"

import { Td } from "@Components/generic"

const ListItem = ({ item }) => {
  return (
    <tr>
      <Td wrap>{item.title}</Td>
      <Td>{item.ext}</Td>
      <Td align="right">{filesize(item.size)}</Td>
    </tr>
  )
}

export default ListItem
