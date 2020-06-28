import React from "react"
import Drawer from "@material-ui/core/Drawer"

import AppMenu from "./app-menu"
import MainMenu from "./main-menu"

const MainSidebar = ({ open, setIsOpen }) => {
  const content = (
    <div className="flex w-64 h-full">
      <AppMenu />
      <MainMenu />
    </div>
  )

  return (
    <div className="sm:w-64 sm:flex-shrink-0">
      <div className="sm:hidden">
        <Drawer
          variant="temporary"
          anchor="left"
          // open={mobileOpen}
          open={open}
          onClose={() => setIsOpen(false)}
          ModalProps={{
            keepMounted: true, // Better open performance on mobile.
          }}
        >
          {content}
        </Drawer>
      </div>
      <div className="hidden sm:block">
        <Drawer variant="permanent" open>
          {content}
        </Drawer>
      </div>
    </div>
  )
}

export default MainSidebar
