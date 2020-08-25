import React from "react"
import { useTranslation } from "react-i18next"
import { useMutation, queryCache } from "react-query"

import { AppBar } from "@Wrappers/layout"
import { Button } from "@Components/generic"
import { requestApi } from "@Utils"
import { Wrapper } from "./components"

const Dashboard = () => {
  const { t } = useTranslation()

  const [fetchData, { reset, error }] = useMutation(
    () => {
      return requestApi("/api/dashboard", "GET")
    },
    {
      onSuccess: ({ data }) => {
        queryCache.setQueryData("currentBusiness", data)
        queryCache.invalidateQueries("currentBusiness")
      },
    }
  )

  const handleGenerateData = () => {
    reset()
    fetchData()
  }

  // TODO: error notification notistack

  return (
    <>
      <AppBar title={t("dashboard.title")}>
        <Button onClick={handleGenerateData} color="white">
          {t("dashboard.btnRefresh")}
        </Button>
      </AppBar>
      <Wrapper generateData={handleGenerateData} />
    </>
  )
}

export default Dashboard
