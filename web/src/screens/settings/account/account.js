import React from "react"
import { useTranslation } from "react-i18next"

import { Panel, Button } from "@Components/generic"
import FormLocalization from "./form-localization"
import { useCurrentContext } from "@Wrappers"

const AccountSettings = () => {
  const { t } = useTranslation()
  const { currentPerson } = useCurrentContext()

  return (
    <div className="overflow-y-scroll pb-28">
      <div className="container mx-auto max-w-2xl mt-8">
        <div className="flex flex-col space-y-8">
          <Panel noPadding>
            <FormLocalization model={currentPerson?.locale} />
          </Panel>
          <Panel>
            <Button color="red" to="/signout">
              {t("btnSignout")}
            </Button>
          </Panel>
        </div>
      </div>
    </div>
  )
}

export default AccountSettings
