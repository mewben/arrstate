import React from "react"
import { navigate } from "gatsby"
import * as Yup from "yup"
import { useMutation, queryCache } from "react-query"

import { ROLES } from "@Enums"
import {
  Form,
  TextField,
  SelectField,
  BaseTextField,
  SubmitButton,
  InputGroup,
  FieldError,
} from "@Components/forms"
import { Error } from "@Components/generic"
import { ButtonConfirm } from "@Components/generic/button"
import { t } from "@Utils/t"
import { DrawerHeader } from "@Wrappers/layout"
import { requestApi } from "@Utils"
import { isEmpty, map, values } from "@Utils/lodash"

const req = t("errors.required")
const validationSchema = Yup.object().shape({
  givenName: Yup.string().required(req),
  familyName: Yup.string(),
  email: Yup.string().email(t("errors.email")).required(req),
  role: Yup.array(),
})

// ------ PersonForm ------- //
const PersonForm = ({ model, onClose }) => {
  const isEdit = model?._id
  const [save, { reset, error }] = useMutation(
    formData => {
      return requestApi("/api/people", isEdit ? "PUT" : "POST", {
        data: formData,
      })
    },
    {
      onSuccess: ({ data }) => {
        queryCache.invalidateQueries(["people"])
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
    const res = await save(payload)
    if (res) {
      isEdit ? onClose() : navigate(`/people/${res?.data?._id}`)
    }
  }

  const onDelete = async () => {
    const res = await remove()
    console.log("res delete", res)
    if (res) {
      navigate("/people", { replace: true })
    }
  }

  const roleOptions = React.useMemo(() => {
    return map(values(ROLES), rol => ({
      label: t(`${rol}`),
      value: rol,
    }))
  }, [])

  const initialModel = {
    givenName: "",
    familyName: "",
    email: "",
    role: [],
    ...model,
  }

  return (
    <div className="flex flex-col w-screen sm:w-96">
      <DrawerHeader
        title={isEdit ? model.name : "New Person"}
        onClose={onClose}
      />
      <Form
        onSubmit={onSubmit}
        validationSchema={validationSchema}
        model={initialModel}
      >
        <div className="grid grid-cols-12 gap-6 p-6">
          <Error error={error} className="col-span-12" />
          <InputGroup name="givenName" id="givenName" label={t("name")}>
            <BaseTextField
              name="givenName"
              id="givenName"
              inputClassName="rounded-none rounded-l-md"
              placeholder={t("name.givenName")}
              autoFocus
            />
            <BaseTextField
              name="familyName"
              inputClassName="rounded-none rounded-r-md"
              placeholder={t("name.familyName")}
            />
          </InputGroup>
          <SelectField
            name="role"
            label={t("role")}
            options={roleOptions}
            multiple
          />
          <TextField name="email" label={t("email")} />
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

export default PersonForm
