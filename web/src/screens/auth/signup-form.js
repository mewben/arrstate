import React from "react"
import * as Yup from "yup"

import { Form, TextField, SubmitButton } from "@Components/forms"
import { t } from "@Utils/t"

const req = t("errors.required")

const validationSchema = Yup.object().shape({
  givenName: Yup.string().required(req),
  familyName: Yup.string(),
  business: Yup.string().max(255).required(req),
  domain: Yup.string().max(255).required(req),
  email: Yup.string().email(t("errors.email")).required(t("errors.required")),
  password: Yup.string()
    .min(6, t("errors.minLength", { count: 6 }))
    .required(t("errors.required")),
})

const SignupForm = () => {
  const onSubmit = data => {
    console.log("data", data)
  }

  return (
    <div>
      <Form onSubmit={onSubmit} validationSchema={validationSchema}>
        <TextField
          name="givenName"
          label={t("givenName")}
          autoComplete="off"
          autoFocus
        />
        <TextField name="familyName" label={t("familyName")} />
        <TextField name="business" label={t("business")} />
        <TextField name="domain" label={t("domain")} />
        <TextField name="email" label={t("email")} />
        <TextField name="password" type="password" label={t("password")} />
        <SubmitButton>{t("signup")}</SubmitButton>
      </Form>
    </div>
  )
}

export default SignupForm
