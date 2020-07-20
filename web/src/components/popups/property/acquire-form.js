import React from "react"
import * as Yup from "yup"
import acc from "accounting"
import { useWatch } from "react-hook-form"
import { useMutation, queryCache } from "react-query"

import { DrawerHeader } from "@Wrappers/layout"
import {
  Form,
  NumberField,
  PeopleSelectField,
  SubmitButton,
  SelectField,
} from "@Components/forms"
import { ROLES, PAYMENT_SCHEMES, PAYMENT_PERIODS } from "@Enums"
import { Error, Loading } from "@Components/generic"
import { requestApi } from "@Utils"
import { t, req } from "@Utils/t"
import { useProject } from "@Hooks"
import { map, values, sortBy } from "@Utils/lodash"
import { fromMoney } from "@Utils/money"

const validationSchema = Yup.object().shape({
  // propertyID: Yup.string().required(req),
  clientID: Yup.string().required(req).nullable(),
  paymentScheme: Yup.string().required(req), // cash, installmment
  // paymentPeriod: Yup.string().required(req), // monthly, yearly
  // paymentPeriod: Yup.string().when("paymentScheme", {
  //   is: val => val === PAYMENT_SCHEMES.INSTALLMENT,
  //   then: Yup.string().required(req),
  // }),
  terms: Yup.number().required(req).min(1, req).nullable(), // 60 months
  downPayment: Yup.number()
    .nullable()
    .when("paymentScheme", {
      is: val => val === PAYMENT_SCHEMES.INSTALLMENT,
      then: Yup.number().required(req).min(1, req),
    }),
  agentID: Yup.string().nullable(),
})

const Wrapper = ({ property, onClose }) => {
  const { status, data: project, error } = useProject(property.projectID)

  const { paymentSchemeOptions, paymentPeriodOptions } = React.useMemo(() => {
    const paymentSchemeOptions = map(values(PAYMENT_SCHEMES), item => ({
      label: t(`${item}`),
      value: item,
    }))
    const paymentPeriodOptions = map(values(PAYMENT_PERIODS), item => ({
      label: t(`${item}`),
      value: item,
    }))

    return { paymentSchemeOptions, paymentPeriodOptions }
  }, [])

  return status === "loading" ? (
    <Loading />
  ) : status === "error" ? (
    <Error error={error} />
  ) : (
    <AcquireForm
      project={project}
      property={property}
      paymentSchemeOptions={paymentSchemeOptions}
      paymentPeriodOptions={paymentPeriodOptions}
      onClose={onClose}
    />
  )
}

// ------ AcquireForm ------- //
const AcquireForm = ({
  property,
  project,
  paymentSchemeOptions,
  paymentPeriodOptions,
  onClose,
}) => {
  const [acquire, { reset, error }] = useMutation(
    formData => {
      return requestApi("/api/properties/acquire", "POST", {
        data: formData,
      })
    },
    {
      onSuccess: ({ data }) =>
        queryCache.setQueryData(["property", data._id], data),
    }
  )

  const onSubmit = async formData => {
    console.log("formData", formData)
    reset()
    const payload = {
      ...formData,
      propertyID: property._id,
      paymentPeriod: PAYMENT_PERIODS.MONTHLY,
    }
    const res = await acquire(payload)
    if (res) {
      onClose()
    }
  }

  const initialModel = {
    clientID: null,
    paymentScheme: PAYMENT_SCHEMES.INSTALLMENT,
    downPayment: null,
    terms: 1,
    agentID: null,
  }

  console.log("initialModel", initialModel)

  return (
    <div className="flex flex-col w-screen sm:w-96">
      <DrawerHeader title="Acquire Property" onClose={onClose}>
        <div className="mt-4 bg-cool-gray-600 p-4 rounded-sm">
          <div className="flex justify-between">
            <h2 className="text-white font-medium text-base">
              {property.name}
            </h2>
            <div className="uppercase text-xs tracking-widest">
              {property.type}
            </div>
          </div>
          {!!project && <div className="text-xs">{project?.name}</div>}
          <div className="mt-2 flex justify-between">
            <div>{acc.formatNumber(property.area, 2)} sq.m</div>
            <div className="text-green-300 font-medium">
              Php {acc.formatNumber(fromMoney(property.price), 2)}
            </div>
          </div>
        </div>
      </DrawerHeader>
      <Form
        onSubmit={onSubmit}
        validationSchema={validationSchema}
        model={initialModel}
      >
        <div className="grid grid-cols-12 gap-6 p-6">
          <Error error={error} className="col-span-12" />
          <PeopleSelectField name="clientID" label={t("client")} />
          <SelectField
            name="paymentScheme"
            label={t("payment scheme")}
            options={paymentSchemeOptions}
            disableClearable
          />
          <InstallmentForm paymentPeriodOptions={paymentPeriodOptions} />
          <PeopleSelectField
            name="agentID"
            role={[ROLES.AGENT]}
            label={t("agent")}
          />
          <div className="col-span-12">
            <SubmitButton>Submit</SubmitButton>
          </div>
        </div>
      </Form>
    </div>
  )
}

const InstallmentForm = ({ paymentPeriodOptions }) => {
  const paymentScheme = useWatch({
    name: "paymentScheme",
  })
  const paymentPeriod = useWatch({
    name: "paymentPeriod",
  })

  if (paymentScheme !== PAYMENT_SCHEMES.INSTALLMENT) return null

  console.log("paymentPeriod", paymentPeriod)

  return (
    <>
      <NumberField
        name="terms"
        label={t("terms")}
        endAddonInline={t(
          paymentPeriod === PAYMENT_PERIODS.MONTHLY ? "months" : "years"
        )}
        inputClassName="pr-16"
      />
      <NumberField
        name="downPayment"
        label={t("down payment")}
        startAddonInline="Php"
        placeholder="0.00"
        inputClassName="pl-12"
      />
    </>
  )
}

export default Wrapper
