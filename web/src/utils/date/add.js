import {
  addSeconds,
  addMinutes,
  addHours,
  addDays,
  addMonths,
  addYears,
  addMilliseconds,
} from "date-fns"

/**
 * add time to date
 * @param date
 * @param amount {number}
 * @param  {('milliseconds'|'seconds'|'minutes'|'hours'|'days'|'months'|'years')} unit
 * @returns {number | *}
 */
export const add = (date, amount, unit = "days") => {
  switch (unit) {
    case "milliseconds":
      return addMilliseconds(date, amount)
    case "seconds":
      return addSeconds(date, amount)
    case "minutes":
      return addMinutes(date, amount)
    case "hours":
      return addHours(date, amount)
    case "days":
      return addDays(date, amount)
    case "months":
      return addMonths(date, amount)
    case "years":
      return addYears(date, amount)
    default:
  }
}
