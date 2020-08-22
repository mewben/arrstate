import React from "react"
import { useTranslation } from "react-i18next"

const Empty = ({ children }) => {
  const { t } = useTranslation()
  const content = children || t("empty.default")
  return (
    <div className="flex w-full h-full items-center justify-center">
      {content}
    </div>
  )
}

export default Empty
