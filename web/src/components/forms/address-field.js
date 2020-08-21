import React from "react"
import { useTranslation } from "react-i18next"

import { TextField, InputGroup, CountriesSelectField } from "@Components/forms"

const AddressField = ({ name }) => {
  const { t } = useTranslation()
  return (
    <>
      <InputGroup>
        <TextField
          name={`${name}.street`}
          placeholder={t("form.address.streetP")}
          label={t("form.address.street")}
          inputClassName="rounded-none rounded-l-md"
        />
        <TextField
          name={`${name}.city`}
          placeholder={t("form.address.cityP")}
          label={t("form.address.city")}
          inputClassName="rounded-none rounded-r-md"
        />
      </InputGroup>
      <InputGroup>
        <TextField
          name={`${name}.state`}
          placeholder={t("form.address.stateP")}
          label={t("form.address.state")}
          inputClassName="rounded-none rounded-l-md"
        />
        <TextField
          name={`${name}.zipCode`}
          placeholder={t("form.address.zipP")}
          label={t("form.address.zip")}
          inputClassName="rounded-none rounded-r-md"
        />
      </InputGroup>
      <CountriesSelectField
        name={`${name}.country`}
        placeholder={t("form.address.countryP")}
        label={t("form.address.country")}
      />
    </>
  )
}

export default AddressField
