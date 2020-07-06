import React from "react"
import { Link, navigate } from "gatsby"
import * as Yup from "yup"
import { useMutation } from "react-query"

import {
  Form,
  TextField,
  BaseTextField,
  SubmitButton,
  InputGroup,
  FieldError,
} from "@Components/forms"
import { Error } from "@Components/generic"
import { t } from "@Utils/t"
import { requestApi } from "@Utils"
import { useAuth } from "@Providers"

const req = t("errors.required")

const validationSchema = Yup.object().shape({
  givenName: Yup.string().required(req),
  familyName: Yup.string(),
  business: Yup.string().max(255).required(req),
  domain: Yup.string().max(255).required(req),
  email: Yup.string().email(t("errors.email")).required(req),
  password: Yup.string()
    .min(6, t("errors.minLength", { count: 6 }))
    .required(req),
})

// ------- SignupForm -------- //
const SignupForm = () => {
  const { authSignIn } = useAuth()
  const [mutate, { reset, error }] = useMutation(formData => {
    return requestApi("/auth/signup", "POST", {
      data: formData,
      noToken: true,
    })
  })

  const onSubmit = async formData => {
    reset()
    const res = await mutate(formData)
    if (res) {
      const params = new URLSearchParams()
      params.set("deviceCode", res.data.deviceCode)
      window.location.assign(
        `${process.env.GATSBY_HTTP_PROTOCOL}${res.data.domain}.${
          process.env.GATSBY_DOMAIN
        }/cb?${params.toString()}`
      )
    }
  }

  return (
    <div className="mt-6 w-full max-w-sm">
      <Form onSubmit={onSubmit} validationSchema={validationSchema}>
        <Error error={error} />
        <div className="grid grid-cols-6 gap-6">
          <div className="col-span-6">
            <InputGroup label={t("name")}>
              <BaseTextField
                name="givenName"
                className="rounded-none rounded-l-md"
                placeholder={t("name.givenName")}
                autoFocus
              />
              <BaseTextField
                name="familyName"
                className="rounded-none rounded-r-md"
                placeholder={t("name.familyName")}
              />
            </InputGroup>
            <FieldError name="givenName" />
          </div>
          <div className="col-span-6">
            <TextField name="business" label={t("business")} />
          </div>
          <div className="col-span-6">
            <TextField
              name="domain"
              label={t("domain")}
              description={t("You can change the name and domain anytime.")}
              endAddon=".realtydomain.com"
            />
          </div>
          <div className="col-span-6">
            <TextField name="email" label={t("email")} />
          </div>
          <div className="col-span-6">
            <TextField name="password" type="password" label={t("password")} />
          </div>
          <div className="col-span-6">
            <SubmitButton size="xl" fullWidth>
              {t("Sign up")}
            </SubmitButton>
            <div className="text-xs text-cool-gray-500 mt-2">
              By creating an account, you are agreeing to our{" "}
              <Link to="/terms" className="font-medium hover:text-gray-900">
                Terms of Service
              </Link>
            </div>
          </div>
        </div>
      </Form>
    </div>
  )
}

export default SignupForm
