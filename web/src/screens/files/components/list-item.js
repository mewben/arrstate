import React from "react"
import filesize from "filesize"

import { Td } from "@Components/generic"
import { uploadURL } from "@Utils"

const ListItem = ({ item }) => {
  return (
    <tr>
      <Td wrap>
        <a
          className="font-medium text-cool-gray-800 hover:text-blue-500"
          href={uploadURL(item)}
          download
        >
          {item.title}
        </a>
      </Td>
      <Td>{item.ext}</Td>
      <Td align="right">{filesize(item.size)}</Td>
    </tr>
  )
}

export default ListItem
