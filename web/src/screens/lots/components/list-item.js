import React from "react"
import { Link } from "gatsby"

// Lots listitem
const ListItem = ({ item }) => {
  return (
    <div>
      <div>
        <Link to={`/lots/${item._id}`}>{item.name}</Link>
      </div>
    </div>
  )
}

export default ListItem
