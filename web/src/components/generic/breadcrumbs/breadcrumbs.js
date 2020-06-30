import React from "react"
import { Link } from "gatsby"
import MuiBreadcrumbs from "@material-ui/core/Breadcrumbs"

const Breadcrumbs = ({
  separator = <span className="text-gray-300">/</span>,
  children,
}) => {
  return (
    <MuiBreadcrumbs separator={separator} aria-label="breadcrumb">
      {children}
    </MuiBreadcrumbs>
  )
}

export const Breadcrumb = ({ to, children }) => {
  return (
    <Link
      to={to}
      className="text-xs font-medium text-gray-500 hover:text-gray-800"
    >
      {children}
    </Link>
  )
}

export default Breadcrumbs
