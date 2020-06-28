import React from "react"
import { Helmet } from "react-helmet"

import MainSidebar from "./main-sidebar"

const LayoutContext = React.createContext()
export const useLayout = () => React.useContext(LayoutContext)

export const LayoutWrapper = ({ children }) => {
  const [isMenuOpen, setIsMenuOpen] = React.useState(false)
  return (
    <LayoutContext.Provider
      value={{
        isMenuOpen,
        setIsMenuOpen,
      }}
    >
      <Helmet
        bodyAttributes={{
          class: "h-screen flex overflow-hidden bg-gray-100",
        }}
      />
      <div className="app-root h-screen w-screen overflow-hidden">
        <div className="flex">
          <MainSidebar open={isMenuOpen} setIsOpen={setIsMenuOpen} />
          <main className="flex flex-col flex-1 overflow-hidden">
            <button onClick={() => setIsMenuOpen(true)}>open</button>
            {children}
          </main>
        </div>
      </div>
    </LayoutContext.Provider>
  )
}
