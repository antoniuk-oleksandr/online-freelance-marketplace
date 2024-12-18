import {Snippet} from "svelte";

export type ModalStore = {
    isOpened: boolean,
    headerStyles?: string,
    title?: string,
    renderContent?: (targetElement: HTMLElement) => void
}