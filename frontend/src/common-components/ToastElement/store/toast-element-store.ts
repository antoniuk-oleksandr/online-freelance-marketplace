import {writable} from "svelte/store";
import {ToastElementStore} from "@/types/ToastElementStore.ts";

export const toastElementStore = writable<ToastElementStore>({
    message: "",
    type: "error",
    show: false,
    exitAnimation: false,
});