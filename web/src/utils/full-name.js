import { compact } from "@Utils/lodash"

export const fullName = (givenName, familyName) => {
  return compact([givenName, familyName]).join(" ")
}
