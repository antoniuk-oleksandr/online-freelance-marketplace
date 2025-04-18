import { writable } from "svelte/store";
import { ModalStore } from "@/types/ModalStore";

export const modalStore = writable<ModalStore>({
    isOpened: false,
});
