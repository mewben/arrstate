import React from "react"
import { Controller } from "react-hook-form"
import { DateRangePicker } from "@material-ui/pickers"
import { useTranslation } from "react-i18next"

const DateRangePickerField = ({ name }) => {
  const { t } = useTranslation()

  return (
    <div>
      <Controller
        name={name}
        render={({ onChange, value }) => {
          return (
            <DateRangePicker
              value={value}
              onChange={onChange}
              // disableCloseOnSelect={false}
              // showTodayButton={true}
              reduceAnimations
              // onAccept={onChange}
              allowSameDateSelection
              renderInput={(startProps, endProps) => {
                return (
                  <div className="flex items-center space-x-2">
                    <input
                      className="form-input"
                      ref={startProps.inputRef}
                      {...startProps.inputProps}
                      placeholder="MM/DD/YYYY"
                    />
                    <span>{t("to")}</span>
                    <input
                      className="form-input"
                      ref={endProps.inputRef}
                      {...endProps.inputProps}
                      placeholder="MM/DD/YYYY"
                    />
                  </div>
                )
              }}
            />
          )
        }}
      />
    </div>
  )
}

export default DateRangePickerField
