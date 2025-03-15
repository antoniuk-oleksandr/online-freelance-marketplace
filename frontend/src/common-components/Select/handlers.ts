import {modalStore} from "@/common-stores/modal-store.ts";
import {getSelectModal} from "@/common-components/Select/helper";
import {SelectMenuProps} from "@/types/SelectMenuProps.ts";
import {SelectProps} from "@/types/SelectProps.ts";

export const handleSelectClick = (
    clickAction: () => void,
    title: string,
    setSelectedItem?: (item: string) => void,
) => {
    setSelectedItem && setSelectedItem(title);
    clickAction();
}

export const handleSelectBackdropClick = (
    e: any,
    setIsOpen?: (value: boolean) => void,
) => {
    if (!setIsOpen) return;

    let {classList} = e.currentTarget.activeElement;
    if ([...classList].includes("Select-trigger")) return;

    ({classList} = (e.target));
    if (![...classList].includes("Select-menu")) setIsOpen(false);
}

export const handleSelectTriggerClick = (
    isOpen: boolean,
    title: string,
    SelectMenuProps: SelectProps,
    setIsOpen?: (value: boolean) => void,
) => {
    if (window.innerWidth < 1024) {
        modalStore.set({
            isOpened: true,
            title,
            renderContent: getSelectModal(SelectMenuProps as SelectMenuProps).render,
            headerStyles: "!p-3"
        })
    } else setIsOpen && setIsOpen(!isOpen);
}
