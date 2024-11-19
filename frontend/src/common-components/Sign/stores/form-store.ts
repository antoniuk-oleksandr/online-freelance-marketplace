import {writable} from "svelte/store";
import type {FormStore} from "@/types/FormStore.ts";

export const formStore = writable<FormStore>({
    data: null,
    errors: null,
    wasSubmitted: false,
    keepSignedIn: false,
})