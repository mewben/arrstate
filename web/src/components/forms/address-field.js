import React from "react"

import { t } from "@Utils/t"
import { TextField, InputGroup, CountriesSelectField } from "@Components/forms"

const AddressField = ({ name }) => {
  return (
    <>
      <InputGroup>
        <TextField
          name={`${name}.street`}
          placeholder={t("Street")}
          label={t("Street")}
          inputClassName="rounded-none rounded-l-md"
        />
        <TextField
          name={`${name}.city`}
          placeholder={t("City")}
          label={t("City")}
          inputClassName="rounded-none rounded-r-md"
        />
      </InputGroup>
      <InputGroup>
        <TextField
          name={`${name}.state`}
          placeholder={t("State")}
          label={t("State")}
          inputClassName="rounded-none rounded-l-md"
        />
        <TextField
          name={`${name}.zipCode`}
          placeholder={t("Post / Zip Code")}
          label={t("Post / Zip Code")}
          inputClassName="rounded-none rounded-r-md"
        />
      </InputGroup>
      <CountriesSelectField
        name={`${name}.country`}
        placeholder={t("Country")}
        label={t("Country")}
      />
    </>
  )
}

export default AddressField
