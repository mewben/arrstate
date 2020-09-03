import React from "react"
import { navigate } from "gatsby"
import { useWatch } from "react-hook-form"
import { useMutation, queryCache } from "react-query"
import { useTranslation } from "react-i18next"

import { ROLES } from "@Enums"
import {
  Form,
  TextField,
  SelectField,
  BaseTextField,
  SubmitButton,
  InputGroup,
  AddressField,
  NumberField,
} from "@Components/forms"
import { Error } from "@Components/generic"
import { ButtonConfirm } from "@Components/generic/button"
import { DrawerHeader } from "@Wrappers/layout"
import { requestApi, getValidationSchema } from "@Utils"
import { includes, map, values } from "@Utils/lodash"

// ------ PersonForm ------- //
const PersonForm = ({ model, onClose }) => {
  const { t } = useTranslation()

  const validationSchema = React.useMemo(() => {
    return getValidationSchema(t, "person")
  }, [t])

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
    name: {
      first: "",
      last: "",
    },
    email: "",
    role: [],
    ...model,
  }

  return (
    <div className="flex flex-col w-screen sm:w-96">
      <DrawerHeader
        title={isEdit ? model.name : t("people.new")}
        onClose={onClose}
      />
      <Form
        onSubmit={onSubmit}
        validationSchema={validationSchema}
        model={initialModel}
      >
        <div className="grid grid-cols-12 gap-6 p-6">
          <Error error={error} className="col-span-12" />
          <InputGroup
            name="name.first"
            id="givenName"
            label={t("name.fullName")}
          >
            <BaseTextField
              name="name.first"
              id="givenName"
              inputClassName="rounded-none rounded-l-md"
              placeholder={t("name.givenName")}
              autoFocus
            />
            <BaseTextField
              name="name.last"
              inputClassName="rounded-none rounded-r-md"
              placeholder={t("name.familyName")}
            />
          </InputGroup>
          <SelectField
            name="role"
            label={t("form.people.role")}
            options={roleOptions}
            multiple
          />
          <CommissionForm />
          <TextField name="email" label={t("email")} />
          <AddressField name="address" />
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

const CommissionForm = () => {
  const { t } = useTranslation()
  const roles = useWatch({
    name: "role",
  })

  if (!includes(roles, ROLES.AGENT)) {
    return null
  }
  return (
    <>
      <NumberField
        name="commissionPerc"
        label={t("form.people.commission")}
        endAddonInline="%"
        inputClassName="pr-16"
        isMoney
      />
    </>
  )
}

export default PersonForm
