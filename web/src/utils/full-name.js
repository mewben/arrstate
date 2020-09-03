import { compact } from "@Utils/lodash"

export const fullName = name => {
  return compact([name?.first, name?.last]).join(" ")
}
