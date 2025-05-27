import { writable } from "svelte/store";
import { ErrorStore } from "@/types/ErrorStore";

export const errorStore = writable<ErrorStore>({
    shown: false,
    error: undefined
});
