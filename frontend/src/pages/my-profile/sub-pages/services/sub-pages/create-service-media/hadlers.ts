import { navigate } from "svelte-routing"
import { createServiceStore } from "../../stores/create-service-store"
import type { CreateServiceMediaFormData } from "@/types/CreateServiceMediaFormData"

export const handleCreateServiceMediaFormSubmit = (data: CreateServiceMediaFormData) => {
  createServiceStore.update((prev) => ({ ...prev, ...data }))

  console.log(data);
  

  // navigate('/my-profile/services/create/requirements')
}
