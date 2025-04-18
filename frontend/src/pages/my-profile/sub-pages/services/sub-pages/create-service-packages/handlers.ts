import type { CreateServicePackageFormData } from "@/types/CreateServicePackageFormData"
import { createServiceStore } from "../../stores/create-service-store"
import { navigate } from "svelte-routing"

export const handleCreateServicePackagesFormSubmit = (data: CreateServicePackageFormData) => {
  createServiceStore.update((prev) => ({ ...prev, ...data }))

  navigate('/my-profile/services/create/requirements')
}
