import React from "react"
import { useTranslation } from "react-i18next"
import cx from "clsx"

import { Button, Panel } from "@Components/generic"

const Widget = ({ data, label, icon, color, link, linkTitle }) => {
  const { t } = useTranslation()
  if (!linkTitle) {
    linkTitle = t("btnViewAll")
  }

  return (
    <Panel noPadding>
      <div className="p-5">
        <div className="flex items-center">
          <div className="flex-shrink-0">
            <div className={cx("text-cool-gray-400", color)}>{icon}</div>
          </div>
          <div className="ml-5 w-0 flex-1">
            <dl>
              <dt className="text-sm leading-5 font-medium text-cool-gray-500 truncate">
                {label}
              </dt>
              <dd>
                <div className="text-lg leading-7 font-medium text-cool-gray-900">
                  {data}
                </div>
              </dd>
            </dl>
          </div>
        </div>
      </div>
      <div className="bg-cool-gray-50 px-5 py-2">
        <Button to={link} size="sm" variant="text">
          {linkTitle}
        </Button>
      </div>
    </Panel>
  )
}

export default Widget
