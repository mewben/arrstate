import React from "react"
import CloseIcon from "@material-ui/icons/CloseOutlined"

export const DrawerHeader = ({ title, children, onClose }) => {
  return (
    <header className="space-y-1 py-6 px-4 bg-cool-gray-700 sm:px-6">
      <div className="flex items-center justify-between space-x-3">
        <h2 className="text-base leading-7 font-medium text-white">{title}</h2>
        <button
          onClick={onClose}
          type="button"
          className="text-gray-200 hover:text-white transition ease-in-out duration-150"
        >
          <CloseIcon />
        </button>
      </div>
      {!!children && (
        <div className="text-sm leading-5 text-gray-300">{children}</div>
      )}
    </header>
  )
}
