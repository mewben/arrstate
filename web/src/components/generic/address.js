import React from "react"
import { useTranslation } from "react-i18next"

export const Address = ({ address }) => {
  const { t } = useTranslation()

  return (
    <div className="address">
      {address?.street && <div>{address?.street}</div>}
      {address?.city && <div>{address?.city}</div>}
      {address?.zipCode && <div>{address?.zipCode}</div>}
      {address?.state && <div>{address?.state}</div>}
      {address?.country && <div>{t(`countries:${address?.country}`)}</div>}
    </div>
  )
}
