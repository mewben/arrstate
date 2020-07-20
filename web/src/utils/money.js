import { isNull } from "@Utils/lodash"

export const toMoney = input => {
  if (!isNull(input)) {
    return parseInt(input * 100)
  }
  return null
}

export const fromMoney = input => {
  if (!isNull(input)) {
    return input / 100
  }
  return null
}
