import React from "react"
import { useTranslation } from "react-i18next"
import { useMutation, queryCache } from "react-query"
import DeleteIcon from "@material-ui/icons/Delete"

import { AppBar } from "@Wrappers/layout"
import { Button, FileUploader } from "@Components/generic"
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
      <div className="overflow-y-scroll">
        <Wrapper generateData={handleGenerateData} />
        <FileUploader />
        <div className="m-8 p-8 border border-gray-500">
          <h1 className="text-2xl">Buttons</h1>
          <div className="m-4">
            <h2 className="text-xl">Contained</h2>
            <div className="flex space-x-4">
              <Button variant="contained" color="cool-gray">
                COOL GRAY
              </Button>
              <Button variant="contained" color="white">
                WHITE
              </Button>
              <Button variant="contained" color="red">
                RED
              </Button>
            </div>
          </div>
          <div className="m-4">
            <h2 className="text-xl">Outlined</h2>
            <div className="flex space-x-4">
              <Button variant="outlined" color="cool-gray">
                COOL GRAY
              </Button>
              <Button variant="outlined" color="white">
                WHITE
              </Button>
              <Button variant="outlined" color="red">
                RED
              </Button>
            </div>
          </div>
          <div className="m-4">
            <h2 className="text-xl">Text</h2>
            <div className="flex space-x-4">
              <Button variant="text" color="cool-gray">
                COOL GRAY
              </Button>
              <Button variant="text" color="white">
                WHITE
              </Button>
              <Button variant="text" color="red">
                RED
              </Button>
            </div>
          </div>
          <div className="m-4">
            <h2 className="text-xl">Icon Contained</h2>
            <div className="flex space-x-4">
              <Button variant="contained" color="cool-gray" circle>
                <DeleteIcon />
              </Button>
              <Button variant="contained" color="white" circle>
                <DeleteIcon />
              </Button>
              <Button variant="contained" color="red" circle>
                <DeleteIcon />
              </Button>
            </div>
          </div>
          <div className="m-4">
            <h2 className="text-xl">Icon Outlined</h2>
            <div className="flex space-x-4">
              <Button variant="outlined" color="cool-gray" circle>
                <DeleteIcon />
              </Button>
              <Button variant="outlined" color="white" circle>
                <DeleteIcon />
              </Button>
              <Button variant="outlined" color="red" circle>
                <DeleteIcon />
              </Button>
            </div>
          </div>
          <div className="m-4">
            <h2 className="text-xl">Icon Text</h2>
            <div className="flex space-x-4">
              <Button variant="text" color="cool-gray" circle>
                <DeleteIcon />
              </Button>
              <Button variant="text" color="white" circle>
                <DeleteIcon />
              </Button>
              <Button variant="text" color="red" circle>
                <DeleteIcon />
              </Button>
            </div>
          </div>

          <div className="m-4">
            <h2 className="text-xl">Sizes</h2>
            <div className="flex items-center space-x-4">
              <Button variant="contained" color="cool-gray" size="xs">
                EXTRA SMALL
              </Button>
              <Button variant="contained" color="cool-gray" size="sm">
                SMALL
              </Button>
              <Button variant="contained" color="cool-gray" size="md">
                MEDIUM
              </Button>
              <Button variant="contained" color="cool-gray" size="lg">
                LARGE
              </Button>
              <Button variant="contained" color="cool-gray" size="xl">
                EXTRA LARGE
              </Button>
              <Button variant="contained" color="cool-gray" size="xl" fullWidth>
                FULL WIDTH
              </Button>
            </div>
          </div>
          <div className="m-4">
            <h2 className="text-xl">Sizes Icon</h2>
            <div className="flex items-center space-x-4">
              <Button variant="contained" color="cool-gray" size="xs" circle>
                <DeleteIcon />
              </Button>
              <Button variant="contained" color="cool-gray" size="sm" circle>
                <DeleteIcon />
              </Button>
              <Button variant="contained" color="cool-gray" size="md" circle>
                <DeleteIcon />
              </Button>
              <Button variant="contained" color="cool-gray" size="lg" circle>
                <DeleteIcon />
              </Button>
              <Button variant="contained" color="cool-gray" size="xl" circle>
                <DeleteIcon />
              </Button>
            </div>
          </div>
        </div>
      </div>
    </>
  )
}

export default Dashboard
