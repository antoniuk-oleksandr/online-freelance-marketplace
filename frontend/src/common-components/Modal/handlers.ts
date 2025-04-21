import { modalStore } from "@/common-stores/modal-store";
import type { ModalStore } from "@/types/ModalStore.ts";

export const handleModalCloseButtonClick = () => {
    modalStore.update((prev) => ({
        ...prev,
        isOpened: false
    }))
}

export const handleModalBackdropClick = (
    e: MouseEvent & { currentTarget: (EventTarget & HTMLDivElement) },
    modalContentRef: HTMLElement | undefined,
    modalBackdropRef: HTMLElement | undefined
) => {
    if (!modalContentRef || !modalBackdropRef) return;

    if (!(modalBackdropRef).contains(e.target as Node)) return;
    if (modalContentRef.contains(e.target as Node)) return;

    modalStore.update((prev) => ({
        ...prev,
        isOpened: false
    }))
}

export const handleModalOpen = (
    modalData: ModalStore | undefined
) => {
    if (modalData && modalData.isOpened) {
        document.body.style.height = "100vh";
        document.body.style.overflow = "hidden";
    } else {
        document.body.style.overflow = "auto";
        document.body.style.height = "auto";
    }
}
