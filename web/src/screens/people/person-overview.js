import React from "react"
import { useTranslation } from "react-i18next"

import { Panel, PersonBadge, Address } from "@Components/generic"
import { map } from "@Utils/lodash"
import { fullName, fullAddress } from "@Utils"
import { Dtd } from "./components"

const PersonOverview = ({ person }) => {
  const { t } = useTranslation()

  return (
    <div className="container mx-auto max-w-2xl mt-8">
      <Panel noPadding>
        <div className="bg-cool-gray-700 h-24 sm:h-20 lg:h-28"></div>
        <div className="divide-y divide-cool-gray-200 pb-6">
          <div className="-mt-12 flow-root px-4 space-y-6 sm:-mt-8 sm:flex sm:items-end sm:px-6 sm:space-x-6 lg:-mt-15 pb-6">
            <div className="flex items-end space-x-8">
              <div className="-m-1 flex">
                <PersonBadge person={person} size="2xl" />
              </div>
              <div className="mb-8">
                <h3 className="font-bold text-xl leading-7 text-cool-gray-900 sm:text-2xl sm:leading-8">
                  {fullName(person.name)}
                </h3>
                <p className="text-sm leading-5 text-cool-gray-500">
                  <div className="flex space-x-2">
                    {map(person.role, rol => (
                      <span key={rol}>{rol}</span>
                    ))}
                  </div>
                </p>
              </div>
            </div>
          </div>
          <Dtd label={t("email")} value={person.email} />
          <Dtd
            label={t("people.address")}
            value={<Address address={person.address} />}
          />
        </div>
      </Panel>
    </div>
  )
}

export default PersonOverview
