import React from "react"
import * as Yup from "yup"
import acc from "accounting"
import { useMutation, queryCache } from "react-query"
import { useTranslation } from "react-i18next"

import { ERRORS } from "@Enums"
import { Error } from "@Components/generic"
import { DrawerHeader } from "@Wrappers/layout"
import { fromMoney, requestApi } from "@Utils"
import {
  Form,
  TextField,
  NumberField,
  PeopleSelectField,
  SubmitButton,
  SelectField,
} from "@Components/forms"

const PayForm = ({ invoice, onClose }) => {
  const { t } = useTranslation()

  const validationSchema = React.useMemo(() => {
    const req = t(ERRORS.REQUIRED)
    return Yup.object().shape({
      receiptNo: Yup.string().required(req),
    })
  }, [])

  const [pay, { reset, error }] = useMutation(
    formData => {
      return requestApi("/api/invoices/pay", "POST", {
        data: formData,
      })
    },
    {
      onSuccess: ({ data }) =>
        queryCache.setQueryData(["invoice", data._id], data),
    }
  )
  const onSubmit = async formData => {
    console.log("formData", formData)
    reset()
    const res = await pay({ ...formData, invoiceID: invoice._id })
    if (res) {
      onClose()
    }
  }

  const initialModel = {
    receiptNo: "",
  }

  return (
    <div className="flex flex-col w-screen sm:w-96">
      <DrawerHeader title={t("form.payment.title")} onClose={onClose}>
        <div className="mt-4 bg-cool-gray-600 p-4 rounded-sm">
          <div className="flex justify-between">
            <h2 className="text-white font-medium text-base">{invoice.name}</h2>
          </div>
          <div className="text-xs">
            {t("form.payment.invoiceNo")}: {invoice.no}
          </div>
          <div className="mt-2 flex justify-end">
            <div className="text-green-300 font-medium">
              Php {acc.formatNumber(fromMoney(invoice.total), 2)}
            </div>
          </div>
        </div>
      </DrawerHeader>
      <Form
        onSubmit={onSubmit}
        validationSchema={validationSchema}
        model={initialModel}
      >
        <div className="grid grid-cols-12 gap-x-6 gap-y-6 p-6">
          <Error error={error} className="col-span-12" />
          <TextField
            name="receiptNo"
            label={t("form.payment.receiptNo")}
            autoFocus
          />
          <div className="col-span-12">
            <SubmitButton>{t("btnSubmit")}</SubmitButton>
          </div>
        </div>
      </Form>
    </div>
  )
}

export default PayForm
