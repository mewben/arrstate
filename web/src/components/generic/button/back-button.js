import React from "react"
import { Link } from "@reach/router"
import ChevronLeftIcon from "@material-ui/icons/ChevronLeft"

const BackButton = ({ to }) => {
  if (!to) return null

  return (
    <div className="border-r border-gray-200 h-full flex items-center pr-4 mr-4">
      <Link
        to={to}
        className="text-xs font-medium text-gray-400 hover:text-gray-800"
      >
        <ChevronLeftIcon fontSize="medium" />
      </Link>
    </div>
  )
}

export default BackButton
