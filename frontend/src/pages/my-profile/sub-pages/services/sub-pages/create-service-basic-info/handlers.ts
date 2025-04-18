import type { CreateServiceBasicInfoFormData } from "@/types/CreateServiceBasicInfoFormData";
import { createServiceStore } from "../../stores/create-service-store";
import { navigate } from "svelte-routing";

export const handleCreateServiceBasicInfoFormSubmit = (data: CreateServiceBasicInfoFormData) => {
  createServiceStore.update((prev) => ({ ...prev, ...data }))

  navigate("/my-profile/services/create/packages");
}
