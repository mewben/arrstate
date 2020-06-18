import React from "react"
import { Link } from "gatsby"

import { SigninForm } from "@Screens/auth"

const SigninPage = () => {
  return (
    <div>
      <h1>Signin</h1>
      <SigninForm />
      <Link to="/signup">Create Account</Link>
    </div>
  )
}

export default SigninPage
