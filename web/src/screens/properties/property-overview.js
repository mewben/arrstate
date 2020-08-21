import React from "react"
import acc from "accounting"
import { useTranslation } from "react-i18next"

import { Panel, Portal, Button } from "@Components/generic"
import { AcquireForm } from "@Components/popups/property"
import { fullName } from "@Utils"
import { fromMoney } from "@Utils/money"
import { Status } from "./components"
import { PersonWrapper } from "@Wrappers"

const PropertyOverview = ({ property }) => {
  console.log("property", property)
  const { t } = useTranslation()

  return (
    <div className="grid grid-cols-6 gap-6">
      <div className="col-span-6 sm:col-span-3 p-4">
        <Panel noPadding>
          <div className="w-full flex items-start justify-between px-4 py-5 sm:p-6 space-x-6">
            <div className="flex-1 truncate">
              <div className="flex items-center space-x-3">
                <h3 className="text-gray-900 text-lg leading-5 font-medium truncate">
                  {property.name}
                </h3>
                <Status status={property.status} />
              </div>
              <p className="mt-1 text-gray-400 leading-5 truncate">
                <span className="uppercase text-xs tracking-widest">
                  {property.type}
                </span>
              </p>
            </div>
            <div>
              {property.status === "available" && (
                <div>
                  <Portal openByClickOn={<Button>{t("btnAcquire")}</Button>}>
                    <AcquireForm property={property} />
                  </Portal>
                </div>
              )}
              {!!property?.acquisition?.clientID && (
                <PersonWrapper personID={property?.acquisition?.clientID}>
                  {({ person }) => {
                    return (
                      <div>
                        <div className="text-gray-400 text-xs">
                          {t("client")}
                        </div>
                        <div className="mt-1 font-medium text-gray-900">
                          {fullName(person.givenName, person.familyName)}
                        </div>
                      </div>
                    )
                  }}
                </PersonWrapper>
              )}
              {!!property?.acquisition?.agentID && (
                <PersonWrapper personID={property?.acquisition?.agentID}>
                  {({ person }) => {
                    return (
                      <div className="mt-4">
                        <div className="text-gray-400 text-xs">
                          {t("agent")}
                        </div>
                        <div className="mt-1 font-medium text-gray-900">
                          {fullName(person.givenName, person.familyName)}
                        </div>
                      </div>
                    )
                  }}
                </PersonWrapper>
              )}
            </div>
          </div>
          <div className="flex p-6 space-x-6">
            <div>
              <div className="text-gray-400 text-xs">
                {t("properties.area")}
              </div>
              <div className="mt-1 font-medium text-gray-900">
                {acc.formatNumber(property.area, 2)} sq.m
              </div>
            </div>
            <div>
              <div className="text-green-300 text-xs">
                {t("properties.price")}
              </div>
              <div className="mt-1 font-medium text-green-500">
                Php {acc.formatNumber(fromMoney(property.price), 2)}
              </div>
            </div>
          </div>
        </Panel>
      </div>
    </div>
  )
}

export default PropertyOverview
