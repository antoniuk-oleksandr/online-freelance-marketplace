import {writable} from "svelte/store";
import {ErrorStore} from "@/types/ErrorStore.ts";

export const errorStore = writable<ErrorStore>({
    shown: false,
    error: undefined
});