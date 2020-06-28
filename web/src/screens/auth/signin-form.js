import React from "react"
import { Link, navigate } from "gatsby"
import * as Yup from "yup"
import { useMutation } from "react-query"

import { Form, TextField, SubmitButton } from "@Components/forms"
import { Error } from "@Components/generic"
import { t } from "@Utils/t"
import { requestApi } from "@Utils"
import { useAuth } from "@Providers"

const req = t("errors.required")

const validationSchema = Yup.object().shape({
  email: Yup.string().email(t("errors.email")).required(req),
  password: Yup.string().required(req),
})

// ------- SigninForm -------- //
const SigninForm = () => {
  const { authSignIn } = useAuth()
  const [mutate, { reset, error }] = useMutation(formData => {
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
    <div className="mt-6 w-full max-w-sm">
      <Form
        onSubmit={onSubmit}
        validationSchema={validationSchema}
        model={{ email: "melvinsoldia@gmail.com", password: "123456" }}
      >
        <Error error={error} />
        <div class="grid grid-cols-6 gap-6">
          <div className="col-span-6">
            <TextField
              name="email"
              label={t("email")}
              autoComplete="off"
              autoFocus
            />
          </div>
          <div className="col-span-6">
            <TextField name="password" label={t("password")} type="password" />
          </div>
          <div className="col-span-6">
            <SubmitButton size="xl" fullWidth>
              {t("Sign in")}
            </SubmitButton>
            <div className="text-xs text-cool-gray-500 mt-2">
              <Link
                to="/forgot-password"
                className="font-medium hover:text-gray-900"
              >
                Forgot your password?
              </Link>
            </div>
          </div>
        </div>
      </Form>
    </div>
  )
}

export default SigninForm
