import React from "react"
import { navigate } from "gatsby"
import * as Yup from "yup"
import { useMutation, queryCache } from "react-query"
import { useTranslation } from "react-i18next"

import { ROLES, ERRORS } from "@Enums"
import {
  Form,
  TextField,
  BaseTextField,
  SubmitButton,
  InputGroup,
  FieldError,
} from "@Components/forms"
import { Error } from "@Components/generic"
import { ButtonConfirm } from "@Components/generic/button"
import { DrawerHeader } from "@Wrappers/layout"
import { requestApi } from "@Utils"
import { isEmpty } from "@Utils/lodash"

// ------ ClientForm ------- //
const ClientForm = ({ model, onClose }) => {
  const { t } = useTranslation()

  const validationSchema = React.useMemo(() => {
    const req = t(ERRORS.REQUIRED)
    return Yup.object().shape({
      givenName: Yup.string().required(req),
      familyName: Yup.string(),
      email: Yup.string().email(t(ERRORS.EMAIL)).required(req),
    })
  }, [])

  const isEdit = model?._id
  const [save, { reset, error }] = useMutation(
    formData => {
      return requestApi("/api/people", isEdit ? "PUT" : "POST", {
        data: formData,
      })
    },
    {
      onSuccess: ({ data }) => {
        queryCache.invalidateQueries(["people", { role: [ROLES.CLIENT] }])
        queryCache.setQueryData(["people", data._id], data)
      },
    }
  )

  // TODO: handle error globally
  const [remove] = useMutation(() => {
    return requestApi(`/api/people/${model._id}`, "DELETE")
  })

  const onSubmit = async formData => {
    reset()
    const payload = {
      ...model,
      ...formData,
    }
    if (isEmpty(payload.role)) {
      payload.role = [ROLES.CLIENT]
    }
    const res = await save(payload)
    if (res) {
      isEdit ? onClose() : navigate(`/clients/${res?.data?._id}`)
    }
  }

  const onDelete = async () => {
    const res = await remove()
    console.log("res delete", res)
    if (res) {
      navigate("/clients", { replace: true })
    }
  }

  const initialModel = {
    givenName: "",
    familyName: "",
    email: "",
    ...model,
  }

  return (
    <div className="flex flex-col w-screen sm:w-96">
      <DrawerHeader
        title={isEdit ? model.name : t("clients.new")}
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
            <InputGroup label={t("name.fullName")}>
              <BaseTextField
                name="givenName"
                className="rounded-none rounded-l-md"
                placeholder={t("name.givenName")}
                autoFocus
              />
              <BaseTextField
                name="familyName"
                className="rounded-none rounded-r-md"
                placeholder={t("name.familyName")}
              />
            </InputGroup>
            <FieldError name="givenName" />
          </div>
          <div className="col-span-6">
            <TextField name="email" label={t("email")} />
          </div>
          <div className="col-span-6">
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

export default ClientForm
