import React from "react"
import { navigate } from "gatsby"
import * as Yup from "yup"
import { useMutation, queryCache } from "react-query"
import { useTranslation } from "react-i18next"

import {
  Form,
  TextField,
  NumberField,
  AddressField,
  SubmitButton,
} from "@Components/forms"
import { Error } from "@Components/generic"
import { ERRORS } from "@Enums"
import { ButtonConfirm } from "@Components/generic/button"
import { DrawerHeader } from "@Wrappers/layout"
import { requestApi } from "@Utils"

// ------ ProjectForm ------- //
const ProjectForm = ({ model, onClose }) => {
  const { t } = useTranslation()

  const validationSchema = React.useMemo(() => {
    const req = t(ERRORS.REQUIRED)
    return Yup.object().shape({
      name: Yup.string().required(req),
      area: Yup.number().min(0).nullable(),
    })
  }, [])

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
    const res = await save({
      ...model,
      ...formData,
    })
    if (res) {
      isEdit ? onClose() : navigate(`/projects/${res?.data?._id}/properties`)
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
    area: null,
    ...model,
  }

  return (
    <div className="flex flex-col w-screen sm:w-96">
      <DrawerHeader
        title={isEdit ? model.name : t("projects.new")}
        onClose={onClose}
      />
      <Form
        onSubmit={onSubmit}
        validationSchema={validationSchema}
        model={initialModel}
      >
        <div className="grid grid-cols-12 gap-6 p-6">
          <Error error={error} className="col-span-12" />
          <TextField name="name" label={t("form.project.name")} autoFocus />
          <NumberField name="area" label={t("form.project.area")} min={0} />
          <AddressField name="address" />
          <div className="col-span-12">
            <div className="flex items-center justify-between">
              <SubmitButton>{t("btnSubmit")}</SubmitButton>
              {isEdit && <ButtonConfirm onConfirm={onDelete} />}
            </div>
          </div>
        </div>
      </Form>
    </div>
  )
}

export default ProjectForm
