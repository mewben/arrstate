import React from "react"
import { Link } from "gatsby"

import { Td } from "@Components/generic"
import { map } from "@Utils/lodash"
import { t } from "@Utils/t"

const ListItem = ({ item }) => {
  return (
    <tr>
      <Td wrap>
        <Link
          to={`/people/${item._id}`}
          className="font-medium text-cool-gray-700 hover:text-blue-500"
        >
          {item.givenName}
        </Link>
      </Td>
      <Td>
        <div className="flex space-x-2">
          {map(item.role, rol => (
            <span key={rol}>{t(rol)}</span>
          ))}
        </div>
      </Td>
    </tr>
  )
}

export default ListItem
