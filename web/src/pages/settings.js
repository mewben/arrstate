import React from "react"
import { Router, Redirect } from "@reach/router"
import { useTranslation } from "react-i18next"

import { PrivateWrapper } from "@Wrappers"
import { AppBar } from "@Wrappers/layout"
import { Account } from "@Screens/settings"

const SettingsPage = () => {
  const { t } = useTranslation()
  const submenu = [
    {
      label: t("settings.menu.account"),
      path: `/settings/account`,
    },
  ]

  return (
    <PrivateWrapper>
      <AppBar title={t("settings.title")} submenu={submenu} />
      <Router className="flex flex-col flex-1 overflow-hidden">
        <Account path="/settings/account" />
        <Redirect from="/settings" to="/settings/account" noThrow />
      </Router>
    </PrivateWrapper>
  )
}

export default SettingsPage
