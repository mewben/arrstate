import React from "react"

import { Portal, Button } from "@Components/generic"
import { LotForm } from "@Components/popups/lot"

const Header = ({ lot }) => {
  return (
    <div>
      <h1>{lot.name}</h1>
      <Portal openByClickOn={<Button>Edit Lot</Button>}>
        <LotForm model={lot} />
      </Portal>
    </div>
  )
}

export default Header
