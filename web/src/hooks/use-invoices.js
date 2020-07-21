import { useQuery } from "react-query"
import { requestApi } from "@Utils"

const fetchInvoices = async (_, params) => {
  const { data } = await requestApi("/api/invoices", "GET", { params })
  return data
}

const fetchInvoice = async (_, invoiceID) => {
  const { data } = await requestApi(`/api/invoices/${invoiceID}`)
  return data
}

export const useInvoices = (params = {}) => {
  return useQuery(["invoices", params], fetchInvoices)
}

export const useInvoice = invoiceID => {
  console.log("useInvoice", invoiceID)
  return useQuery(["invoice", invoiceID], fetchInvoice)
}
