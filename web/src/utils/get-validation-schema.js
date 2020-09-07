import { object, string, array } from "yup"
import { ERRORS } from "@Enums"

export const getValidationSchema = (t, entityType) => {
  const req = t(ERRORS.REQUIRED)

  switch (entityType) {
    case "signup":
      return object().shape({
        name: object().shape({
          first: string().required(req),
          last: string(),
        }),
        business: string().max(255).required(req),
        domain: string().max(255).required(req),
        email: string().email(t(ERRORS.EMAIL)).required(req),
        password: string()
          .min(6, t(ERRORS.MIN_LENGTH, { count: 6 }))
          .required(req),
      })

    case "person":
      return object().shape({
        name: object().shape({
          first: string().required(req),
          last: string(),
        }),
        email: string().email(t(ERRORS.EMAIL)).required(req),
        role: array(),
      })

    default:
      return null
  }
}
