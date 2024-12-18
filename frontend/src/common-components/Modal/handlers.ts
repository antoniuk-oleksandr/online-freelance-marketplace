import {modalStore} from "@/common-stores/modal-store.ts";
import type {ModalStore} from "@/types/ModalStore.ts";

export const handleModalCloseButtonClick = () => {
    modalStore.update((prev) => ({
        ...prev,
        isOpened: false
    }))
}

export const handleModalBackdropClick = (
    e: MouseEvent & { currentTarget: (EventTarget & HTMLDivElement) }
) => {
    if (e.currentTarget.id !== 'modal-backdrop') return;

    modalStore.update((prev) => ({
        ...prev,
        isOpened: false
    }))
}

export const handleModalOpen = (
    modalData: ModalStore | undefined
) => {
    if (modalData && modalData.isOpened) {
        document.body.style.overflow = "hidden";
    } else document.body.style.overflow = "auto";
}