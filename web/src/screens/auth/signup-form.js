import React from "react"
import { Link, navigate } from "gatsby"
import { useMutation } from "react-query"
import { useTranslation } from "react-i18next"

import {
  Form,
  TextField,
  BaseTextField,
  SubmitButton,
  InputGroup,
  FieldError,
} from "@Components/forms"
import { Error } from "@Components/generic"
import { requestApi, getValidationSchema } from "@Utils"

// ------- SignupForm -------- //
const SignupForm = () => {
  const { t } = useTranslation()
  const validationSchema = React.useMemo(() => {
    return getValidationSchema(t, "signup")
  }, [t])

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
        <div className="grid grid-cols-12 gap-6">
          <Error error={error} className="col-span-12" />
          <InputGroup
            name="name.first"
            id="givenName"
            label={t("name.fullName")}
          >
            <BaseTextField
              name="name.first"
              id="givenName"
              inputClassName="rounded-none rounded-l-md"
              placeholder={t("name.givenName")}
              autoFocus
            />
            <BaseTextField
              name="name.last"
              inputClassName="rounded-none rounded-r-md"
              placeholder={t("name.familyName")}
            />
          </InputGroup>
          <TextField name="business" label={t("business.name")} />
          <TextField name="name.first" label={t("fsldkjfdk")} />
          <TextField
            name="domain"
            label={t("domain")}
            description={t("signup.domainHelp")}
            endAddon=".arrstate.com"
          />
          <TextField name="email" type="email" label={t("email")} />
          <TextField name="password" type="password" label={t("password")} />
          <div className="col-span-12">
            <SubmitButton size="xl" fullWidth>
              {t("signup.btn")}
            </SubmitButton>
            <div className="text-xs text-cool-gray-500 mt-2">
              {t("signup.termsHelp")}{" "}
              <Link to="/terms" className="font-medium hover:text-gray-900">
                {t("signup.termsOfService")}
              </Link>
            </div>
          </div>
        </div>
      </Form>
    </div>
  )
}

export default SignupForm
