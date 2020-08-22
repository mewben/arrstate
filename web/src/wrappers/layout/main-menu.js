import React from "react"
import { Link } from "gatsby"
import FolderOpenIcon from "@material-ui/icons/FolderOpen"
import TextureIcon from "@material-ui/icons/Texture"
import RecentActorsIcon from "@material-ui/icons/RecentActors"
import HomeIcon from "@material-ui/icons/Home"
import { useTranslation } from "react-i18next"

import { map } from "@Utils/lodash"

const MainMenu = () => {
  const { t } = useTranslation()

  const mainmenu = [
    {
      path: "/",
      label: t("menu.dashboard"),
      icon: <HomeIcon />,
      partiallyActive: false,
    },
    {
      path: "/projects",
      label: t("menu.projects"),
      icon: <FolderOpenIcon />,
    },
    {
      path: "/properties",
      label: t("menu.properties"),
      icon: <TextureIcon />,
    },
    {
      path: "/people",
      label: t("menu.people"),
      icon: <RecentActorsIcon />,
    },
  ]

  return (
    <div className="flex-1 px-2">
      {map(mainmenu, menu => {
        return (
          <Link
            key={menu.path}
            to={menu.path}
            className="mt-1 group flex items-center px-2 py-2 text-sm leading-5 font-medium text-gray-600 rounded-md hover:text-gray-900 hover:bg-gray-50 focus:outline-none focus:bg-gray-100 transition ease-in-out duration-150"
            activeClassName="text-gray-900 bg-gray-200"
            partiallyActive={menu.partiallyActive}
          >
            <span className="mr-3 text-gray-400 group-hover:text-gray-500 group-focus:text-gray-500 transition ease-in-out duration-150">
              {menu.icon}
            </span>
            <span>{menu.label}</span>
          </Link>
        )
      })}
    </div>
  )
}

export default MainMenu
