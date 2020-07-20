import React from "react"
import PropTypes from "prop-types"
import { navigate } from "gatsby"
import * as Yup from "yup"
import { useMutation, queryCache } from "react-query"

import { PROPERTY_TYPES } from "@Enums"
import {
  Form,
  TextField,
  NumberField,
  SelectField,
  SubmitButton,
} from "@Components/forms"
import { Error, Loading } from "@Components/generic"
import { Button, ButtonConfirm } from "@Components/generic/button"
import { t, req } from "@Utils/t"
import { useProjectOptions } from "@Hooks"
import { DrawerHeader } from "@Wrappers/layout"
import { get, map, values } from "@Utils/lodash"
import { requestApi } from "@Utils"

const validationSchema = Yup.object().shape({
  name: Yup.string().required(req),
  type: Yup.string().required(req).nullable(),
  area: Yup.number().min(0).nullable(),
  price: Yup.number().min(0).nullable(),
  priceAddon: Yup.number().min(0).nullable(),
  // projectID: Yup.mixed().notOneOf([null, undefined], req),
})

// ------ PropertyForm ------- //
const PropertyForm = ({ model, projectID, onClose }) => {
  const isEdit = model?._id
  const [save, { reset, error }] = useMutation(
    formData => {
      return requestApi("/api/properties", isEdit ? "PUT" : "POST", {
        data: formData,
      })
    },
    {
      onSuccess: ({ data }) =>
        queryCache.setQueryData(["property", data._id], data),
    }
  )

  // TODO: handle error globally
  const [remove] = useMutation(() => {
    return requestApi(`/api/properties/${model._id}`, "DELETE")
  })

  const onSubmit = async formData => {
    reset()
    console.log("formData", formData)

    const res = await save({
      ...model,
      ...formData,
    })
    if (res) {
      isEdit ? onClose() : navigate(`/properties/${res?.data?._id}`)
    }
  }

  const onDelete = async () => {
    const res = await remove()
    console.log("res delete", res)
    if (res) {
      navigate("/properties", { replace: true })
    }
  }

  const {
    status,
    options: projectOptions,
    error: projectError,
  } = useProjectOptions()

  const propertyTypeOptions = React.useMemo(() => {
    return map(values(PROPERTY_TYPES), typ => ({
      label: t(`${typ}`),
      value: typ,
    }))
  }, [])

  if (status === "loading") {
    return <Loading />
  } else if (status === "error") {
    return <Error error={projectError} />
  }

  const initialModel = {
    projectID,
    name: "",
    type: PROPERTY_TYPES.LOT,
    area: null,
    price: null,
    priceAddon: null,
    ...model,
  }

  return (
    <div className="flex flex-col w-screen sm:w-96">
      <DrawerHeader
        title={isEdit ? model.name : "New Property"}
        onClose={onClose}
      />
      <Form
        onSubmit={onSubmit}
        validationSchema={validationSchema}
        model={initialModel}
      >
        <div className="grid grid-cols-12 gap-6 p-6">
          <Error error={error} className="col-span-12" />
          <TextField name="name" label={t("property.name")} autoFocus />
          <SelectField
            name="type"
            label={t("property.type")}
            options={propertyTypeOptions}
            disableClearable
          />
          <SelectField
            name="projectID"
            label={t("property.project")}
            options={projectOptions}
          />
          <NumberField
            name="area"
            label={t("property.area")}
            min={0}
            endAddonInline="sq.m"
            inputClassName="pr-16"
            placeholder="0.00"
          />
          <NumberField
            name="price"
            label={t("property.price")}
            min={0}
            isMoney
            startAddonInline="Php"
            placeholder="0.00"
            inputClassName="pl-12"
          />
          <NumberField
            name="priceAddon"
            label={t("property.priceAddon")}
            min={0}
            isMoney
            startAddonInline="Php"
            placeholder="0.00"
            inputClassName="pl-12"
          />
          <div className="col-span-12">
            <div className="flex items-center justify-between">
              <SubmitButton>Submit</SubmitButton>
              {isEdit && <ButtonConfirm onConfirm={onDelete} />}
            </div>
          </div>
        </div>
      </Form>
    </div>
  )
}

PropertyForm.propTypes = {
  model: PropTypes.object,
  projectID: PropTypes.string,
}

export default PropertyForm
