import React from "react"
import { navigate } from "gatsby"
import * as Yup from "yup"
import { useMutation, queryCache } from "react-query"

import { Form, TextField, SubmitButton } from "@Components/forms"
import { Error } from "@Components/generic"
import { Button, ButtonConfirm } from "@Components/generic/button"
import { t } from "@Utils/t"
import { requestApi } from "@Utils"

const req = t("errors.required")
const validationSchema = Yup.object().shape({
  name: Yup.string().required(req),
  area: Yup.number().min(0),
})

// ------ ProjectForm ------- //
const ProjectForm = ({ model, onClose }) => {
  const isEdit = model?._id
  const [save, { reset, error }] = useMutation(
    formData => {
      return requestApi("/api/projects", isEdit ? "PUT" : "POST", {
        data: formData,
      })
    },
    {
      onSuccess: ({ data }) =>
        queryCache.setQueryData(["project", data._id], data),
    }
  )

  // TODO: handle error globally
  const [remove] = useMutation(() => {
    return requestApi(`/api/projects/${model._id}`, "DELETE")
  })

  const onSubmit = async formData => {
    reset()
    formData.area = +formData.area // convert to number

    const res = await save({
      ...model,
      ...formData,
    })
    if (res) {
      isEdit ? onClose() : navigate(`/projects/${res?.data?._id}/lots`)
    }
  }

  const onDelete = async () => {
    const res = await remove()
    console.log("res delete", res)
    if (res) {
      navigate("/projects", { replace: true })
    }
  }

  const initialModel = {
    name: "",
    area: 0,
    ...model,
  }

  return (
    <div>
      <Form
        onSubmit={onSubmit}
        validationSchema={validationSchema}
        model={initialModel}
      >
        <Error error={error} />
        <TextField name="name" label={t("project.name")} autoFocus />
        <TextField
          name="area"
          type="number"
          label={t("project.area")}
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

export default ProjectForm
