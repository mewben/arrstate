import React from "react"
import PropTypes from "prop-types"
import { navigate } from "gatsby"
import * as Yup from "yup"
import { useMutation, queryCache } from "react-query"

import { Form, TextField, SelectField, SubmitButton } from "@Components/forms"
import { Error, Loading } from "@Components/generic"
import { Button, ButtonConfirm } from "@Components/generic/button"
import { t } from "@Utils/t"
import { useProjectOptions } from "@Hooks"
import { get, map } from "@Utils/lodash"
import { requestApi } from "@Utils"

const req = t("errors.required")
const validationSchema = Yup.object().shape({
  name: Yup.string().required(req),
  area: Yup.number().min(0),
  price: Yup.number().min(0),
  priceAddon: Yup.number().min(0),
  // projectID: Yup.mixed().notOneOf([null, undefined], req),
})

// ------ LotForm ------- //
const LotForm = ({ model, projectID, onClose }) => {
  const isEdit = model?._id
  const [save, { reset, error }] = useMutation(
    formData => {
      return requestApi("/api/lots", isEdit ? "PUT" : "POST", {
        data: formData,
      })
    },
    {
      onSuccess: ({ data }) => queryCache.setQueryData(["lot", data._id], data),
    }
  )

  // TODO: handle error globally
  const [remove] = useMutation(() => {
    return requestApi(`/api/lots/${model._id}`, "DELETE")
  })

  const onSubmit = async formData => {
    reset()

    // prepare formData
    formData.area = +formData.area // convert to number
    formData.price = +formData.price
    formData.priceAddon = +formData.priceAddon
    formData.projectID = get(formData, "projectID.value", null)

    console.log("formData", formData)

    const res = await save({
      ...model,
      ...formData,
    })
    if (res) {
      isEdit ? onClose() : navigate(`/lots/${res?.data?._id}`)
    }
  }

  const onDelete = async () => {
    const res = await remove()
    console.log("res delete", res)
    if (res) {
      navigate("/lots", { replace: true })
    }
  }

  const {
    status,
    options: projectOptions,
    selectedOption: projectOption,
    error: projectError,
  } = useProjectOptions(get(model, "projectID", projectID))

  if (status === "loading") {
    return <Loading />
  } else if (status === "error") {
    return <Error error={projectError} />
  }

  const initialModel = {
    name: "",
    area: 0,
    price: 0,
    priceAddon: 0,
    ...model,
    projectID: projectOption,
  }

  return (
    <div>
      <Form
        onSubmit={onSubmit}
        validationSchema={validationSchema}
        model={initialModel}
      >
        <Error error={error} />
        <TextField name="name" label={t("lot.name")} autoFocus />
        <SelectField
          name="projectID"
          label={t("lot.project")}
          options={projectOptions}
        />
        <TextField
          name="area"
          type="number"
          label={t("lot.area")}
          step="0.0001"
          min="0"
        />
        <TextField
          name="price"
          type="number"
          label={t("lot.price")}
          step="0.0001"
          min="0"
        />
        <TextField
          name="priceAddon"
          type="number"
          label={t("lot.priceAddon")}
          step="0.0001"
          min="0"
        />
        {isEdit && <ButtonConfirm onConfirm={onDelete} />}
        <SubmitButton>Submit</SubmitButton>
        <Button onClick={onClose}>Close</Button>
      </Form>
    </div>
  )
}

LotForm.propTypes = {
  model: PropTypes.object,
  projectID: PropTypes.string,
}

export default LotForm
