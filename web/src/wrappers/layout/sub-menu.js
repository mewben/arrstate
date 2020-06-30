import React from "react"
import { Link } from "gatsby"

export const SubMenu = ({ children }) => {
  return <div className="flex overflow-x-scroll space-x-4">{children}</div>
}

export const SubMenuItem = ({ to, children }) => {
  return (
    <Link
      to={to}
      className="px-2 py-1 font-medium text-sm leading-5 rounded-md text-gray-500 hover:text-gray-700 focus:outline-none focus:text-gray-700 focus:bg-gray-100"
      activeClassName="text-gray-700 bg-gray-100"
    >
      {children}
    </Link>
  )
}
