import React from "react"

import { Time } from "@Components/generic"
import { ProjectWrapper, PropertyWrapper, PersonWrapper } from "@Wrappers"
import { fullName } from "@Utils"

const Intro = ({ block, invoice }) => {
  return (
    <div>
      <h1 className="text-xl font-medium">Invoice</h1>
      <div className="mt-4">
        <div>
          <span className="text-gray-500 mr-2">Invoice #:</span>
          {invoice.no}
        </div>
        <div>
          <span className="text-gray-500 mr-2">Invoice Ref:</span>
          {invoice.name}
        </div>
        {!!invoice.issueDate && (
          <div>
            <span className="text-gray-500 mr-2">Issue Date:</span>
            <Time d={invoice.issueDate} dateOnly />
          </div>
        )}
        {!!invoice.dueDate && (
          <div>
            <span className="text-gray-500 mr-2">Due Date:</span>
            <Time d={invoice.dueDate} dateOnly />
          </div>
        )}
        {!!invoice.paidAt && (
          <div>
            <span className="text-gray-500 mr-2">Paid Date:</span>
            <Time d={invoice.paidAt} dateOnly />
          </div>
        )}
        {!!invoice.projectID && (
          <ProjectWrapper projectID={invoice.projectID}>
            {({ project }) => {
              return (
                <div>
                  <span className="text-gray-500 mr-2">Project:</span>
                  {project.name}
                </div>
              )
            }}
          </ProjectWrapper>
        )}
        {!!invoice.propertyID && (
          <PropertyWrapper propertyID={invoice.propertyID}>
            {({ property }) => {
              return (
                <div>
                  <span className="text-gray-500 mr-2">Property:</span>
                  {property.name}
                </div>
              )
            }}
          </PropertyWrapper>
        )}
      </div>
      <div className="mt-8 flex items-start w-full space-x-4">
        <div className="text-gray-400 uppercase tracking-wider">To:</div>
        {!!invoice?.to?._id && (
          <PersonWrapper personID={invoice?.to?._id}>
            {({ person }) => {
              return (
                <div className="text-gray-600">
                  <div className="font-medium text-gray-900">
                    {fullName(person.givenName, person.familyName)}
                  </div>
                  <div className="mt-1">
                    <div>{person.email}</div>
                  </div>
                </div>
              )
            }}
          </PersonWrapper>
        )}
      </div>
    </div>
  )
}

export default Intro
