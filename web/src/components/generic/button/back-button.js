import React from "react"
import { navigate } from "@reach/router"
import ChevronLeftIcon from "@material-ui/icons/ChevronLeft"

const BackButton = ({ to }) => {
  if (!to) return null

  return (
    <div className="border-r border-gray-200 h-full flex items-center pr-4 mr-4">
      <button
        // to={to}
        onClick={() => navigate(-1)}
        className="text-xs font-medium text-gray-400 hover:text-gray-800"
      >
        <ChevronLeftIcon fontSize="default" />
      </button>
    </div>
  )
}

export default BackButton
