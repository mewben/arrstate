import React from "react"
import { Router, Redirect } from "@reach/router"
import { useTranslation } from "react-i18next"

import { PrivateWrapper } from "@Wrappers"
import { AppBar, SubMenu, SubMenuItem } from "@Wrappers/layout"
import { Account } from "@Screens/settings"
import { map } from "@Utils/lodash"

const SettingsPage = () => {
  const { t } = useTranslation()
  const submenu = [
    {
      label: t("settings.menu.account"),
      path: `/settings/account`,
    },
  ]

  const renderSubmenu = () => {
    return (
      <SubMenu>
        {map(submenu, (item, i) => {
          return (
            <SubMenuItem key={i} to={item.path}>
              {item.label}
            </SubMenuItem>
          )
        })}
      </SubMenu>
    )
  }

  return (
    <PrivateWrapper>
      <AppBar title={t("settings.title")} submenu={renderSubmenu()} />
      <Router className="flex flex-col flex-1 overflow-hidden">
        <Account path="/settings/account" />
        <Redirect from="/settings" to="/settings/account" />
      </Router>
    </PrivateWrapper>
  )
}

export default SettingsPage
