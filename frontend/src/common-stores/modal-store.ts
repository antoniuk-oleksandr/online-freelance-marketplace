import {writable} from "svelte/store";
import {ModalStore} from "@/types/ModalStore.ts";

export const modalStore = writable<ModalStore>({
    isOpened: false,
});