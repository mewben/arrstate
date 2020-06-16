import React from "react"

const AuthContext = React.createContext()
export const useAuth = () => React.useContext(AuthContext)

export const AuthProvider = ({ children }) => {
  const [isLoading, setIsLoading] = React.useState(false)
  const [isAuthenticated, setIsAuthenticated] = React.useState(false)

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
