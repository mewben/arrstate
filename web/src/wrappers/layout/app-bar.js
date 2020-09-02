import React from "react"

import { BackButton } from "@Components/generic/button"
import { map, isEmpty } from "@Utils/lodash"
import { SubMenu, SubMenuItem } from "./sub-menu"

export const AppBar = ({
  title,
  backTo,
  submenu,
  submenuRenderer,
  children,
}) => {
  const renderSubmenu = () => {
    if (submenuRenderer) return submenuRenderer()
    if (isEmpty(submenu)) return null

    return (
      <div className="flex items-center px-4 h-10 border-t border-gray-200">
        <SubMenu>
          {map(submenu, (item, i) => {
            return (
              <SubMenuItem key={i} to={item.path}>
                {item.label}
              </SubMenuItem>
            )
          })}
        </SubMenu>
      </div>
    )
  }

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
      {renderSubmenu()}
    </div>
  )
}
