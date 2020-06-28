import React from "react"
import { Link } from "gatsby"

import { SigninForm, Hero } from "@Screens/auth"

const SigninPage = () => {
  return (
    <div className="min-h-screen bg-white flex">
      <div className="flex flex-col items-center justify-center md:w-2/5 sm:w-1/2 p-4 ">
        <div className="flex-1 flex flex-col justify-center items-center w-full">
          <h1 className="font-medium text-3xl text-cool-gray-900">Sign in</h1>
          <SigninForm />
        </div>
        <div className="text-gray-500">
          Don't have an account?{" "}
          <Link to="/signup" className="font-bold hover:text-gray-900">
            Create one
          </Link>
        </div>
      </div>
      <Hero />
    </div>
  )
}

export default SigninPage
