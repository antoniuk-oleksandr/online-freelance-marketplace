import { StatusEnum } from "@/types/StatusEnum"

export const formatStatus = (statusText: StatusEnum) => {
  const value = Object.values(StatusEnum)[statusText - 1] as string
  return value.replace(/([a-z])([A-Z])/g, '$1 $2')
}
