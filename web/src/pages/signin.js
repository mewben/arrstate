import React from "react"
import { Link } from "gatsby"
import { useTranslation } from "react-i18next"

import { SigninForm, Hero } from "@Screens/auth"
import { PublicWrapper } from "@Wrappers"

const SigninPage = () => {
  const { t } = useTranslation()
  return (
    <PublicWrapper>
      <div className="min-h-screen bg-white flex">
        <div className="flex flex-col items-center justify-center md:w-2/5 sm:w-1/2 p-4 ">
          <div className="flex-1 flex flex-col justify-center items-center w-full">
            <h1 className="font-medium text-3xl text-cool-gray-900">
              {t("signin.title")}
            </h1>
            <SigninForm />
          </div>
          <div className="text-gray-500">
            {t("signin.noAccount")}{" "}
            <Link to="/signup" className="font-bold hover:text-gray-900">
              {t("signin.btnCreate")}
            </Link>
          </div>
        </div>
        <Hero />
      </div>
    </PublicWrapper>
  )
}

export default SigninPage
