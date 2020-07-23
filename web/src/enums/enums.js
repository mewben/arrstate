export const ENTITIES = {
  PROJECT: "project",
  PROPERTY: "property",
  PERSON: "person",
  INVOICE: "invoice",
}

export const ROLES = {
  OWNER: "owner",
  COOWNER: "coowner",
  AGENT: "agent",
  CLIENT: "client",
}

export const PROPERTY_TYPES = {
  LOT: "lot",
  HOUSE: "house",
}

export const PAYMENT_SCHEMES = {
  INSTALLMENT: "installment",
  CASH: "cash",
}

export const PAYMENT_PERIODS = {
  MONTHLY: "monthly",
  YEARLY: "yearly",
}

export const GENERIC_BLOCKS = {
  CONTENT: "content",
}

export const INVOICE_BLOCKS = {
  ...GENERIC_BLOCKS,
  INTRO: "intro",
  ITEM: "item",
  SUMMARY: "summary",
}

export const INVOICE_STATUS = {
  DRAFT: "draft",
  PENDING: "pending",
  PAID: "paid",
  OVERDUE: "overdue",
}
