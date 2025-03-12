import { writable } from "svelte/store";

export const orderSubmitRequirementsFormFilesStore = writable<File[]>([]);
