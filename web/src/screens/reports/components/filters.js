import React from "react"
import { useTranslation } from "react-i18next"

import { Form, DateRangePickerField, SubmitButton } from "@Components/forms"

const Filters = ({ isLoading, onSubmit }) => {
  const { t } = useTranslation()

  return (
    <div className="bg-cool-gray-200 py-2 px-4 border-b border-cool-gray-300">
      <Form model={{ range: [null, null] }} onSubmit={onSubmit}>
        <div className="flex items-center space-x-4">
          <DateRangePickerField name="range" />
          <SubmitButton disabled={isLoading}>{t("btnGenerate")}</SubmitButton>
        </div>
      </Form>
    </div>
  )
}

export default Filters
