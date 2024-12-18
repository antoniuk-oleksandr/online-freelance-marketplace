import {modalStore} from "@/common-stores/modal-store.ts";
import {getDropdownModal} from "@/common-components/Dropdown/helper.ts";
import {DropdownMenuProps} from "@/types/DropdownMenuProps.ts";
import {DropdownProps} from "@/types/DropdownProps.ts";

export const handleDropdownClick = (
    clickAction: () => void,
    title: string,
    setSelectedItem?: (item: string) => void,
) => {
    setSelectedItem && setSelectedItem(title);
    clickAction();
}

export const handleDropdownBackdropClick = (
    e: any,
    setIsOpen?: (value: boolean) => void,
) => {
    if (!setIsOpen) return;

    let {classList} = e.currentTarget.activeElement;
    if ([...classList].includes("dropdown-trigger")) return;

    ({classList} = (e.target));
    if (![...classList].includes("dropdown-menu")) setIsOpen(false);
}

export const handleDropdownTriggerClick = (
    isOpen: boolean,
    title: string,
    dropdownMenuProps: DropdownProps,
    setIsOpen?: (value: boolean) => void,
) => {
    if (window.innerWidth < 1024) {
        modalStore.set({
            isOpened: true,
            title,
            renderContent: getDropdownModal(dropdownMenuProps as DropdownMenuProps).render,
            headerStyles: "!p-3"
        })
    } else setIsOpen && setIsOpen(!isOpen);
}