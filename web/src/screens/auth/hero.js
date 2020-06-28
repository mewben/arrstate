import React from "react"

const Hero = () => {
  return (
    <div className="hidden sm:block relative flex-1">
      <img
        className="absolute inset-0 h-full w-full object-cover"
        src="https://images.unsplash.com/photo-1516971849755-77ddd3e004b0?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=1867&q=80"
        alt=""
      />
      <div className="absolute inset-0 flex items-center justify-center">
        <img className="h-20 w-20" src={`../../logo.png`} />
      </div>
    </div>
  )
}

export default Hero
