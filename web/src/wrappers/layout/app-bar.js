import React from "react"
import { Link } from "gatsby"

import { BackButton } from "@Components/generic/button"
import { Breadcrumbs, Breadcrumb } from "@Components/generic"
import { SubMenu, SubMenuItem } from "@Wrappers/layout"

export const AppBar = ({ title, backTo, submenu, children }) => {
  // const renderSubBar = () => {
  //   if (!nav && !submenu) return null
  //   return (
  //     <div className="flex items-center px-4 h-10 border-t border-gray-200">
  //       {!!nav && (
  //         <div className="border-r border-gray-200 h-full flex items-center pr-4 mr-4">
  //           {/* {nav} */}
  //           <Breadcrumbs>
  //             <Breadcrumb to="/">Projects</Breadcrumb>
  //             <Breadcrumb to="/">Lots</Breadcrumb>
  //             <Breadcrumb to="/">Projects</Breadcrumb>
  //           </Breadcrumbs>
  //         </div>
  //       )}
  //       {submenu}
  //       <SubMenu>
  //         <SubMenuItem to="/">Overview</SubMenuItem>
  //       </SubMenu>
  //     </div>
  //   )
  // }

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
