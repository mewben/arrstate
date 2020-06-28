import React from "react"

export const AppBar = ({ title, children }) => {
  return (
    <div className="flex px-4 justify-between items-center relative z-10 flex-shrink-0 h-14 bg-white shadow-sm">
      <h1 className="text-base font-medium leading-6 text-cool-gray-900">
        {title}
      </h1>
      {children}
    </div>
  )
}
