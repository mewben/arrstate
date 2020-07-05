import React from "react"
import { Link } from "gatsby"

import { Td } from "@Components/generic"

const ListItem = ({ item }) => {
  return (
    <tr>
      <Td>
        <Link
          to={`/people/${item._id}`}
          className="font-medium text-cool-gray-700 hover:text-blue-500"
        >
          {item.givenName}
        </Link>
      </Td>
    </tr>
  )
}

export default ListItem
