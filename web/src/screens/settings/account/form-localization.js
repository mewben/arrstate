import React from "react"
import { useTranslation } from "react-i18next"
import { formatDate, requestApi } from "@Utils"
import { useMutation, queryCache } from "react-query"

import {
  Form,
  SelectField,
  LanguageSelectField,
  SubmitButton,
} from "@Components/forms"
import { useCurrentContext } from "@Wrappers"
import { Time } from "@Components/generic"
import { DATE_FORMATS } from "@Enums/date-formats"
import { TIME_FORMATS, TIMESTAMP_FORMATS } from "@Enums/time-formats"
import { map } from "@Utils/lodash"

const FormLocalization = ({ model }) => {
  const { t } = useTranslation()
  const {
    currentPerson: { _id: personID },
  } = useCurrentContext()

  const [save, { reset, error }] = useMutation(
    formData => {
      return requestApi(`/api/people/${personID}/locale`, "PUT", {
        data: formData,
      })
    },
    {
      onSuccess: ({ data }) => {
        queryCache.setQueryData("currentPerson", data)
        queryCache.invalidateQueries("currentPerson")
      },
    }
  )

  const {
    dateFormatOptions,
    timeFormatOptions,
    timestampFormatOptions,
  } = React.useMemo(() => {
    const dateFormatOptions = map(DATE_FORMATS, df => ({
      value: df,
      label: formatDate(undefined, df),
    }))

    const timeFormatOptions = map(TIME_FORMATS, tf => ({
      value: tf,
      label: formatDate(undefined, tf),
    }))

    const timestampFormatOptions = map(TIMESTAMP_FORMATS, tsf => ({
      value: tsf,
      labelText: t(`timestampFormats.${tsf}`),
      label: (
        <>
          {t(`timestampFormats.${tsf}`)} (<Time format={tsf} d={new Date()} />)
        </>
      ),
    }))

    return { dateFormatOptions, timeFormatOptions, timestampFormatOptions }
  }, [])

  const onSubmit = async formData => {
    const payload = {
      ...model,
      ...formData,
    }
    reset()
    await save(payload)
  }

  return (
    <div>
      <div className="px-6 pt-6">
        <h3 className="text-lg leading-6 font-medium text-gray-900">
          {t("account.localization.title")}
        </h3>
      </div>
      <Form model={model} onSubmit={onSubmit}>
        <div className="grid grid-cols-12 gap-x-6 gap-y-6 p-6">
          <LanguageSelectField
            name="language"
            leftLabel={t("form.localization.language")}
            disableClearable
          />
          <SelectField
            name="dateFormat"
            leftLabel={t("form.localization.dateFormat")}
            options={dateFormatOptions}
            disableClearable
          />
          <SelectField
            name="timeFormat"
            leftLabel={t("form.localization.timeFormat")}
            options={timeFormatOptions}
            disableClearable
          />
          {/* <SelectField
            name="timestampFormat"
            leftLabel={t("form.localization.timestampFormat")}
            options={timestampFormatOptions}
            disableClearable
          /> */}
          <div>
            <SubmitButton>{t("btnSave")}</SubmitButton>
          </div>
        </div>
      </Form>
    </div>
  )
}

export default FormLocalization
