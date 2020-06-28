import React from "react"
import { Link } from "gatsby"
import SettingsIcon from "@material-ui/icons/Settings"

const AppMenu = () => {
  return (
    <div className="flex flex-col justify-between items-center p-2 border-r border-gray-200">
      <Link to="/" className="group">
        <img
          src={`../../logo.png`}
          alt="app"
          className="h-8 w-8 rounded hover:bg-gray-100 p-0.5"
        />
      </Link>
      <Link
        to="/settings"
        className="rounded hover:bg-gray-100 h-8 w-8 text-gray-400 flex items-center justify-center"
      >
        <SettingsIcon />
      </Link>
    </div>
  )
}

export default AppMenu
