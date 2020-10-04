import React from "react"
import { Link, navigate } from "gatsby"
import * as Yup from "yup"
import { useMutation, queryCache } from "react-query"
import { useTranslation } from "react-i18next"

import { Form, TextField, SubmitButton } from "@Components/forms"
import { Error } from "@Components/generic"
import { requestApi } from "@Utils"
import { useAuth } from "@Providers"
import { ERRORS } from "@Enums"

// ------- SigninForm -------- //
const SigninForm = () => {
  const { t } = useTranslation()
  const { authSignIn } = useAuth()

  const validationSchema = React.useMemo(() => {
    const req = t(ERRORS.REQUIRED)
    return Yup.object().shape({
      email: Yup.string().email(t(ERRORS.EMAIL)).required(req),
      password: Yup.string().required(req),
    })
  }, [])

  const [signin, { reset, error }] = useMutation(
    formData => {
      return requestApi("/auth/signin", "POST", {
        data: formData,
        noToken: true,
      })
    },
    {
      onSuccess: ({ data }) =>
        queryCache.setQueryData("currentUser", data?.user),
    }
  )

  const onSubmit = async formData => {
    reset()
    const res = await signin(formData)
    if (res) {
      // store token,
      // redirect to '/'
      authSignIn(res.data.token)
      navigate("/", { replace: true })
    }
  }

  return (
    <div className="mt-6 w-full max-w-sm">
      <Form
        onSubmit={onSubmit}
        validationSchema={validationSchema}
        model={{ email: "", password: "" }}
      >
        <div className="grid grid-cols-12 gap-x-6 gap-y-6">
          <Error error={error} className="col-span-12" />
          <TextField
            name="email"
            type="email"
            label={t("email")}
            autoComplete="off"
            autoFocus
          />
          <TextField name="password" label={t("password")} type="password" />
          <div className="col-span-12">
            <SubmitButton size="xl" fullWidth>
              {t("signin.btn")}
            </SubmitButton>
            <div className="text-xs text-cool-gray-500 mt-2">
              <Link
                to="/forgot-password"
                className="font-medium hover:text-gray-900"
              >
                {t("signin.forgot")}
              </Link>
            </div>
          </div>
        </div>
      </Form>
    </div>
  )
}

export default SigninForm
