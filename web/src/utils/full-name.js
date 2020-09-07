import { compact } from "@Utils/lodash"

export const fullName = name => {
  return compact([name?.first, name?.last]).join(" ")
}

export const initialsName = name => {
  return name?.first?.charAt(0) + name?.last?.charAt(0) || "ME"
}

export const fullAddress = (address, t) => {
  return compact([
    address?.street,
    address?.city,
    address?.state,
    address?.zipCode,
    address?.country && t(`countries:${address?.country}`),
  ]).join(", ")
}
