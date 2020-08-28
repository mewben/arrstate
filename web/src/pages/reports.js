import React from "react"
import { Router, Redirect } from "@reach/router"
import { useTranslation } from "react-i18next"

import { PrivateWrapper } from "@Wrappers"
import { AppBar } from "@Wrappers/layout"
import { Income } from "@Screens/reports"

const ReportsPage = () => {
  const { t } = useTranslation()
  const submenu = [
    {
      label: t("reports.menu.income"),
      path: `/reports/income`,
    },
  ]
  return (
    <PrivateWrapper>
      <AppBar title={t("reports.title")} submenu={submenu} />
      <Router className="flex flex-col flex-1 overflow-hidden">
        <Income path="/reports/income" />
        <Redirect from="/reports" to="/reports/income" />
      </Router>
    </PrivateWrapper>
  )
}

export default ReportsPage
