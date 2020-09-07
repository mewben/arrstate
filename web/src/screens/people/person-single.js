import React from "react"
import { Router } from "@reach/router"
import { useQuery } from "react-query"
import { useTranslation } from "react-i18next"

import { AppBar } from "@Wrappers/layout"
import { Loading, Error, Portal, Button } from "@Components/generic"
import { PersonForm } from "@Components/popups/person"
import { fullName } from "@Utils"
import { PersonWrapper } from "@Wrappers"
import PersonOverview from "./person-overview"

const PersonSingle = ({ personID }) => {
  const { t } = useTranslation()
  return (
    <PersonWrapper personID={personID}>
      {({ person }) => {
        return (
          <>
            <AppBar title={fullName(person?.name)} backTo>
              <Portal openByClickOn={<Button>{t("people.edit")}</Button>}>
                <PersonForm model={person} />
              </Portal>
            </AppBar>
            <Router className="flex-1 overflow-y-scroll pb-28">
              <PersonOverview path="/" person={person} />
            </Router>
          </>
        )
      }}
    </PersonWrapper>
  )
}

export default PersonSingle
