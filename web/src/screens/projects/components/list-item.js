import React from "react"
import { Link } from "gatsby"
import acc from "accounting"

const ListItem = ({ item }) => {
  return (
    <div>
      <div>
        <Link to={`/projects/${item._id}/lots`}>{item.name}</Link>
      </div>
      <div>
        {acc.formatNumber(item.area, 2)} {item.unit}
      </div>
    </div>
  )
}

export default ListItem
