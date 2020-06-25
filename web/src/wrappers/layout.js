import React from "react"
import { Link } from "gatsby"

import { map } from "@Utils/lodash"

export const LayoutWrapper = ({ children }) => {
  return (
    <div>
      <MainMenu />
      {children}
    </div>
  )
}

const MainMenu = () => {
  const mainmenu = [
    {
      path: "/",
      label: "Home",
    },
    {
      path: "/projects",
      label: "Projects",
    },
    {
      path: "/lots",
      label: "Lots",
    },
    {
      path: "/clients",
      label: "Clients",
    },
    {
      path: "/agents",
      label: "Agents",
    },
  ]

  return (
    <nav>
      {map(mainmenu, menu => {
        return <Link to={menu.path}>{menu.label}</Link>
      })}
    </nav>
  )
}
