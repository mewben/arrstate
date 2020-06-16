import React from "react"
import * as Yup from "yup"

import { Form, TextField } from "@Components/forms"
import { t } from "@Utils/t"

const req = t("errors.required")

const validationSchema = Yup.object().shape({
  email: Yup.string().email(t("errors.email")).required(req),
  password: Yup.string().required(req),
})

const SigninForm = () => {
  const onSubmit = data => {
    console.log("data", data)
  }

  return (
    <div>
      <Form onSubmit={onSubmit} validationSchema={validationSchema}>
        <TextField
          name="email"
          label={t("email")}
          autoComplete="off"
          autoFocus
        />
        <TextField name="password" label={t("password")} type="password" />
        <button type="submit">{t("signin")}</button>
      </Form>
    </div>
  )
}

export default SigninForm
