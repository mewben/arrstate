import React from "react"
import { useTranslation } from "react-i18next"

import { Time } from "@Components/generic"
import { ProjectWrapper, PropertyWrapper, PersonWrapper } from "@Wrappers"
import { fullName } from "@Utils"

const Intro = ({ block, invoice, isReceipt }) => {
  const { t } = useTranslation()
  const renderReceipt = () => {
    return (
      <>
        <h1 className="text-xl font-medium">{t("receipts.title")}</h1>
        <div className="mt-4">
          <div>
            <span className="text-gray-500 mr-2">{t("receipts.no")}:</span>
            {invoice.receiptNo}
          </div>
          <div>
            <span className="text-gray-500 mr-2">{t("invoices.no")}:</span>
            {invoice.no}
          </div>
          <div>
            <span className="text-gray-500 mr-2">
              {t("invoices.paidDate")}:
            </span>
            <Time d={invoice.paidAt} dateOnly />
          </div>
        </div>
      </>
    )
  }

  const renderInvoice = () => {
    return (
      <>
        <h1 className="text-xl font-medium">{t("invoices.title")}</h1>
        <div className="mt-4">
          <div>
            <span className="text-gray-500 mr-2">{t("invoices.no")}:</span>
            {invoice.no}
          </div>
          <div>
            <span className="text-gray-500 mr-2">{t("invoices.ref")}:</span>
            {invoice.name}
          </div>
          {!!invoice.issueDate && (
            <div>
              <span className="text-gray-500 mr-2">
                {t("invoices.issueDate")}:
              </span>
              <Time d={invoice.issueDate} dateOnly />
            </div>
          )}
          {!!invoice.dueDate && (
            <div>
              <span className="text-gray-500 mr-2">
                {t("invoices.dueDate")}:
              </span>
              <Time d={invoice.dueDate} dateOnly />
            </div>
          )}
          {!!invoice.paidAt && (
            <div>
              <span className="text-gray-500 mr-2">
                {t("invoices.paidDate")}:
              </span>
              <Time d={invoice.paidAt} dateOnly />
            </div>
          )}
          {!!invoice.projectID && (
            <ProjectWrapper projectID={invoice.projectID}>
              {({ project }) => {
                return (
                  <div>
                    <span className="text-gray-500 mr-2">
                      {t("invoices.project")}:
                    </span>
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
                    <span className="text-gray-500 mr-2">
                      {t("invoices.property")}:
                    </span>
                    {property.name}
                  </div>
                )
              }}
            </PropertyWrapper>
          )}
        </div>
      </>
    )
  }

  return (
    <div>
      {isReceipt ? renderReceipt() : renderInvoice()}
      <div className="mt-8 flex items-start w-full space-x-4">
        <div className="text-gray-400 uppercase tracking-wider">
          {t("invoices.to")}:
        </div>
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
