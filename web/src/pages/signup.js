import React from "react"
import { Link } from "gatsby"

import { SignupForm } from "@Screens/auth"

const SigninPage = () => {
  return (
    <div>
      <h1>Sign up</h1>
      <SignupForm />
      <Link to="/signin">Sign in</Link>
    </div>
  )
}

export default SigninPage
