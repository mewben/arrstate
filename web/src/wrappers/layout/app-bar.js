import React from "react"

import { BackButton } from "@Components/generic/button"

export const AppBar = ({ title, backTo, submenu, children }) => {
  return (
    <div className="flex flex-col flex-shrink-0 bg-white shadow-sm relative z-10 border-b border-gray-200">
      <div className="flex px-4 justify-between items-center relative z-10 h-12">
        <div className="flex h-full items-center">
          <BackButton to={backTo} />
          <h1 className="text-base font-medium leading-6 text-cool-gray-900">
            {title}
          </h1>
        </div>
        {children}
      </div>
      {!!submenu && (
        <div className="flex items-center px-4 h-10 border-t border-gray-200">
          {submenu}
        </div>
      )}
    </div>
  )
}
