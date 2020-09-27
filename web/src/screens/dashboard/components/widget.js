import React from "react"
import { Link } from "gatsby"
import { useTranslation } from "react-i18next"
import acc from "accounting"
import cx from "clsx"

import { Button } from "@Components/generic/button"

const Widget = ({ data, icon, color, link }) => {
  const { t } = useTranslation()

  if (!data) {
    return null
  }

  return (
    <div className="bg-white overflow-hidden shadow rounded-lg">
      <div className="px-4 py-5 sm:p-6">
        <div className="flex items-center">
          <div className={cx("flex-shrink-0 rounded-md text-white p-3", color)}>
            {icon}
          </div>
          <div className="ml-5 w-0 flex-1">
            <dl>
              <dt className="text-sm leading-5 font-medium text-gray-500 truncate">
                {data.label}
              </dt>
              <dd className="flex items-baseline">
                <div className="text-2xl leading-8 font-semibold text-gray-900">
                  {acc.formatNumber(data.total)}
                </div>
              </dd>
            </dl>
          </div>
        </div>
      </div>
      <div className="bg-gray-50 px-4 py-4 sm:px-6">
        <div className="text-sm leading-5">
          <Button to={link} size="xs">
            {t("btnViewAll")}
          </Button>
        </div>
      </div>
    </div>
  )
}

export default Widget
