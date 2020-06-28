import React from "react"
import { navigate } from "gatsby"
import * as Yup from "yup"
import { useMutation, queryCache } from "react-query"

import { Form, TextField, SubmitButton } from "@Components/forms"
import { Error } from "@Components/generic"
import { Button, ButtonConfirm } from "@Components/generic/button"
import { t } from "@Utils/t"
import { DrawerHeader } from "@Wrappers/layout"
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
      onSuccess: ({ data }) => {
        queryCache.invalidateQueries("projects", { exact: true })
        queryCache.setQueryData(["project", data._id], data)
      },
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
    <div className="flex flex-col w-screen sm:w-96">
      <DrawerHeader
        title={isEdit ? model.name : "New Project"}
        onClose={onClose}
      />
      <Form
        onSubmit={onSubmit}
        validationSchema={validationSchema}
        model={initialModel}
      >
        <Error error={error} />
        <div className="grid grid-cols-6 gap-6 p-6">
          <div className="col-span-6">
            <TextField name="name" label={t("project.name")} autoFocus />
          </div>
          <div className="col-span-6">
            <TextField
              name="area"
              type="number"
              label={t("project.area")}
              step="0.0001"
              min="0"
            />
          </div>
          <div className="col-span-6">
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

export default ProjectForm
