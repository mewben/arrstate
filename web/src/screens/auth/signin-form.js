import React from "react"
import * as Yup from "yup"

import { Form, TextField, SubmitButton } from "@Components/forms"
import { t } from "@Utils/t"
import { signIn } from "@Store/actions/auth-actions"
import { extractError } from "@Utils"

const req = t("errors.required")

const validationSchema = Yup.object().shape({
  email: Yup.string().email(t("errors.email")).required(req),
  password: Yup.string().required(req),
})

const SigninForm = () => {
  const onSubmit = async data => {
    console.log("data", data)
    const res = await signIn(data)
    console.log("typeof res", typeof res, JSON.stringify(res, null, "  "))
    if (res.error) {
      console.log("res error", res)
      console.log("errorMesssage", res.data)
    } else {
      console.log("no error", res)
    }
  }

  return (
    <div>
      <Form
        onSubmit={onSubmit}
        validationSchema={validationSchema}
        model={{ email: "test@email.com", password: "password2" }}
      >
        <TextField
          name="email"
          label={t("email")}
          autoComplete="off"
          autoFocus
        />
        <TextField name="password" label={t("password")} type="password" />
        <SubmitButton>{t("signin")}</SubmitButton>
      </Form>
    </div>
  )
}

export default SigninForm
