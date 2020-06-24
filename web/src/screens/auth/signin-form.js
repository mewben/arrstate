import React from "react"
import { navigate } from "gatsby"
import * as Yup from "yup"
import { useMutation } from "react-query"

import { Form, TextField, SubmitButton } from "@Components/forms"
import { t } from "@Utils/t"
import { requestApi, extractError } from "@Utils"
import { useAuth } from "@Providers"

const req = t("errors.required")

const validationSchema = Yup.object().shape({
  email: Yup.string().email(t("errors.email")).required(req),
  password: Yup.string().required(req),
})

// ------- SigninForm -------- //
const SigninForm = () => {
  const { authSignIn } = useAuth()
  const [mutate, { reset, error, isError }] = useMutation(formData => {
    return requestApi("/auth/signin", "POST", {
      data: formData,
      noToken: true,
    })
  })

  const onSubmit = async formData => {
    reset()
    const res = await mutate(formData)
    if (res) {
      // store token,
      // redirect to '/'
      console.log("ressss", res.data)
      authSignIn(res.data.token)
      navigate("/", { replace: true })
    }
  }

  return (
    <div>
      <Form
        onSubmit={onSubmit}
        validationSchema={validationSchema}
        model={{ email: "test@email.com", password: "password" }}
      >
        {isError && <div>{extractError(error)}</div>}
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
