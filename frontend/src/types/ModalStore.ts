import { Snippet } from "svelte";

export type ModalStore = {
    isOpened: boolean,
    headerStyles?: string,
    title?: string,
    renderContent?: Snippet<[props?: any]>,
    props?: any
}
