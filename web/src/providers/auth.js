import React from "react"
import jwtDecode from "jwt-decode"

import { TOKEN_STORAGE_KEY } from "@Enums/config"

const AuthContext = React.createContext()
export const useAuth = () => React.useContext(AuthContext)
export let authToken

export const AuthProvider = ({ children }) => {
  const [isLoading, setIsLoading] = React.useState(true)
  const [isAuthenticated, setIsAuthenticated] = React.useState(false)

  React.useEffect(() => {
    const initAuth = () => {
      authToken = new AuthToken()
      setIsAuthenticated(!authToken.isExpired)
      setIsLoading(false)
    }

    initAuth()
  }, [])

  console.log("isLoading", isLoading)
  console.log("isAuthenticated", isAuthenticated)

  return (
    <AuthContext.Provider
      value={{
        isLoading,
        isAuthenticated,
      }}
    >
      {children}
    </AuthContext.Provider>
  )
}

class AuthToken {
  token
  decodedToken

  constructor(token) {
    if (!token) {
      // try to get from the localStorage
      this.token = localStorage.getItem(TOKEN_STORAGE_KEY)
    } else {
      this.token = token
    }
    // default to expired
    this.decodedToken = { sub: "", exp: 0 }
    // then try and decode the jwt using jwt-decode
    try {
      console.log("thistoken", this.token)
      this.decodedToken = jwtDecode(this.token)
    } catch (e) {
      console.info("error jwtDecode", e)
    }
  }

  get authorizationString() {
    return `Bearer ${this.token}`
  }

  get expiresAt() {
    return new Date(this.decodedToken.exp * 1000)
  }

  get isExpired() {
    console.log("decodedToken", this.decodedToken)
    console.log(new Date())
    console.log(this.expiresAt)
    console.log("isExpired")
    return new Date() > this.expiresAt
  }

  static store(token) {
    localStorage.setItem(TOKEN_STORAGE_KEY, token)
  }
}
