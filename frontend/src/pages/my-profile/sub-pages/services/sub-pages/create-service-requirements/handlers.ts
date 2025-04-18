import { navigate } from "svelte-routing"
import { createServiceStore } from "../../stores/create-service-store"
import type { CreateServiceRequirementsFormData } from "@/types/CreateServiceRequirementsFormData"

export const handleCreateServiceRequirementsFormSubmit = (data: CreateServiceRequirementsFormData) => {
  createServiceStore.update((prev) => ({ ...prev, ...data }))

  navigate('/my-profile/services/create/media')
}
