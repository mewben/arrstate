import React from "react"
import FolderOpenIcon from "@material-ui/icons/FolderOpen"
import RecentActorsIcon from "@material-ui/icons/RecentActors"
import TextureIcon from "@material-ui/icons/Texture"
import { useTranslation } from "react-i18next"
import useSwr from "swr"
import MoneyIcon from "@material-ui/icons/Payment"
import { fromMoney } from "@Utils/money"
import acc from "accounting"

import { Loading } from "@Components/generic"
import { useParams } from "@Utils"
import Widget from "./widget"
import SalesWidget from "./sales-widget"

const Wrapper = () => {
  const { t } = useTranslation()
  const params = useParams()

  const { data, error } = useSwr(["/api/dashboard", params.toString()])

  if (!data && !error) {
    return <Loading />
  }

  const counts = data?.counts || {}

  return (
    <div className="container p-4 mx-auto">
      <div className="grid grid-cols-12 gap-x-6 gap-y-6">
        <div className="col-span-3">
          <Widget
            data={`Php ${acc.formatNumber(fromMoney(counts.sales), 2)}`}
            label={t("salesTotal")}
            icon={<MoneyIcon fontSize="large" />}
            color="text-green-400"
            link="/reports/income"
            linkTitle={t("btnViewReports")}
          />
        </div>
        <div className="col-span-3">
          <Widget
            data={acc.formatNumber(counts.projects)}
            label={t("projects.title")}
            link="/projects"
            icon={<FolderOpenIcon fontSize="large" />}
            color="text-yellow-200"
          />
        </div>
        <div className="col-span-3">
          <Widget
            data={acc.formatNumber(counts.properties)}
            label={t("properties.title")}
            link="/properties"
            icon={<TextureIcon fontSize="large" />}
            color="text-orange-400"
          />
        </div>
        <div className="col-span-3">
          <Widget
            data={acc.formatNumber(counts.people)}
            label={t("people.title")}
            link="/people"
            icon={<RecentActorsIcon fontSize="large" />}
            color="text-blue-400"
          />
        </div>
        <SalesWidget data={data?.sales} />
      </div>
    </div>
  )
}

export default Wrapper
