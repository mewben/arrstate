import React from "react"
import { Link } from "gatsby"
import { useTranslation } from "react-i18next"

import { SignupForm, Hero } from "@Screens/auth"

const SignupPage = () => {
  const { t } = useTranslation()

  return (
    <div className="min-h-screen bg-white flex">
      <div className="flex flex-col items-center justify-center md:w-2/5 sm:w-1/2 p-4 ">
        <div className="flex-1 flex flex-col justify-center items-center w-full">
          <h1 className="font-medium text-3xl text-cool-gray-900">
            {t("signup.title")}
          </h1>
          <SignupForm />
        </div>
        <div className="text-gray-500">
          {t("signup.hasAccount")}{" "}
          <Link to="/signin" className="font-bold hover:text-gray-900">
            {t("signin.title")}
          </Link>
        </div>
      </div>
      <Hero />
    </div>
  )
}

export default SignupPage
