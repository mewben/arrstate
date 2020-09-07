import React from "react"

const Dtd = ({ label, value }) => {
  return (
    <div className="sm:flex sm:space-x-6 sm:border-t sm:border-cool-gray-200 sm:px-6 sm:py-5">
      <dt className="text-sm leading-5 font-medium text-cool-gray-500 sm:w-40 sm:flex-shrink-0 lg:w-48">
        {label}
      </dt>
      <dd className="mt-1 text-sm leading-5 text-cool-gray-900 sm:mt-0 sm:col-span-2">
        {value}
      </dd>
    </div>
  )
}

export default Dtd
