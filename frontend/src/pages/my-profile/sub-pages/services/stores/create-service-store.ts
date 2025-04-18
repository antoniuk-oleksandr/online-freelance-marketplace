import type { CreateServiceFormData } from "@/types/CreateServiceFormData";
import { writable } from "svelte/store";

export const createServiceStore = writable<CreateServiceFormData>({
  title: '',
  description: '',
  categoryId: -1,
  packages: [],
  questions: [],
  files: [],
})
