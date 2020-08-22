import React from "react"
import FolderOpenIcon from "@material-ui/icons/FolderOpen"
import RecentActorsIcon from "@material-ui/icons/RecentActors"
import TextureIcon from "@material-ui/icons/Texture"
import { useTranslation } from "react-i18next"

import { useMeContext } from "@Wrappers"
import { Empty, Button } from "@Components/generic"

import Widget from "./widget"

const Wrapper = ({ generateData }) => {
  const { t } = useTranslation()
  const {
    currentBusiness: { dashboard },
  } = useMeContext()

  if (!dashboard) {
    return (
      <Empty>
        <div className="text-center">
          <h3 className="text-lg leading-6 font-medium text-gray-900">
            {t("dashboard.empty")}
          </h3>
          <div className="mt-2">
            <Button onClick={generateData}>{t("dashboard.btnGenerate")}</Button>
          </div>
        </div>
      </Empty>
    )
  }

  return (
    <div className="container p-4 mx-auto">
      <div className="grid grid-cols-12 gap-8">
        <div className="col-span-4">
          <Widget
            data={dashboard?.projects}
            link="/projects"
            icon={<FolderOpenIcon />}
            color="bg-green-400"
          />
        </div>
        <div className="col-span-4">
          <Widget
            data={dashboard?.properties}
            link="/properties"
            icon={<TextureIcon />}
            color="bg-orange-400"
          />
        </div>
        <div className="col-span-4">
          <Widget
            data={dashboard?.people}
            link="/people"
            icon={<RecentActorsIcon />}
            color="bg-blue-400"
          />
        </div>
      </div>
    </div>
  )
}

export default Wrapper
