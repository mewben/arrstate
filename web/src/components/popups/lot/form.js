import React from "react"
import { navigate } from "gatsby"
import * as Yup from "yup"
import { useMutation, queryCache } from "react-query"

import { Form, TextField, SubmitButton } from "@Components/forms"
import { Button, ButtonConfirm } from "@Components/generic/button"
import { t } from "@Utils/t"
import { requestApi, extractError } from "@Utils"

const req = t("errors.required")
const validationSchema = Yup.object().shape({
  name: Yup.string().required(req),
  area: Yup.number().min(0),
  price: Yup.number().min(0),
  priceAddon: Yup.number().min(0),
  projectID: Yup.string(),
})

// ------ LotForm ------- //
const LotForm = ({ model, project, onClose }) => {
  const isEdit = model?._id
  const [save, { reset, error, isError }] = useMutation(
    formData => {
      return requestApi("/api/lots", isEdit ? "PUT" : "POST", {
        data: formData,
      })
    },
    {
      onSuccess: ({ data }) => queryCache.setQueryData(["lot", data._id]),
    }
  )

  // TODO: handle error globally
  const [remove] = useMutation(() => {
    return requestApi(`/api/lots/${model._id}`, "DELETE")
  })

  const onSubmit = async formData => {
    console.log("onSubmit", formData)
    reset()
    if (formData.area) {
      formData.area = +formData.area // convert to number
    }

    if (formData.price) {
      formData.price = +formData.price
    }

    if (formData.priceAddon) {
      formData.priceAddon = +formData.priceAddon
    }

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

  const initialModel = {
    name: "",
    area: 0,
    price: 0,
    priceAddon: 0,
    ...model,
  }

  return (
    <div>
      <Form
        onSubmit={onSubmit}
        validationSchema={validationSchema}
        model={initialModel}
      >
        {isError && <div>{extractError(error)}</div>}
        <TextField name="projectName" label={t("project.name")} readOnly />
        <TextField name="name" label={t("lot.name")} autoFocus />
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

export default LotForm
