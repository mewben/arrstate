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

  const authSignIn = token => {
    authToken.process(token)
    setIsAuthenticated(!authToken.isExpired)
  }

  const authSignout = () => {
    if (!authToken) {
      authToken = new AuthToken()
    }
    authToken.clear()
    setIsAuthenticated(!authToken.isExpired)
  }

  return (
    <AuthContext.Provider
      value={{
        isLoading,
        isAuthenticated,
        authSignIn,
        authSignout,
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
    this.process(this.token)
  }

  process(token) {
    this.token = token
    // default to expired
    this.decodedToken = { sub: "", exp: 0 }
    // then try and decode the jwt using jwt-decode
    try {
      this.decodedToken = jwtDecode(this.token)
      this.store(token)
    } catch (e) {
      console.info("error jwtDecode", e)
    }
  }

  clear() {
    this.token = ""
    this.decodedToken = { sub: "", exp: 0 }
    localStorage.removeItem(TOKEN_STORAGE_KEY)
  }

  get authorizationString() {
    return `Bearer ${this.token}`
  }

  get expiresAt() {
    return new Date(this.decodedToken.exp * 1000)
  }

  get isExpired() {
    return new Date() > this.expiresAt
  }

  store(token) {
    localStorage.setItem(TOKEN_STORAGE_KEY, token)
  }
}
