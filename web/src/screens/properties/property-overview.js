import React from "react"
import acc from "accounting"

import { Panel, Portal, Button } from "@Components/generic"
import { AcquireForm } from "@Components/popups/property"

const PropertyOverview = ({ property }) => {
  console.log("property", property)
  return (
    <div className="grid grid-cols-6 gap-6">
      <div className="col-span-6 sm:col-span-3 p-4">
        <Panel noPadding>
          <div className="w-full flex items-center justify-between px-4 py-5 sm:p-6 space-x-6">
            <div className="flex-1 truncate">
              <div className="flex items-center space-x-3">
                <h3 className="text-gray-900 text-lg leading-5 font-medium truncate">
                  {property.name}
                </h3>
                <span className="flex-shrink-0 inline-block px-2 py-0.5 text-green-800 text-xs leading-4 font-medium bg-green-100 rounded-full">
                  Available
                </span>
              </div>
              <p className="mt-1 text-gray-400 leading-5 truncate">
                <span className="uppercase text-xs tracking-widest">
                  {property.type}
                </span>
              </p>
            </div>
            <div>
              <Portal openByClickOn={<Button>Acquire</Button>}>
                <AcquireForm property={property} />
              </Portal>
            </div>
          </div>
          <div className="flex p-6 space-x-6">
            <div>
              <div className="text-gray-400 text-xs">Area</div>
              <div className="mt-1 font-medium text-gray-900">
                {acc.formatNumber(property.area, 2)} sq.m
              </div>
            </div>
            <div>
              <div className="text-green-300 text-xs">Price</div>
              <div className="mt-1 font-medium text-green-500">
                Php {acc.formatNumber(property.price, 2)}
              </div>
            </div>
          </div>
          <div className="hidden">
            <div className="-mt-px flex">
              <div className="w-0 flex-1 flex border-r border-gray-200">
                <a
                  href="#"
                  className="relative -mr-px w-0 flex-1 inline-flex items-center justify-center py-4 text-sm leading-5 text-gray-700 font-medium border border-transparent rounded-bl-lg hover:text-gray-500 focus:outline-none focus:shadow-outline-blue focus:border-blue-300 focus:z-10 transition ease-in-out duration-150"
                >
                  <svg
                    className="w-5 h-5 text-gray-400"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z" />
                    <path d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z" />
                  </svg>
                  <span className="ml-3">Email</span>
                </a>
              </div>
              <div className="-ml-px w-0 flex-1 flex">
                <a
                  href="#"
                  className="relative w-0 flex-1 inline-flex items-center justify-center py-4 text-sm leading-5 text-gray-700 font-medium border border-transparent rounded-br-lg hover:text-gray-500 focus:outline-none focus:shadow-outline-blue focus:border-blue-300 focus:z-10 transition ease-in-out duration-150"
                >
                  <svg
                    className="w-5 h-5 text-gray-400"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path d="M2 3a1 1 0 011-1h2.153a1 1 0 01.986.836l.74 4.435a1 1 0 01-.54 1.06l-1.548.773a11.037 11.037 0 006.105 6.105l.774-1.548a1 1 0 011.059-.54l4.435.74a1 1 0 01.836.986V17a1 1 0 01-1 1h-2C7.82 18 2 12.18 2 5V3z" />
                  </svg>
                  <span className="ml-3">Call</span>
                </a>
              </div>
            </div>
          </div>
        </Panel>
      </div>
    </div>
  )
}

export default PropertyOverview
