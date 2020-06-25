import React from "react"

import { Portal, Button } from "@Components/generic"
import { LotForm } from "@Components/popups/lot"

import { List } from "./components"

const LotList = () => {
  return (
    <div>
      <div>
        <h1>Lots</h1>
        <Portal openByClickOn={<Button>New Lot</Button>}>
          <LotForm />
        </Portal>
      </div>
      <List />
    </div>
  )
}

export default LotList
